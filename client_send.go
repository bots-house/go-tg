package tg

import (
	"context"
	"encoding/json"
	"time"

	"github.com/pkg/errors"
)

// Peer define generic interface for Telegram chat ID.
type Peer interface {
	Peer() string
}

type ReplyMarkup interface {
	isReplyMarkup()
}

// TextOpts contains optional parameters for SendText method.
type TextOpts struct {
	// Mode for parsing entities in the message text.
	ParseMode string

	// Disables link previews for links in this message
	DisableWebPagePreview bool

	// Sends the message silently. Users will receive a notification with no sound.
	DisableNotification bool

	// If the message is a reply, ID of the original message
	ReplyToMessageID MessageID

	// InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply
	ReplyMarkup ReplyMarkup
}

// SendText message. Opts contains optional parameters.
func (client *Client) SendText(
	ctx context.Context,
	to Peer,
	text string,
	opts *TextOpts,
) (*Message, error) {
	r := NewRequest("sendMessage")

	r.SetPeer("chat_id", to)
	r.SetString("text", text)

	if opts != nil {
		r.SetOptString("parse_mode", opts.ParseMode)
		r.SetOptBool("disable_web_page_preview", opts.DisableWebPagePreview)
		r.SetOptBool("disable_notification", opts.DisableNotification)
		r.SetOptInt("reply_to_message_id", int(opts.ReplyToMessageID))
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

type PhotoOpts struct {
	// Photo caption (may also be used when resending photos by file_id), 0-1024 characters after entities parsing
	Caption string

	// Mode for parsing entities in the photo caption. See formatting options for more details.
	ParseMode string

	// Sends the message silently. Users will receive a notification with no sound.
	DisableNotification bool

	// If the message is a reply, ID of the original message
	ReplyToMessageID MessageID

	// InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply
	ReplyMarkup ReplyMarkup
}

// SendPhoto message. Opts contains optional parameters.
func (client *Client) SendPhoto(
	ctx context.Context,
	to Peer,
	photo *InputFile,
	opts *PhotoOpts,
) (*Message, error) {
	r := NewRequest("sendPhoto")

	r.SetPeer("chat_id", to)
	r.SetInputFile("photo", photo)

	if opts != nil {
		r.SetOptString("caption", opts.Caption)
		r.SetOptString("parse_mode", opts.ParseMode)
		r.SetOptBool("disable_notification", opts.DisableNotification)
		r.SetOptInt("reply_to_message_id", int(opts.ReplyToMessageID))
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

type AudioOpts struct {
	// Optional. Caption of the audio to be sent, 0-1024 characters after entities parsing
	Caption string `json:"caption,omitempty"`

	// Optional. Mode for parsing entities in the audio caption.
	ParseMode string `json:"parse_mode,omitempty"`

	// Optional. Duration if the audio in seconds.
	Duration Duration

	// Optional. Performer.
	Performer string

	// Optional. Title.
	Title string

	// Optional. Thumbnail of the file sent. Can be ignored if thumbnail generation
	// for the file is supported server-side.
	// The thumbnail should be in JPEG format and less than 200 kB in size.
	// A thumbnail‘s width and height should not exceed 320.
	// Ignored if the file is not uploaded using multipart/form-data.
	// Thumbnails can’t be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>”
	// if the thumbnail was uploaded using multipart/form-data under <file_attach_name>.
	Thumb *InputFile `json:"thumb"`

	// Sends the message silently. Users will receive a notification with no sound.
	DisableNotification bool

	// If the message is a reply, ID of the original message
	ReplyToMessageID MessageID

	// InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply
	ReplyMarkup ReplyMarkup
}

// SendAudio message. Opts contains optional parameters.
func (client *Client) SendAudio(
	ctx context.Context,
	to Peer,
	audio *InputFile,
	opts *AudioOpts,
) (*Message, error) {
	r := NewRequest("sendAudio")

	r.SetPeer("chat_id", to)
	r.SetInputFile("audio", audio)

	if opts != nil {
		r.SetOptString("caption", opts.Caption)
		r.SetOptString("parse_mode", opts.ParseMode)
		r.SetOptBool("disable_notification", opts.DisableNotification)
		r.SetOptInt("reply_to_message_id", int(opts.ReplyToMessageID))
		if err := r.SetOptJSON("reply_markup", opts.ReplyMarkup); err != nil {
			return nil, errors.Wrap(err, "marshal reply markup")
		}
		r.SetOptInt("duration", int(opts.Duration))
		r.SetOptString("performer", opts.Performer)
		r.SetOptString("title", opts.Title)
		if opts.Thumb != nil {
			r.SetInputFile("thumb", opts.Thumb)
		}
	}
	result := &Message{}

	if err := client.Invoke(ctx, r, result); err != nil {
		return nil, errors.Wrap(err, "invoke")
	}
	return result, nil
}

type VideoOpts struct {
	// Optional. Caption of the video to be sent, 0-1024 characters after entities parsing
	Caption string `json:"caption,omitempty"`

	// Optional. Mode for parsing entities in the video caption.
	ParseMode string `json:"parse_mode,omitempty"`

	// Optional. Video width
	Width int `json:"width,omitempty"`

	// Optional. Video height
	Height int `json:"height,omitempty"`

	// Optional. Video duration
	Duration Duration `json:"duration,omitempty"`

	// Optional. Pass True, if the uploaded video is suitable for streaming
	SupportsStreaming bool `json:"supports_streaming,omitempty"`

	// Optional. Thumbnail of the file sent. Can be ignored if thumbnail generation
	// for the file is supported server-side.
	// The thumbnail should be in JPEG format and less than 200 kB in size.
	// A thumbnail‘s width and height should not exceed 320.
	// Ignored if the file is not uploaded using multipart/form-data.
	// Thumbnails can’t be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>”
	// if the thumbnail was uploaded using multipart/form-data under <file_attach_name>.
	Thumb *InputFile `json:"thumb"`

	// Sends the message silently. Users will receive a notification with no sound.
	DisableNotification bool

	// If the message is a reply, ID of the original message
	ReplyToMessageID MessageID

	// InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply
	ReplyMarkup ReplyMarkup
}

// SendVideo message. Opts contains optional parameters.
func (client *Client) SendVideo(
	ctx context.Context,
	to Peer,
	video *InputFile,
	opts *VideoOpts,
) (*Message, error) {
	r := NewRequest("sendVideo")

	r.SetPeer("chat_id", to)
	r.SetInputFile("video", video)

	if opts != nil {
		r.SetOptString("caption", opts.Caption)
		r.SetOptString("parse_mode", opts.ParseMode)
		r.SetOptBool("disable_notification", opts.DisableNotification)
		r.SetOptInt("reply_to_message_id", int(opts.ReplyToMessageID))
		if err := r.SetOptJSON("reply_markup", opts.ReplyMarkup); err != nil {
			return nil, errors.Wrap(err, "marshal reply markup")
		}
		r.SetOptInt("duration", int(opts.Duration))
		r.SetOptBool("supports_streaming", opts.SupportsStreaming)
		r.SetOptInt("width", opts.Width)
		r.SetOptInt("height", opts.Height)

		if opts.Thumb != nil {
			r.SetInputFile("thumb", opts.Thumb)
		}
	}
	result := &Message{}

	if err := client.Invoke(ctx, r, result); err != nil {
		return nil, errors.Wrap(err, "invoke")
	}
	return result, nil
}

type AnimationOpts struct {
	// Optional. Thumbnail of the file sent. Can be ignored if thumbnail generation
	// for the file is supported server-side.
	// The thumbnail should be in JPEG format and less than 200 kB in size.
	// A thumbnail‘s width and height should not exceed 320.
	// Ignored if the file is not uploaded using multipart/form-data.
	// Thumbnails can’t be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>”
	// if the thumbnail was uploaded using multipart/form-data under <file_attach_name>.
	Thumb *InputFile `json:"thumb"`

	// Optional. Caption of the animation to be sent, 0-1024 characters after entities parsing
	Caption string `json:"caption,omitempty"`

	// Optional. Mode for parsing entities in the animation caption.
	ParseMode string `json:"parse_mode,omitempty"`

	// Optional. animation width
	Width int `json:"width,omitempty"`

	// Optional. animation height
	Height int `json:"height,omitempty"`

	// Optional. animation duration
	Duration Duration `json:"duration,omitempty"`

	// Optional. Pass True, if the uploaded animation is suitable for streaming
	SupportsStreaming bool `json:"supports_streaming,omitempty"`

	// Sends the message silently. Users will receive a notification with no sound.
	DisableNotification bool

	// If the message is a reply, ID of the original message
	ReplyToMessageID MessageID

	// InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply
	ReplyMarkup ReplyMarkup
}

// SendAnimation message. Opts contains optional parameters.
func (client *Client) SendAnimation(
	ctx context.Context,
	to Peer,
	animation *InputFile,
	opts *AnimationOpts,
) (*Message, error) {
	r := NewRequest("sendAnimation")

	r.SetPeer("chat_id", to)
	r.SetInputFile("animation", animation)

	if opts != nil {
		r.SetOptString("caption", opts.Caption)
		r.SetOptString("parse_mode", opts.ParseMode)
		r.SetOptBool("disable_notification", opts.DisableNotification)
		r.SetOptInt("reply_to_message_id", int(opts.ReplyToMessageID))
		if err := r.SetOptJSON("reply_markup", opts.ReplyMarkup); err != nil {
			return nil, errors.Wrap(err, "marshal reply markup")
		}
		r.SetOptInt("duration", int(opts.Duration))
		r.SetOptBool("supports_streaming", opts.SupportsStreaming)
		r.SetOptInt("width", opts.Width)
		r.SetOptInt("height", opts.Height)

		if opts.Thumb != nil {
			r.SetInputFile("thumb", opts.Thumb)
		}
	}
	result := &Message{}

	if err := client.Invoke(ctx, r, result); err != nil {
		return nil, errors.Wrap(err, "invoke")
	}
	return result, nil
}

type DocumentOpts struct {

	// Optional. Caption of the document to be sent, 0-1024 characters after entities parsing
	Caption string `json:"caption,omitempty"`

	// Optional. Mode for parsing entities in the document caption.
	ParseMode string `json:"parse_mode,omitempty"`

	// Optional. Thumbnail of the file sent. Can be ignored if thumbnail generation
	// for the file is supported server-side.
	// The thumbnail should be in JPEG format and less than 200 kB in size.
	// A thumbnail‘s width and height should not exceed 320.
	// Ignored if the file is not uploaded using multipart/form-data.
	// Thumbnails can’t be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>”
	// if the thumbnail was uploaded using multipart/form-data under <file_attach_name>.
	Thumb *InputFile `json:"thumb"`

	// Sends the message silently. Users will receive a notification with no sound.
	DisableNotification bool

	// If the message is a reply, ID of the original message
	ReplyToMessageID MessageID

	// InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply
	ReplyMarkup ReplyMarkup
}

// SendDocument message. Opts contains optional parameters.
func (client *Client) SendDocument(
	ctx context.Context,
	to Peer,
	document *InputFile,
	opts *DocumentOpts,
) (*Message, error) {
	r := NewRequest("sendDocument")

	r.SetPeer("chat_id", to)
	r.SetInputFile("document", document)

	if opts != nil {
		r.SetOptString("caption", opts.Caption)
		r.SetOptString("parse_mode", opts.ParseMode)
		r.SetOptBool("disable_notification", opts.DisableNotification)
		r.SetOptInt("reply_to_message_id", int(opts.ReplyToMessageID))
		if err := r.SetOptJSON("reply_markup", opts.ReplyMarkup); err != nil {
			return nil, errors.Wrap(err, "marshal reply markup")
		}
		if opts.Thumb != nil {
			r.SetInputFile("thumb", opts.Thumb)
		}
	}
	result := &Message{}

	if err := client.Invoke(ctx, r, result); err != nil {
		return nil, errors.Wrap(err, "invoke")
	}
	return result, nil
}

type LocationOpts struct {
	// Period in seconds for which the location will be updated (see Live Locations, should be between 60 and 86400.
	LivePeriod time.Duration

	// Sends the message silently. Users will receive a notification with no sound.
	DisableNotification bool

	// If the message is a reply, ID of the original message
	ReplyToMessageID MessageID

	// InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply
	ReplyMarkup ReplyMarkup
}

const (
	MinLiveLocationPeriod = time.Second * 60
	MaxLiveLocationPeriod = time.Hour * 24
)

// SendLocation message. Opts contains optional parameters.
func (client *Client) SendLocation(
	ctx context.Context,
	to Peer,
	location Location,
	opts *LocationOpts,
) (*Message, error) {
	r := NewRequest("sendLocation")

	r.SetPeer("chat_id", to)
	r.SetFloat64("latitude", location.Latitude)
	r.SetFloat64("longitude", location.Longitude)

	if opts != nil {
		r.SetOptDuration("live_period", opts.LivePeriod)
		r.SetOptBool("disable_notification", opts.DisableNotification)
		r.SetOptInt("reply_to_message_id", int(opts.ReplyToMessageID))
		if err := r.SetOptJSON("reply_markup", opts.ReplyMarkup); err != nil {
			return nil, errors.Wrap(err, "marshal reply markup")
		}
	}

	result := &Message{}

	if err := client.Invoke(ctx, r, result); err != nil {
		return nil, errors.Wrap(err, "invoke")
	}

	result.client = client

	return result, nil
}

type InputMedia interface {
	isInputMedia()
	json.Marshaler
}

type inputMedia struct {
	Type string `json:"type"`
	*InputMediaVideo
	*InputMediaPhoto
	*InputMediaAudio
	*InputMediaDocument
	*InputMediaAnimation
}

// InputMediaAudio a photo to be sent.
type InputMediaAudio struct {
	// File to send
	Media *InputFile

	// Optional. Caption of the audio to be sent, 0-1024 characters after entities parsing
	Caption string `json:"caption,omitempty"`

	// Optional. Mode for parsing entities in the audio caption.
	ParseMode string `json:"parse_mode,omitempty"`

	// Optional. Duration if the audio in seconds.
	Duration Duration

	// Optional. Performer.
	Performer string

	// Optional. Title.
	Title string

	// Optional. Thumbnail of the file sent. Can be ignored if thumbnail generation
	// for the file is supported server-side.
	// The thumbnail should be in JPEG format and less than 200 kB in size.
	// A thumbnail‘s width and height should not exceed 320.
	// Ignored if the file is not uploaded using multipart/form-data.
	// Thumbnails can’t be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>”
	// if the thumbnail was uploaded using multipart/form-data under <file_attach_name>.
	Thumb *InputFile `json:"thumb"`
}

func (ima InputMediaAudio) isInputMedia() {}

func (ima InputMediaAudio) MarshalJSON() ([]byte, error) {
	return json.Marshal(inputMedia{
		Type:            "audio",
		InputMediaAudio: &ima,
	})
}

// InputMediaDocument a document to be sent.
type InputMediaDocument struct {
	// File to send
	Media *InputFile

	// Optional. Caption of the document to be sent, 0-1024 characters after entities parsing
	Caption string `json:"caption,omitempty"`

	// Optional. Mode for parsing entities in the document caption.
	ParseMode string `json:"parse_mode,omitempty"`

	// Optional. Thumbnail of the file sent. Can be ignored if thumbnail generation
	// for the file is supported server-side.
	// The thumbnail should be in JPEG format and less than 200 kB in size.
	// A thumbnail‘s width and height should not exceed 320.
	// Ignored if the file is not uploaded using multipart/form-data.
	// Thumbnails can’t be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>”
	// if the thumbnail was uploaded using multipart/form-data under <file_attach_name>.
	Thumb *InputFile `json:"thumb"`
}

func (imd InputMediaDocument) isInputMedia() {}

func (imd InputMediaDocument) MarshalJSON() ([]byte, error) {
	return json.Marshal(inputMedia{
		Type:               "document",
		InputMediaDocument: &imd,
	})
}

// InputMediaPhoto a photo to be sent.
type InputMediaPhoto struct {
	// File to send
	Media *InputFile `json:"media"`

	// Optional. Caption of the photo to be sent, 0-1024 characters after entities parsing
	Caption string `json:"caption,omitempty"`

	// Optional. Mode for parsing entities in the photo caption.
	ParseMode string `json:"parse_mode,omitempty"`
}

func (imp InputMediaPhoto) isInputMedia() {}

func (imp InputMediaPhoto) MarshalJSON() ([]byte, error) {
	return json.Marshal(inputMedia{
		Type:            "photo",
		InputMediaPhoto: &imp,
	})
}

// InputMediaVideo a video to be sent.
type InputMediaVideo struct {
	// File to send.
	Media *InputFile `json:"media"`

	// Optional. Thumbnail of the file sent. Can be ignored if thumbnail generation
	// for the file is supported server-side.
	// The thumbnail should be in JPEG format and less than 200 kB in size.
	// A thumbnail‘s width and height should not exceed 320.
	// Ignored if the file is not uploaded using multipart/form-data.
	// Thumbnails can’t be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>”
	// if the thumbnail was uploaded using multipart/form-data under <file_attach_name>.
	Thumb *InputFile `json:"thumb"`

	// Optional. Caption of the video to be sent, 0-1024 characters after entities parsing
	Caption string `json:"caption,omitempty"`

	// Optional. Mode for parsing entities in the video caption.
	ParseMode string `json:"parse_mode,omitempty"`

	// Optional. Video width
	Width int `json:"width,omitempty"`

	// Optional. Video height
	Height int `json:"height,omitempty"`

	// Optional. Video duration
	Duration Duration `json:"duration,omitempty"`

	// Optional. Pass True, if the uploaded video is suitable for streaming
	SupportsStreaming bool `json:"supports_streaming,omitempty"`
}

func (imv InputMediaVideo) isInputMedia() {}

func (imv InputMediaVideo) MarshalJSON() ([]byte, error) {
	im := inputMedia{
		Type:            "video",
		InputMediaVideo: &imv,
	}

	return json.Marshal(im)
}

// InputMediaAnimation a animation to be sent.
type InputMediaAnimation struct {
	// File to send.
	Media *InputFile `json:"media"`

	// Optional. Thumbnail of the file sent. Can be ignored if thumbnail generation
	// for the file is supported server-side.
	// The thumbnail should be in JPEG format and less than 200 kB in size.
	// A thumbnail‘s width and height should not exceed 320.
	// Ignored if the file is not uploaded using multipart/form-data.
	// Thumbnails can’t be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>”
	// if the thumbnail was uploaded using multipart/form-data under <file_attach_name>.
	Thumb *InputFile `json:"thumb"`

	// Optional. Caption of the animation to be sent, 0-1024 characters after entities parsing
	Caption string `json:"caption,omitempty"`

	// Optional. Mode for parsing entities in the animation caption.
	ParseMode string `json:"parse_mode,omitempty"`

	// Optional. animation width
	Width int `json:"width,omitempty"`

	// Optional. animation height
	Height int `json:"height,omitempty"`

	// Optional. animation duration
	Duration Duration `json:"duration,omitempty"`

	// Optional. Pass True, if the uploaded animation is suitable for streaming
	SupportsStreaming bool `json:"supports_streaming,omitempty"`
}

func (ima InputMediaAnimation) isInputMedia() {}

func (ima InputMediaAnimation) MarshalJSON() ([]byte, error) {
	im := inputMedia{
		Type:                "animation",
		InputMediaAnimation: &ima,
	}

	return json.Marshal(im)
}

type MediaGroupOpts struct {
	// Sends the messages silently. Users will receive a notification with no sound.
	DisableNotification bool

	// If the messages are a reply, ID of the original message
	ReplyToMessageID MessageID
}

func (client *Client) SendMediaGroup(
	ctx context.Context,
	to Peer,
	group []InputMedia,
	opts *MediaGroupOpts,
) ([]*Message, error) {
	r := NewRequest("sendMediaGroup")

	r.SetPeer("chat_id", to)

	addInputMedia := func(file *InputFile) {
		addr := r.SetInputFile("", file)
		file.setAddr(addr)
	}

	for _, item := range group {
		switch v := item.(type) {
		case InputMediaVideo:
			addInputMedia(v.Media)

			if v.Thumb != nil {
				addInputMedia(v.Thumb)
			}
		case InputMediaPhoto:
			addInputMedia(v.Media)
		case InputMediaAudio:
			addInputMedia(v.Media)

			if v.Thumb != nil {
				addInputMedia(v.Thumb)
			}
		case InputMediaAnimation:
			addInputMedia(v.Media)

			if v.Thumb != nil {
				addInputMedia(v.Thumb)
			}
		case InputMediaDocument:
			addInputMedia(v.Media)

			if v.Thumb != nil {
				addInputMedia(v.Thumb)
			}
		default:
			panic("unexpected type when sendMediaGroup")
		}
	}

	if err := r.SetJSON("media", group); err != nil {
		return nil, errors.Wrap(err, "marshal media")
	}

	if opts != nil {
		r.SetOptBool("disable_notification", opts.DisableNotification)
		r.SetOptInt("reply_to_message_id", int(opts.ReplyToMessageID))
	}

	result := make([]*Message, len(group))

	if err := client.Invoke(ctx, r, result); err != nil {
		return nil, errors.Wrap(err, "invoke")
	}

	return result, nil
}
