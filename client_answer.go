package tg

import (
	"context"

	"github.com/pkg/errors"
)

type AnswerCallbackQueryOpts struct {
	// Optional. Text of the notification. If not specified, nothing will be shown to the user, 0-200 characters
	Text string

	// Optional. If true, an alert will be shown by the client instead of a notification at the top of the chat screen. Defaults to false.
	ShowAlert bool

	// Optional. URL that will be opened by the user's client.
	URL string

	// The maximum amount of time in seconds that the result of the callback query may be cached client-side. Telegram apps will support caching starting in version 3.14. Defaults to 0.
	CacheTime int
}

func (client *Client) AnswerCallbackQuery(
	ctx context.Context,
	callbackQueryID string,
	opts *AnswerCallbackQueryOpts,
) (bool, error) {
	r := NewRequest("answerCallbackQuery")

	r.SetString("callback_query_id", callbackQueryID)

	if opts != nil {
		r.SetOptString("text", opts.Text)
		r.SetOptBool("show_alert", opts.ShowAlert)
		r.SetOptString("url", opts.URL)
		r.SetOptInt("cache_time", opts.CacheTime)
	}

	var success bool

	if err := client.Invoke(ctx, r, &success); err != nil {
		return false, errors.Wrap(err, "invoke")
	}
	return success, nil
}
