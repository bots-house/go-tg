package core

import (
	"context"
	"time"

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
	disableWebPagePreview bool,
	scheduledAt time.Time,
) *Post {
	return &Post{
		LotID:                 lotID,
		Text:                  text,
		DisableWebPagePreview: disableWebPagePreview,
		ScheduledAt:           scheduledAt,
		Buttons:               PostButtons{Like: true},
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
	Like bool
}

type PostStore interface {
	// Add post to store
	Add(ctx context.Context, post *Post) error

	// Update post in store
	Update(ctx context.Context, post *Post) error

	// Pull returns expired, not published scheduled posts
	Pull(ctx context.Context) (PostSlice, error)
}
