package models

import (
	"github.com/bots-house/birzzha/api/gen/models"
	"github.com/bots-house/birzzha/core"
	"github.com/bots-house/birzzha/pkg/storage"
	"github.com/bots-house/birzzha/service/landing"
	"github.com/go-openapi/swag"
)

func newReviewUser(s storage.Storage, user core.ReviewUser) *models.ReviewUser {
	var avatar string

	if user.Avatar.Valid {
		avatar = s.PublicURL(avatar)
	}

	return &models.ReviewUser{
		FirstName: swag.String(user.FirstName),
		LastName:  nullStringToString(user.LastName),
		Username:  nullStringToString(user.Username),
		Avatar:    swag.String(avatar),
	}
}

func newReview(s storage.Storage, review *core.Review) *models.Review {
	return &models.Review{
		ID:        swag.Int64(int64(review.ID)),
		User:      newReviewUser(s, review.User),
		Text:      swag.String(review.Text),
		CreatedAt: timeToUnix(review.CreatedAt),
	}
}

func newReviewSlice(s storage.Storage, reviews core.ReviewSlice) []*models.Review {
	result := make([]*models.Review, len(reviews))

	for i, review := range reviews {
		result[i] = newReview(s, review)
	}

	return result
}

func NewReviewList(s storage.Storage, in *landing.ReviewList) *models.ReviewList {
	return &models.ReviewList{
		Total: swag.Int64(int64(in.Total)),
		Items: newReviewSlice(s, in.Items),
	}
}
