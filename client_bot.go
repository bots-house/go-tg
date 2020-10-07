package tg

// This file contains bot related methods.

import (
	"context"
	"time"

	"github.com/pkg/errors"
)

// GetMe returns basic information about the bot in form of a UserBot object.
func (client *Client) GetMe(ctx context.Context) (*UserBot, error) {
	bot := &UserBot{}

	req := NewRequest("getMe")

	if err := client.Invoke(ctx, req, bot); err != nil {
		return nil, errors.Wrap(err, "get me")
	}

	return bot, nil
}

// SetMyCommands method  change the list of the bot's commands.
// Returns True on success.
func (client *Client) SetMyCommands(ctx context.Context, cmds []BotCommand) error {
	r := NewRequest("setMyCommands")

	if err := r.SetJSON("commands", cmds); err != nil {
		return errors.Wrap(err, "marshal commands")
	}

	return client.invokeExceptedTrue(ctx, r)
}

// GetMyCommands returns current list of the bot's commands.
func (client *Client) GetMyCommands(ctx context.Context) ([]BotCommand, error) {
	r := NewRequest("getMyCommands")

	result := []BotCommand{}

	if err := client.Invoke(ctx, r, &result); err != nil {
		return nil, errors.Wrap(err, "invoke")
	}

	return result, nil
}

type PollingOptions struct {
	// Identifier of the first update to be returned.
	// Must be greater by one than the highest among the identifiers of previously received updates.
	// By default, updates starting with the earliest unconfirmed update are returned.
	// An update is considered confirmed as soon as GetUpdates is called with an offset higher than its UpdateID.
	// The negative offset can be specified to retrieve updates
	// starting from -offset update from the end of the updates queue. All previous updates will forgotten.
	Offset UpdateID

	// Limits the number of updates to be retrieved. Values between 1-100 are accepted. Defaults to 100.
	Limit int

	// Timeout in seconds for long polling.
	// Defaults to 0, i.e. usual short polling.
	// Should be positive, short polling should be used for testing purposes only.
	Timeout time.Duration

	// A list of the update types you want your bot to receive.
	// For example, specify []string{"message", "edited_channel_post", "callback_query"}
	// to only receive updates of these types.
	// See Update for a complete list of available update types.
	// Specify an empty list to receive all updates regardless of type (default).
	// If not specified, the previous setting will be used.
	// Please note that this parameter doesn't affect updates created before the call to the getUpdates,
	// so unwanted updates may be received for a short period of time.
	AllowedUpdates []string
}

// GetUpdates returns incoming updates using long polling.
func (client *Client) GetUpdates(ctx context.Context, opts *PollingOptions) ([]*Update, error) {
	r := NewRequest("getUpdates")

	limit := 100

	if opts != nil {
		limit = opts.Limit

		r.SetOptInt("offset", int(opts.Offset))
		r.SetOptInt("limit", opts.Limit)
		r.SetOptDuration("timeout", opts.Timeout)
		if err := r.SetJSON("allowed_updates", opts.AllowedUpdates); err != nil {
			return nil, errors.Wrap(err, "marshal AllowedUpdates")
		}
	}

	updates := make([]*Update, 0, limit)

	if err := client.Invoke(ctx, r, &updates); err != nil {
		return nil, errors.Wrap(err, "invoke")
	}

	for _, update := range updates {
		update.BindClient(client)
	}

	return updates, nil
}

// GetWebhookInfo returns current webhook status. Requires no parameters.
// If the bot is using getUpdates, will return an object with the url field empty.
func (client *Client) GetWebhookInfo(ctx context.Context) (*WebhookInfo, error) {
	r := NewRequest("getWebhookInfo")

	result := &WebhookInfo{}

	if err := client.Invoke(ctx, r, result); err != nil {
		return nil, errors.Wrap(err, "invoke")
	}

	return result, nil
}

type SetWebhookOptions struct {
	// HTTPS url to send updates to. Use an empty string to remove webhook integration.
	URL string

	// Maximum allowed number of simultaneous HTTPS connections to the webhook for update delivery, 1-100.
	// Defaults to 40. Use lower values to limit the load on your bot's server, and higher values to increase your bot's throughput.
	MaxConnections int

	// A list of the update types you want your bot to receive.
	// For example, specify []string{"message", "edited_channel_post", "callback_query"}
	// to only receive updates of these types.
	// See Update for a complete list of available update types.
	// Specify an empty list to receive all updates regardless of type (default).
	// If not specified, the previous setting will be used.
	// Please note that this parameter doesn't affect updates created before the call to the getUpdates,
	// so unwanted updates may be received for a short period of time.
	AllowedUpdates []string
}

// SetWebhook specify a url and receive incoming updates via an outgoing webhook.
func (client *Client) SetWebhook(ctx context.Context, opts *SetWebhookOptions) error {
	r := NewRequest("setWebhook")

	if opts != nil {
		r.SetOptString("url", opts.URL)
		r.SetOptInt("max_connections", opts.MaxConnections)
		if err := r.SetJSON("allowed_updates", opts.AllowedUpdates); err != nil {
			return errors.Wrap(err, "marshal AllowedUpdates")
		}
	}

	return client.invokeExceptedTrue(ctx, r)
}

type ChatOptions struct {
	ChatID   ChatID
	Username string
}

func (client *Client) GetChat(ctx context.Context, opts *ChatOptions) (*Chat, error) {
	r := NewRequest("getChat")
	result := &Chat{}

	if opts != nil {
		if opts.ChatID != 0 {
			r.SetOptInt64("chat_id", int64(opts.ChatID))
		} else {
			r.SetOptString("chat_id", opts.Username)
		}
	}
	if err := client.Invoke(ctx, r, result); err != nil {
		return nil, errors.Wrap(err, "invloke")
	}

	return result, nil
}
