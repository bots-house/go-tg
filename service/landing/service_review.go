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

func (srv *Service) GetReviews(ctx context.Context, offset int, limit int) (*ReviewList, error) {
	reviews, err := srv.Review.Query().Limit(limit).Offset(offset).OrderByCreatedAt().All(ctx)
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
