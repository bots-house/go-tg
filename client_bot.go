package tg

// This file contains bot related methods.

import (
	"context"

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