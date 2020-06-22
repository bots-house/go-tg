package core

import (
	"context"
	"errors"
	"time"

	"github.com/volatiletech/null/v8"
)

type JoinedFrom string

const (
	JoinedFromSite = JoinedFrom("site")
	JoinedFromBot  = JoinedFrom("bot")
)

// Unique User ID in Birzzha.
type UserID int

// User represents Birzzha user.
type User struct {
	// ID of user in Birzzha. Unique
	ID UserID

	// User info from Telegram
	Telegram UserTelegram

	// First name.
	FirstName string

	// Last name.
	LastName null.String

	// If, true we don't sync data with Telegram
	IsNameEdited bool

	// Path to avatar in file store.
	Avatar null.String

	// True, if user is admin.
	IsAdmin bool

	// Users can sign up from bot or site.
	JoinedFrom JoinedFrom

	// Joined At
	JoinedAt time.Time

	// Updated at
	UpdatedAt null.Time
}

// UserTelegram contains Telegram user identities.
type UserTelegram struct {
	// ID of user in Telegram
	ID int

	// Username of user in Telegram
	Username null.String

	// LanguageCode from Telegram.
	LanguageCode null.String
}

var (
	ErrUserNotFound = errors.New("user not found")
)

// UserStoreQuery define complex ops with store.
type UserStoreQuery interface {
	// Filter by User.ID
	ID(id UserID) UserStoreQuery

	// Filter by User.Telegram.ID
	TelegramID(id int) UserStoreQuery

	// Query only one item from store.
	One(ctx context.Context) (*User, error)
}

// UserStore persistance interface for user.
type UserStore interface {
	// Add user to store. Also updates ID.
	Add(ctx context.Context, user *User) error

	// Update user in store.
	Update(ctx context.Context, user *User) error

	// Complex queries.
	Query() UserStoreQuery
}
