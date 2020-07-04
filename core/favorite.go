package core

import (
	"context"
	"time"
)

type FavoriteID int

type Favorite struct {
	// Unique ID of favorite.
	ID FavoriteID

	// Reference to lot.
	LotID LotID

	// Reference to user.
	UserID UserID

	// Time when favorite was created.
	CreatedAt time.Time
}

type FavoriteSlice []*Favorite

func (fs FavoriteSlice) HasLot(id LotID) bool {
	for _, favorite := range fs {
		if id == favorite.LotID {
			return true
		}
	}
	return false
}

func NewFavorite(lotID LotID, userID UserID) *Favorite {
	return &Favorite{
		LotID:     lotID,
		UserID:    userID,
		CreatedAt: time.Now(),
	}
}

type FavoriteStoreQuery interface {
	ID(ids ...FavoriteID) FavoriteStoreQuery
	LotID(ids ...LotID) FavoriteStoreQuery
	UserID(id UserID) FavoriteStoreQuery
	Delete(ctx context.Context) error
	All(ctx context.Context) (FavoriteSlice, error)
	One(ctx context.Context) (*Favorite, error)
}

var ErrFavoriteNotFound = NewError("favorite_not_found", "favorite not found")

type FavoriteStore interface {
	Add(ctx context.Context, favorite *Favorite) error
	Query() FavoriteStoreQuery
}
