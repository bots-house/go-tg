package models

import (
	"github.com/bots-house/birzzha/api/gen/models"
	"github.com/bots-house/birzzha/core"
	"github.com/bots-house/birzzha/service/landing"
	"github.com/go-openapi/swag"
)

func newReviewUser(user core.ReviewUser) *models.ReviewUser {
	return &models.ReviewUser{
		FirstName: swag.String(user.FirstName),
		LastName:  nullStringToString(user.LastName),
		Username:  nullStringToString(user.Username),
		Avatar:    nullStringToString(user.Avatar),
	}
}

func newReview(review *core.Review) *models.Review {
	return &models.Review{
		ID:        swag.Int64(int64(review.ID)),
		User:      newReviewUser(review.User),
		Text:      swag.String(review.Text),
		CreatedAt: timeToUnix(review.CreatedAt),
	}
}

func newReviewSlice(reviews core.ReviewSlice) []*models.Review {
	result := make([]*models.Review, len(reviews))

	for i, review := range reviews {
		result[i] = newReview(review)
	}

	return result
}

func NewReviewList(in *landing.ReviewList) *models.ReviewList {
	return &models.ReviewList{
		Total: swag.Int64(int64(in.Total)),
		Items: newReviewSlice(in.Items),
	}
}
