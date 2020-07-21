package landing

import (
	"context"

	"github.com/bots-house/birzzha/core"
	"github.com/pkg/errors"
)

type ReviewList struct {
	Total int
	Items core.ReviewSlice
}

func (srv *Service) GetReviews(ctx context.Context, offset, limit int) (*ReviewList, error) {
	reviews, err := srv.Review.Query().Offset(offset).Limit(limit).All(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get reviews")
	}

	total, err := srv.Review.Query().Count(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get reviews count")
	}

	return &ReviewList{
		Total: total,
		Items: reviews,
	}, nil
}
