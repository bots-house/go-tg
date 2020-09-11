package admin

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/bots-house/birzzha/core"
	"github.com/bots-house/birzzha/store"
	tgbotapi "github.com/bots-house/telegram-bot-api"
	"github.com/pkg/errors"
	"github.com/volatiletech/null/v8"
)

type PostInput struct {
	LotID                 core.LotID
	Text                  string
	Title                 null.String
	DisableWebPagePreview bool
	ScheduledAt           time.Time
	LotLinkButton         bool
}

var (
	ErrLotIsAlreadyPublished = core.NewError("lot_is_already_published", "lot is already published")
)

func (srv *Service) CreatePost(ctx context.Context, user *core.User, in *PostInput) (*core.Post, error) {
	if err := srv.IsAdmin(user); err != nil {
		return nil, err
	}
	post := core.NewPost(in.LotID, in.Text, in.Title, in.DisableWebPagePreview, in.ScheduledAt, core.PostStatusScheduled, in.LotLinkButton)

	if err := srv.Txier(ctx, func(ctx context.Context) error {

		if err := srv.Post.Add(ctx, post); err != nil {
			return errors.Wrap(err, "add post")
		}

		if in.LotID != 0 {
			lot, err := srv.Lot.Query().ID(in.LotID).One(ctx)
			if err != nil {
				return errors.Wrap(err, "get lot")
			}

			if lot.Status == core.LotStatusPublished {
				return ErrLotIsAlreadyPublished
			}

			lot.Status = core.LotStatusScheduled
			lot.ScheduledAt = null.TimeFrom(in.ScheduledAt)
			if err := srv.Lot.Update(ctx, lot); err != nil {
				return errors.Wrap(err, "update lot")
			}

			srv.Notify.SendUser(lot.OwnerID, userNotifyScheduledLot{
				Lot: lot,
			})
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return post, nil
}

type PostItemLot struct {
	Name     string
	Username null.String
	JoinLink null.String
	Avatar   null.String
}

type PostItem struct {
	*core.Post
	Lot *PostItemLot
}

func joinSitePath(site, path string) string {
	path = strings.TrimPrefix(path, "/")
	site = strings.TrimSuffix(site, "/")

	return site + "/" + path
}

func (srv *Service) UpdatePost(ctx context.Context, user *core.User, id core.PostID, in *PostInput) (*PostItem, error) {
	if err := srv.IsAdmin(user); err != nil {
		return nil, err
	}

	post, err := srv.Post.Query().ID(id).One(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get post")
	}

	settings, err := srv.Settings.Get(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get settings")
	}

	post.Title = in.Title
	post.ScheduledAt = in.ScheduledAt
	post.DisableWebPagePreview = in.DisableWebPagePreview
	post.Buttons.LotLink = in.LotLinkButton

	if post.MessageID.Valid {
		if post.Text != in.Text {
			post.Text = in.Text
			_, err := srv.TgClient.EditMessageText(tgbotapi.EditMessageTextConfig{
				BaseEdit: tgbotapi.BaseEdit{
					ChatID:    settings.Channel.PrivateID,
					MessageID: post.MessageID.Int,
				},
				Text:                  in.Text,
				DisableWebPagePreview: in.DisableWebPagePreview,
				ParseMode:             "HTML",
			})
			if err != nil {
				return nil, errors.Wrap(err, "edit message")
			}
		}
		msg := tgbotapi.EditMessageReplyMarkupConfig{
			BaseEdit: tgbotapi.BaseEdit{
				ChatID:    settings.Channel.PrivateID,
				MessageID: post.MessageID.Int,
			},
		}
		if in.LotLinkButton {
			if post.LotID != 0 {
				markup := tgbotapi.NewInlineKeyboardMarkup(
					tgbotapi.NewInlineKeyboardRow(
						tgbotapi.InlineKeyboardButton{
							Text: "Подробнее",
							LoginURL: &tgbotapi.LoginURL{
								URL: joinSitePath(srv.Config.Site, fmt.Sprintf("lots/%d?from=channel", post.LotID)),
							},
						},
					),
				)
				msg.BaseEdit.ReplyMarkup = &markup

			}
		}
		_, err := srv.TgClient.EditMessageText(msg)
		if err != nil {
			return nil, errors.Wrap(err, "edit message")
		}

	}
	item := &PostItem{
		Post: post,
	}

	if in.LotID != 0 {
		lot, err := srv.Lot.Query().ID(in.LotID).One(ctx)
		if err != nil {
			return nil, errors.Wrap(err, "get lot")
		}
		item.Lot = &PostItemLot{
			Name:     lot.Name,
			Username: lot.Username,
			JoinLink: lot.JoinLink,
			Avatar:   lot.Avatar,
		}

		post.LotID = in.LotID
		item.Post.LotID = in.LotID
	}

	if err := srv.Post.Update(ctx, post); err != nil {
		return nil, errors.Wrap(err, "update post")
	}

	return item, nil
}

func (srv *Service) DeletePost(ctx context.Context, user *core.User, id core.PostID, deleteFromChannel bool) error {
	if err := srv.IsAdmin(user); err != nil {
		return err
	}

	settings, err := srv.Settings.Get(ctx)
	if err != nil {
		return errors.Wrap(err, "get settings")
	}

	post, err := srv.Post.Query().ID(id).One(ctx)
	if err != nil {
		return errors.Wrap(err, "get post")
	}

	if deleteFromChannel {
		if post.MessageID.Valid {
			_, err := srv.TgClient.DeleteMessage(tgbotapi.DeleteMessageConfig{
				ChatID:    settings.Channel.PrivateID,
				MessageID: post.MessageID.Int,
			})
			if err != nil {
				return errors.Wrap(err, "delete post from channel")
			}
		}
	}

	return srv.Post.Delete(ctx, id)
}

type FullPost struct {
	Items []*PostItem
	Total int
}

func (srv *Service) newPostItem(post *core.Post, lot *core.Lot) *PostItem {
	item := &PostItem{
		Post: post,
	}
	if lot != nil {
		item.Lot = &PostItemLot{
			Name:     lot.Name,
			Username: lot.Username,
			JoinLink: lot.JoinLink,
			Avatar:   lot.Avatar,
		}
	}
	return item
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
