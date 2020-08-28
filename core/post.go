package core

import (
	"context"
	"time"

	"github.com/bots-house/birzzha/store"
	"github.com/volatiletech/null/v8"
)

// PostID it's unique post id
type PostID int

// Post it's scheduled publications to Telegram channel.
type Post struct {
	// Unique ID of post
	ID PostID

	// Related lot ID
	LotID LotID

	// Text of post
	Text string

	// Title of post
	Title null.String

	// Inline buttons of post
	Buttons PostButtons

	// Disable web page preview
	DisableWebPagePreview bool

	// Time when post should be published
	ScheduledAt time.Time

	// Time when post was posted
	PublishedAt null.Time
}

func NewPost(
	lotID LotID,
	text string,
	title null.String,
	disableWebPagePreview bool,
	scheduledAt time.Time,
	lotLinkButton bool,
) *Post {
	return &Post{
		LotID:                 lotID,
		Text:                  text,
		Title:                 title,
		DisableWebPagePreview: disableWebPagePreview,
		ScheduledAt:           scheduledAt,
		Buttons:               PostButtons{Like: true, LotLink: lotLinkButton},
	}
}

type PostSlice []*Post

func (ps PostSlice) LotIDs() []LotID {
	ids := make([]LotID, len(ps))
	for i, post := range ps {
		ids[i] = post.LotID
	}
	return ids
}

var (
	ErrPostNotFound = NewError("post_not_found", "post not found")
)

type PostButtons struct {
	Like    bool
	LotLink bool
}

type PostField int8

const (
	PostFieldScheduledAt PostField = iota + 1
)

var (
	stringToPostField = map[string]PostField{
		"scheduled_at": PostFieldScheduledAt,
	}

	postFieldToString = mirrorStringToPostField(stringToPostField)
)

func mirrorStringToPostField(in map[string]PostField) map[PostField]string {
	result := make(map[PostField]string, len(in))

	for k, v := range in {
		result[v] = k
	}

	return result
}

var ErrInvalidPostField = NewError("invalid_post_field", "invalid post field")

func ParsePostField(v string) (PostField, error) {
	f, ok := stringToPostField[v]
	if !ok {
		return PostField(-1), ErrInvalidPostField
	}
	return f, nil
}

func (pf PostField) String() string {
	return postFieldToString[pf]
}

type PostStore interface {
	// Add post to store
	Add(ctx context.Context, post *Post) error

	// Update post in store
	Update(ctx context.Context, post *Post) error

	// Delete post from store
	Delete(ctx context.Context, id PostID) error

	// Pull returns expired, not published scheduled posts
	Pull(ctx context.Context) (PostSlice, error)

	Query() PostStoreQuery
}

type PostStoreQuery interface {
	ID(ids ...PostID) PostStoreQuery
	Limit(limit int) PostStoreQuery
	Offset(offset int) PostStoreQuery
	SortBy(field PostField, typ store.SortType) PostStoreQuery
	One(ctx context.Context) (*Post, error)
	All(ctx context.Context) (PostSlice, error)
}
