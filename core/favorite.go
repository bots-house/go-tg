package core

import (
	"context"
	"time"

	"github.com/bots-house/birzzha/store"
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

type FavoriteField int8

const (
	FavoriteFieldCreatedAt FavoriteField = iota + 1
)

var (
	stringToFavoriteField = map[string]FavoriteField{
		"created_at": FavoriteFieldCreatedAt,
	}

	favoriteFieldToString = mirrorStringToFavoriteField(stringToFavoriteField)
)

func mirrorStringToFavoriteField(in map[string]FavoriteField) map[FavoriteField]string {
	result := make(map[FavoriteField]string, len(in))

	for k, v := range in {
		result[v] = k
	}
	return result
}

var ErrInvalidFavoriteField = NewError("invalid_favorite_field", "invalid favorite field")

func ParseFavoriteField(v string) (FavoriteField, error) {
	f, ok := stringToFavoriteField[v]
	if !ok {
		return FavoriteField(-1), ErrInvalidFavoriteField
	}
	return f, nil
}

func (ff FavoriteField) String() string {
	return favoriteFieldToString[ff]
}

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
	SortBy(field FavoriteField, typ store.SortType) FavoriteStoreQuery
	Delete(ctx context.Context) error
	All(ctx context.Context) (FavoriteSlice, error)
	One(ctx context.Context) (*Favorite, error)
}

var ErrFavoriteNotFound = NewError("favorite_not_found", "favorite not found")

type FavoriteStore interface {
	Add(ctx context.Context, favorite *Favorite) error
	Query() FavoriteStoreQuery
}
