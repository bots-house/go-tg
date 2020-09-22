package tg

import (
	"context"
)

// Handler responds to Telegram Bot API update.
type Handler interface {
	HandleUpdate(ctx context.Context, update *Update) error
}

// The HandlerFunc type is an adapter to allow the use of
// ordinary functions as bot handlers. If f is a function
// with the appropriate signature, HandlerFunc(f) is a
// Handler that calls f.
type HandlerFunc func(ctx context.Context, update *Update) error

// HandleUpdate calls handler(ctx, update).
func (handler HandlerFunc) HandleUpdate(ctx context.Context, update *Update) error {
	return handler(ctx, update)
}
