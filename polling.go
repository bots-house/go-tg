package tg

import (
	"context"
	"log"
	"sync"
	"time"
)

type Logger interface {
	Printf(format string, args ...interface{})
}

type ErrorHandler func(ctx context.Context, update *Update, err error)

// Poller listen to updates using long polling.
type Poller struct {
	// Client represents Telegram Bot API client.
	Client *Client

	// Handler represents Update handler.
	Handler Handler

	// Options contains options using for polling.
	Options PollingOptions

	// ErrorLogger used for logging polling errors.
	ErrorLogger Logger

	// ErrorHandler called ever
	ErrorHandler ErrorHandler

	// HandlerTimeout define time to process update.
	HandlerTimeout time.Duration

	// Delay before next request if error happened.
	RetryAfter time.Duration

	wg sync.WaitGroup
}

func (poller *Poller) onPollingError(ctx context.Context, err error) {
	format := "polling error: %#v, retry after %s"
	args := []interface{}{
		err,
		poller.RetryAfter,
	}

	if ctx.Err() == context.Canceled {
		return
	}

	if poller.ErrorLogger == nil {
		log.Printf(format, args...)
	} else {
		poller.ErrorLogger.Printf(format, args...)
	}

}

func (poller *Poller) onHandlerError(ctx context.Context, update *Update, err error) {
	if poller.ErrorHandler != nil {
		poller.ErrorHandler(ctx, update, err)
	} else {
		log.Printf("handler error: %#v", err)
	}
}

func (poller *Poller) processUpdates(ctx context.Context, updates []*Update) {
	if poller.Handler == nil {
		return
	}

	for _, update := range updates {
		update := update

		poller.wg.Add(1)
		go func() {
			defer poller.wg.Done()

			handlerCtx := ctx

			if poller.HandlerTimeout != 0 {
				var cancel context.CancelFunc
				handlerCtx, cancel = context.WithTimeout(ctx, poller.HandlerTimeout)
				defer cancel()
			}

			if err := poller.Handler.HandleUpdate(handlerCtx, update); err != nil {
				poller.onHandlerError(ctx, update, err)
			}
		}()
	}

}

func (poller *Poller) getNextOffset(updates []*Update) UpdateID {
	last := updates[len(updates)-1]
	return last.ID + 1
}

// Run polling until stopped.
func (poller *Poller) Run(ctx context.Context) {
	opts := poller.Options
	for {
		select {
		case <-ctx.Done():
			poller.wg.Wait()
			return
		default:
			updates, err := poller.Client.GetUpdates(ctx, &opts)
			if err != nil {
				poller.onPollingError(ctx, err)
				continue
			}

			if len(updates) > 0 {
				opts.Offset = poller.getNextOffset(updates)
				go poller.processUpdates(ctx, updates)
			}
		}
	}
}

// Polling updates.
func Polling(ctx context.Context, client *Client, handler Handler) {
	poller := Poller{
		Client:         client,
		Handler:        handler,
		HandlerTimeout: time.Second * 3,
	}

	poller.Run(ctx)
}
