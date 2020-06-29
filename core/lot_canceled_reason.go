package core

import (
	"context"
	"time"

	"github.com/volatiletech/null/v8"
)

type LotCanceledReasonID int

type LotCanceledReason struct {
	// ID of discontinued reason
	ID LotCanceledReasonID

	// Text of discontinued reason
	Why string

	// If true, reason was displayed when go by direct link.
	IsPublic bool

	// Time when canceled reason was created
	CreatedAt time.Time

	// Time when canceled reason was updated
	UpdatedAt null.Time
}

type LotCanceledReasonStoreQuery interface {
	One(ctx context.Context) (*LotCanceledReason, error)
	All(ctx context.Context) ([]*LotCanceledReason, error)
}

type LotCanceledReasonStore interface {
	Add(ctx context.Context, reason *LotCanceledReason) error
	Update(ctx context.Context, reason *LotCanceledReason) error
	Query() LotCanceledReasonStoreQuery
}
