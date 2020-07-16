package core

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/volatiletech/null/v8"
)

type JoinedFrom string

const (
	JoinedFromSite = JoinedFrom("site")
	JoinedFromBot  = JoinedFrom("bot")
)

func (jf JoinedFrom) String() string {
	return string(jf)
}

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

type UserSlice []*User

// Name returns user full name.
func (user *User) Name() string {
	name := user.FirstName

	if user.LastName.Valid {
		name += " " + user.LastName.String
	}

	return name
}

func (user *User) TelegramLink() string {
	return "tg://user?id=" + strconv.Itoa(user.Telegram.ID)
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

func (ut *UserTelegram) TelegramLink() null.String {
	if ut.Username.Valid {
		return null.StringFrom(fmt.Sprintf("https://t.me/%s", ut.Username.String))
	}
	return null.String{}
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

	Limit(limit int) UserStoreQuery
	Offset(offset int) UserStoreQuery
	OrderByJoinedAt() UserStoreQuery

	// Query only one item from store.
	One(ctx context.Context) (*User, error)
	Count(ctx context.Context) (int, error)
	All(ctx context.Context) (UserSlice, error)
}

// UserStore persistence interface for user.
type UserStore interface {
	// Add user to store. Also updates ID.
	Add(ctx context.Context, user *User) error

	// Update user in store.
	Update(ctx context.Context, user *User) error

	// Complex queries.
	Query() UserStoreQuery
}
