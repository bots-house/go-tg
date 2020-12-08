package tg

import (
	"context"

	"github.com/pkg/errors"
)

// Use this method to delete a message, including service messages, with the following limitations:
// - A message can only be deleted if it was sent less than 48 hours ago.
// - A dice message in a private chat can only be deleted if it was sent more than 24 hours ago.
// - Bots can delete outgoing messages in private chats, groups, and supergroups.
// - Bots can delete incoming messages in private chats.
// - Bots granted can_post_messages permissions can delete outgoing messages in channels.
// - If the bot is an administrator of a group, it can delete any message there.
// - If the bot has can_delete_messages permission in a supergroup or a channel, it can delete any message there.
// Returns True on success.
func (client *Client) DeleteMessage(
	ctx context.Context,
	chat Peer,
	msg MessageID,
) (bool, error) {
	r := NewRequest("deleteMessage")

	var success bool

	r.SetPeer("chat_id", chat)
	r.SetInt("message_id", int(msg))

	if err := client.Invoke(ctx, r, &success); err != nil {
		return false, errors.Wrap(err, "invoke")
	}
	return success, nil
}
