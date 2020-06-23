package tg

import (
	"context"

	"github.com/bots-house/birzzha/core"
)

type Channel struct {
	// ID of channel (bot api format)
	ID int64

	// Channel name (title)
	Name string

	// Avatar of channel (optional)
	Avatar string

	// Members count
	MembersCount int

	// @username (optional)
	Username string

	// Description of channel
	Description string

	// Average views count in channel
	DailyCoverage int
}

// Group represents telegram group or supergroups
type Group struct {
	// ID of chat (bot api format)
	ID int64

	// Chat name (title)
	Name string

	// Avatar of chat (optional)
	Avatar string

	// Members count
	MembersCount int

	// @username (optional)
	Username string

	// Description of chat
	Description string
}

// ResolveResult is union of all possible result of Resolver.Resolve()
type ResolveResult struct {
	Channel *Channel
	Group   *Group
}

var (
	ErrInvalidQuery = core.NewError(
		"invalid_query",
		"query is invalid format",
	)
)

type Resolver interface {
	Resolve(ctx context.Context, input string) (*ResolveResult, error)
}
