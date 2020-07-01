package core

import (
	"context"
	"errors"
	"time"

	"github.com/volatiletech/null/v8"
)

type ReviewID int

type Review struct {
	// ID of review in Birzzha. Unique
	ID ReviewID

	// User info from Telegram.
	User ReviewUser

	// Review text.
	Text string

	// Created at.
	CreatedAt time.Time
}

type ReviewSlice []*Review

func NewReview(
	user ReviewUser,
	text string,
	createdAt time.Time,
) *Review {
	return &Review{
		User:      user,
		Text:      text,
		CreatedAt: createdAt,
	}
}

type ReviewUser struct {
	// External Telegram ID.
	TelegramID null.Int

	// First name.
	FirstName string

	// Last name.
	LastName null.String

	// Username of user in Telegram.
	Username null.String

	// Path to avatar in fule store.
	Avatar null.String
}

func NewReviewUser(
	telegramID null.Int,
	firstName string,
	lastName null.String,
	username null.String,
	avatar null.String,
) ReviewUser {
	return ReviewUser{
		TelegramID: telegramID,
		FirstName:  firstName,
		LastName:   lastName,
		Username:   username,
		Avatar:     avatar,
	}
}

type ReviewStoreQuery interface {
	Offset(offset int) ReviewStoreQuery
	Limit(limit int) ReviewStoreQuery
	OrderByCreatedAt() ReviewStoreQuery
	All(ctx context.Context) (ReviewSlice, error)
	Count(ctx context.Context) (int, error)
}

var (
	ErrReviewNotFound = errors.New("review not found")
)

type ReviewStore interface {
	Add(ctx context.Context, review *Review) error
	Query() ReviewStoreQuery
}
