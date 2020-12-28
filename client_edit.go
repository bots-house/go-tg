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

func (client *Client) EditMessageText(
	ctx context.Context,
	chat Peer,
	msg,
	inline MessageID,
	text string,
	opts *TextOpts,
) (*Message, error) {
	r := NewRequest("editMessageText")

	r.SetPeer("chat_id", chat)
	r.SetOptInt("message_id", int(msg))
	r.SetOptInt("inline_message_id", int(inline))
	r.SetString("text", text)
	if opts != nil {
		r.SetOptString("parse_mode", opts.ParseMode)
		r.SetOptBool("disable_web_page_preview", opts.DisableWebPagePreview)
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

func (client *Client) EditMessageCaption(
	ctx context.Context,
	chat Peer,
	msg,
	inline MessageID,
	caption string,
	opts *TextOpts,
) (*Message, error) {
	r := NewRequest("editMessageCaption")

	r.SetPeer("chat_id", chat)
	r.SetOptInt("message_id", int(msg))
	r.SetOptInt("inline_message_id", int(inline))
	r.SetString("caption", caption)
	if opts != nil {
		r.SetOptString("parse_mode", opts.ParseMode)
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

func (client *Client) EditMessageMedia(
	ctx context.Context,
	chat Peer,
	msg,
	inline MessageID,
	media InputMedia,
	replyMarkup ReplyMarkup,
) (*Message, error) {
	r := NewRequest("editMessageMedia")

	r.SetPeer("chat_id", chat)
	r.SetOptInt("message_id", int(msg))
	r.SetOptInt("inline_message_id", int(inline))

	addInputMedia := func(file *InputFile) {
		addr := r.SetInputFile("", file)
		file.setAddr(addr)
	}

	switch v := media.(type) {
	case InputMediaVideo:
		addInputMedia(v.Media)

		if v.Thumb != nil {
			addInputMedia(v.Thumb)
		}
	case InputMediaPhoto:
		addInputMedia(v.Media)
	case InputMediaAudio:
		addInputMedia(v.Media)
	case InputMediaDocument:
		addInputMedia(v.Media)
	case InputMediaAnimation:
		addInputMedia(v.Media)
	default:
		panic("unexpected type when sendMediaGroup")
	}

	if err := r.SetJSON("media", media); err != nil {
		return nil, errors.Wrap(err, "marshal media")
	}

	if err := r.SetOptJSON("reply_markup", replyMarkup); err != nil {
		return nil, errors.Wrap(err, "marshal reply markup")
	}
	result := &Message{}

	if err := client.Invoke(ctx, r, result); err != nil {
		return nil, errors.Wrap(err, "invoke")
	}
	return result, nil
}

func (client *Client) EditMessageReplyMarkup(
	ctx context.Context,
	chat Peer,
	msg,
	inline MessageID,
	replyMarkup ReplyMarkup,
) (*Message, error) {
	r := NewRequest("editMessageReplyMarkup")

	r.SetPeer("chat_id", chat)
	r.SetOptInt("message_id", int(msg))
	r.SetOptInt("inline_message_id", int(inline))
	if err := r.SetOptJSON("reply_markup", replyMarkup); err != nil {
		return nil, errors.Wrap(err, "marshal reply markup")
	}

	result := &Message{}

	if err := client.Invoke(ctx, r, result); err != nil {
		return nil, errors.Wrap(err, "invoke")
	}

	return result, nil
}
