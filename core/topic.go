package core

import (
	"context"
	"errors"
	"time"

	"github.com/gosimple/slug"
)

// TopicID it's unique topic of Birzzha.
type TopicID int

// Topic of lot.
type Topic struct {
	ID        TopicID
	Name      string
	Slug      string
	CreatedAt time.Time
}

type TopicSlice []*Topic

func (ts TopicSlice) IDs() []TopicID {
	result := make([]TopicID, len(ts))
	for i, v := range ts {
		result[i] = v.ID
	}
	return result
}

func NewTopic(name string) *Topic {
	return &Topic{
		Name:      name,
		Slug:      slug.Make(name),
		CreatedAt: time.Now(),
	}
}

type TopicStoreQuery interface {
	ID(ids ...TopicID) TopicStoreQuery
	One(ctx context.Context) (*Topic, error)
	All(ctx context.Context) (TopicSlice, error)
}

var (
	ErrTopicNotFound = errors.New("topic not found")
)

type TopicStore interface {
	Add(ctx context.Context, topic *Topic) error
	Update(ctx context.Context, topic *Topic) error
	Delete(ctx context.Context, id TopicID) error
	Query() TopicStoreQuery
}
