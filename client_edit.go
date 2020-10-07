package tg

import (
	"context"

	"github.com/pkg/errors"
)

// EditLiveLocation this method edit live location messages.
// A location can be edited until its LivePeriod expires
// or editing is explicitly disabled by a call to StopMessageLiveLocation.
// On success, if the edited message was sent by the bot,
// the edited Message is returned, otherwise True is returned.
//
// Only ReplyMarkup is usable from options.
func (client *Client) EditLiveLocation(
	ctx context.Context,
	chat Peer,
	msg MessageID,
	location Location,
	opts *LocationOpts,
) (*Message, error) {
	r := NewRequest("editMessageLiveLocation")

	r.SetPeer("chat_id", chat)
	r.SetInt("message_id", int(msg))
	r.SetFloat64("latitude", location.Latitude)
	r.SetFloat64("longitude", location.Longitude)

	if opts != nil {
		if err := r.SetOptJSON("reply_markup", opts.ReplyMarkup); err != nil {
			return nil, errors.Wrap(err, "marshal reply markup")
		}
	}

	result := &Message{}

	if err := client.Invoke(ctx, r, result); err != nil {
		return nil, errors.Wrap(err, "invoke")
	}

	return result, nil
}