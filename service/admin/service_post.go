package admin

import (
	"context"
	"time"

	"github.com/bots-house/birzzha/core"
	"github.com/bots-house/birzzha/store"
	"github.com/pkg/errors"
	"github.com/volatiletech/null/v8"
)

type PostInput struct {
	LotID                 core.LotID
	Text                  string
	Title                 string
	DisableWebPagePreview bool
	ScheduledAt           time.Time
}

var (
	ErrLotIsAlreadyPublished = core.NewError("lot_is_already_published", "lot is already published")
)

func (srv *Service) CreatePost(ctx context.Context, user *core.User, in *PostInput) (*core.Post, error) {
	if err := srv.IsAdmin(user); err != nil {
		return nil, err
	}
	post := core.NewPost(in.LotID, in.Text, in.Title, in.DisableWebPagePreview, in.ScheduledAt)

	if err := srv.Txier(ctx, func(ctx context.Context) error {
		lot, err := srv.Lot.Query().ID(in.LotID).One(ctx)
		if err != nil {
			return errors.Wrap(err, "get lot")
		}

		if lot.Status == core.LotStatusPublished {
			return ErrLotIsAlreadyPublished
		}

		if err := srv.Post.Add(ctx, post); err != nil {
			return errors.Wrap(err, "add post")
		}

		lot.Status = core.LotStatusScheduled
		lot.ScheduledAt = null.TimeFrom(in.ScheduledAt)
		if err := srv.Lot.Update(ctx, lot); err != nil {
			return errors.Wrap(err, "update lot")
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return post, nil
}

type PostItem struct {
	*core.Post
	LotName  string
	Username null.String
	JoinLink null.String
	Avatar   null.String
}

func (srv *Service) UpdatePost(ctx context.Context, user *core.User, id core.PostID, in *PostInput) (*PostItem, error) {
	if err := srv.IsAdmin(user); err != nil {
		return nil, err
	}

	post, err := srv.Post.Query().ID(id).One(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get post")
	}

	lot, err := srv.Lot.Query().ID(post.LotID).One(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get lot")
	}

	post.LotID = in.LotID
	post.Title = in.Title
	post.Text = in.Text
	post.ScheduledAt = in.ScheduledAt
	post.DisableWebPagePreview = in.DisableWebPagePreview

	item := &PostItem{
		Post:     post,
		LotName:  lot.Name,
		Username: lot.Username,
		JoinLink: lot.JoinLink,
		Avatar:   lot.Avatar,
	}

	if err := srv.Post.Update(ctx, post); err != nil {
		return nil, errors.Wrap(err, "update post")
	}

	return item, nil
}

func (srv *Service) DeletePost(ctx context.Context, user *core.User, id core.PostID) error {
	if err := srv.IsAdmin(user); err != nil {
		return err
	}

	return srv.Post.Delete(ctx, id)
}

type FullPost struct {
	Items []*PostItem
	Total int
}

func (srv *Service) newPostItem(post *core.Post, lot *core.Lot) *PostItem {
	return &PostItem{
		Post:     post,
		LotName:  lot.Name,
		Username: lot.Username,
		Avatar:   lot.Avatar,
		JoinLink: lot.JoinLink,
	}
}

func (srv *Service) newPostItemSlice(posts core.PostSlice, lots core.LotSlice) []*PostItem {
	items := make([]*PostItem, len(posts))
	for i, post := range posts {
		items[i] = srv.newPostItem(post, lots.Find(post.LotID))
	}
	return items
}

func (srv *Service) GetPosts(ctx context.Context, user *core.User, limit int, offset int) (*FullPost, error) {
	if err := srv.IsAdmin(user); err != nil {
		return nil, err
	}

	posts, err := srv.Post.Query().
		Limit(limit).
		Offset(offset).
		SortBy(core.PostFieldScheduledAt, store.SortTypeAsc).
		All(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get posts")
	}

	lots, err := srv.Lot.Query().ID(posts.LotIDs()...).All(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get lots")
	}

	return &FullPost{
		Items: srv.newPostItemSlice(posts, lots),
		Total: len(posts),
	}, nil
}
