package core

import "context"

// LotTopic it's relation between lots and topics
type LotTopic struct {
	LotID   LotID
	TopicID TopicID
}

type LotTopicStore interface {
	Set(ctx context.Context, lot LotID, topics []TopicID) error
	Get(ctx context.Context, lot LotID) (TopicSlice, error)
}
