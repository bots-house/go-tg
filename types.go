package tg

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// FileID it's ID of uploaded file.
type FileID string

// UserID it's unique user ID in Telegram.
type UserID int

// Peer implements Peer interface.
func (id UserID) Peer() string {
	return strconv.Itoa(int(id))
}

// Username it's Telegram username.
type Username string

func (un Username) Peer() string {
	return un.String()
}

func (un Username) String() string {
	return "@" + string(un)
}

// User represents Telegram user.
type User struct {
	// Unique identifier for this user or bot
	ID UserID `json:"id"`

	// True, if this user is a bot
	IsBot bool `json:"is_bot"`

	// User's or bot's first name
	FirstName string `json:"first_name"`

	// Optional. User's or bot's last name
	LastName string `json:"last_name,omitempty"`

	// Optional. User's or bot's username
	Username Username `json:"username,omitempty"`

	// Optional. IETF language tag of the user's language
	LanguageCode string `json:"language_code,omitempty"`
}

// UserBot represents Telegram bot user info.
type UserBot struct {
	User `json:"-"`

	// Optional. True, if the bot can be invited to groups.
	CanJoinGroups bool `json:"can_join_groups"`

	// Optional. True, if privacy mode is disabled for the bot.
	CanReadAllGroupMessages bool `json:"can_read_all_group_messages"`

	// Optional. True, if the bot supports inline queries.
	SupportsInlineQueries bool `json:"supports_inline_queries"`
}

func (userBot *UserBot) UnmarshalJSON(data []byte) error {
	// unmarshal user fields
	var user User

	if err := json.Unmarshal(data, &user); err != nil {
		return err
	}

	userBot.User = user

	// unmarshal bot fields
	var bot struct {
		CanJoinGroups           bool `json:"can_join_groups"`
		CanReadAllGroupMessages bool `json:"can_read_all_group_messages"`
		SupportsInlineQueries   bool `json:"supports_inline_queries"`
	}

	if err := json.Unmarshal(data, &bot); err != nil {
		return err
	}

	userBot.CanJoinGroups = bot.CanJoinGroups
	userBot.CanReadAllGroupMessages = bot.CanReadAllGroupMessages
	userBot.SupportsInlineQueries = bot.SupportsInlineQueries

	return nil
}

// BotCommand represents a bot command.
type BotCommand struct {
	// Text of the command, 1-32 characters.
	// Can contain only lowercase English letters, digits and underscores.
	Command string `json:"command"`

	// Description of the command, 3-256 characters.
	Description string `json:"description"`
}

// UpdateID the update‘s unique identifier.
// Update identifiers start from a certain positive number and increase sequentially.
// This ID becomes especially handy if you’re using Webhooks, since it allows you to ignore repeated updates or
// to restore the correct update sequence, should they get out of order.
// If there are no new updates for at least a week, then identifier of the next update will be chosen
// randomly instead of sequentially.
type UpdateID int

// Update represents an incoming update.
// At most one of the optional parameters can be present in any given update.
type Update struct {
	// Update unique ID
	ID UpdateID `json:"update_id"`

	// Optional. New incoming message of any kind — text, photo, sticker, etc.
	Message *Message `json:"message,omitempty"`

	// Optional. New version of a message that is known to the bot and was edited.
	EditedMessage *Message `json:"edited_message,omitempty"`

	// Optional. New incoming channel post of any kind — text, photo, sticker, etc.
	ChannelPost *Message `json:"channel_post,omitempty"`

	// Optional. New version of a channel post that is known to the bot and was edited.
	EditedChannelPost *Message `json:"edited_channel_post,omitempty"`

	// Optional. New incoming inline query
	InlineQuery json.RawMessage `json:"inline_query,omitempty"`

	// Optional. The result of an inline query that was chosen by a user and sent to their chat partner.
	ChosenInlineResult json.RawMessage `json:"chosen_inline_result,omitempty"`

	// Optional. New incoming callback query
	CallbackQuery *CallbackQuery `json:"callback_query,omitempty"`

	// Optional. New incoming shipping query. Only for invoices with flexible price
	ShippingQuery json.RawMessage `json:"shipping_query,omitempty"`

	// Optional. New incoming pre-checkout query. Contains full information about checkout
	PreCheckoutQuery json.RawMessage `json:"pre_checkout_query,omitempty"`

	// Optional. New poll state.
	// Bots receive only updates about stopped polls and polls, which are sent by the bot
	Poll *Poll `json:"poll,omitempty"`

	// Optional. A user changed their answer in a non-anonymous poll.
	// Bots receive new votes only in polls that were sent by the bot itself.
	PollAnswer *PollAnswer `json:"poll_answer,omitempty"`

	// Client contains client who received this update.
	// Use BindClient to propagate binding to child fields.
	Client *Client `json:"-"`
}

// BindClient set Update and child struct client.
func (update *Update) BindClient(client *Client) {
	update.Client = client

	if update.Message != nil {
		update.Message.client = client
	}

	if update.EditedMessage != nil {
		update.EditedMessage.client = client
	}

	if update.ChannelPost != nil {
		update.ChannelPost.client = client
	}

	if update.EditedChannelPost != nil {
		update.EditedChannelPost.client = client
	}
}

// UnixTime represents Unix timestamp.
type UnixTime int64

// Duration
type Duration int

// ChatID represents ID of Telegram chat.
type ChatID int64

func (id ChatID) Peer() string {
	return strconv.FormatInt(int64(id), 10)
}

// ChatPhoto object represents a chat photo.
type ChatPhoto struct {
	// File identifier of small (160x160) chat photo.
	// This FileID can be used only for photo download and only for as long as the photo is not changed.
	SmallFileID FileID `json:"small_file_id"`

	// Unique file identifier of small (160x160) chat photo,
	// which is supposed to be the same over time and for different bots.
	// Can't be used to download or reuse the file.
	SmallFileUniqueID FileID `json:"small_file_unique_id"`

	// File identifier of big (640x640) chat photo.
	// This file_id can be used only for photo download and only for as long as the photo is not changed.
	BigFileID FileID `json:"big_file_id"`

	// Unique file identifier of big (640x640) chat photo, which is supposed to be the same over time
	// and for different bots. Can't be used to download or reuse the file.
	BigFileUniqueID FileID `json:"big_file_unique_id"`
}

// PhotoSize object represents one size of a photo or a file / sticker thumbnail.
type PhotoSize struct {
	// Identifier for this file, which can be used to download or reuse the file
	FileID FileID `json:"file_id"`

	// Unique identifier for this file, which is supposed to be the same over time and for different bots.
	// Can't be used to download or reuse the file.
	FileUniqueID FileID `json:"file_unique_id"`

	// Photo width
	Width int `json:"width"`

	// Photo height
	Height int `json:"height"`

	// Optional. File size
	FileSize int `json:"file_size"`
}

type PhotoSizeSlice []PhotoSize

type UserProfilePhotos struct {
	// Total number of profile pictures the target user has.
	TotalCount int `json:"total_count,omitempty"`

	// Requested profile pictures (in up to 4 sizes each).
	Photos []PhotoSizeSlice `json:"photos,omitempty"`
}

// ChatPermissions describes actions that a non-administrator user is allowed to take in a chat.
type ChatPermissions struct {
	// Optional. True, if the user is allowed to send text messages, contacts, locations and venues
	CanSendMessages bool `json:"can_send_messages,omitempty"`

	// Optional. True, if the user is allowed to send audios, documents, photos, videos, video notes and voice notes,
	// implies can_send_messages
	CanSendMediaMessages bool `json:"can_send_media_messages,omitempty"`

	// Optional. True, if the user is allowed to send polls, implies can_send_messages
	CanSendPolls bool `json:"can_send_polls,omitempty"`

	// Optional. True, if the user is allowed to send animations, games, stickers and use inline bots,
	// implies can_send_media_messages
	CanSendOtherMessages bool `json:"can_send_other_messages,omitempty"`

	// Optional. True, if the user is allowed to add web page previews to their messages,
	// implies can_send_media_messages
	CanAddWebPagePreviews bool `json:"can_add_web_page_previews,omitempty"`

	// Optional. True, if the user is allowed to change the chat title, photo and other settings.
	// Ignored in public supergroups.
	CanChangeInfo bool `json:"can_change_info,omitempty"`

	// Optional. True, if the user is allowed to invite new users to the chat
	CanInviteUsers bool `json:"can_invite_users,omitempty"`

	// Optional. True, if the user is allowed to pin messages. Ignored in public supergroups.
	CanPinMessages bool `json:"can_pin_messages,omitempty"`
}

// Chat object represents a chat.
type Chat struct {
	// Unique identifier for this chat.
	ID ChatID `json:"id"`

	// Type of chat, can be either “private”, “group”, “supergroup” or “channel”
	Type string `json:"type"`

	// Optional. Title, for supergroups, channels and group chats
	Title string `json:"title,omitempty"`

	// Optional. Username, for private chats, supergroups and channels if available
	Username Username `json:"username,omitempty"`

	// Optional. First name of the other party in a private chat
	FirstName string `json:"first_name,omitempty"`

	// Optional. Last name of the other party in a private chat
	LastName string `json:"last_name,omitempty"`

	// Optional. Chat photo. Returned only in GetChat.
	Photo *ChatPhoto `json:"photo,omitempty"`

	// Optional. Description, for groups, supergroups and channel chats. Returned only in getChat.
	Description string `json:"description,omitempty"`

	// Optional. Chat invite link, for groups, supergroups and channel chats.
	// Each administrator in a chat generates their own invite links,
	// so the bot must first generate the link using exportChatInviteLink.
	//
	// Returned only in getChat.
	InviteLink string `json:"invite_link,omitempty"`

	// Optional. Pinned message, for groups, supergroups and channels. Returned only in getChat.
	PinnedMessage *Message `json:"pinned_message,omitempty"`

	// Optional. Default chat member permissions, for groups and supergroups. Returned only in getChat.
	Permissions *ChatPermissions `json:"permissions,omitempty"`

	// Optional. For supergroups, the minimum allowed delay between consecutive messages sent by each unpriviledged user.
	// Returned only in getChat.
	SlowModeDelay Duration `json:"slow_mode_delay,omitempty"`

	// Optional. For supergroups, name of group sticker set. Returned only in getChat.
	StickerSetName string `json:"sticker_set_name,omitempty"`

	// Optional. True, if the bot can change the group sticker set. Returned only in getChat.
	CanSetStickerSet bool `json:"can_set_sticker_set,omitempty"`
}

// IsPrivate returns if the Chat is a private conversation.
func (c *Chat) IsPrivate() bool {
	return c.Type == "private"
}

// IsGroup returns if the Chat is a group.
func (c Chat) IsGroup() bool {
	return c.Type == "group"
}

// IsSuperGroup returns if the Chat is a supergroup.
func (c Chat) IsSuperGroup() bool {
	return c.Type == "supergroup"
}

// IsChannel returns if the Chat is a channel.
func (c Chat) IsChannel() bool {
	return c.Type == "channel"
}

// MessageEntity object represents one special entity in a text message. For example, hashtags, usernames, URLs, etc.
type MessageEntity struct {
	// Type of the entity.
	Type string `json:"type"`

	// Offset in UTF-16 code units to the start of the entity
	Offset int `json:"offset"`

	// Length of the entity in UTF-16 code units
	Length int `json:"length"`

	// Optional. For “text_link” only, url that will be opened after user taps on the text
	URL string `json:"url,omitempty"`

	// Optional. For “text_mention” only, the mentioned user
	User *User `json:"user"`

	// Optional. For “pre” only, the programming language of the entity text
	Language string `json:"language,omitempty"`
}

// MessageID unique message identifier inside this chat
type MessageID int

// Message object represents a message.
type Message struct {
	// Unique message identifier inside this chat
	ID MessageID `json:"message_id"`

	// Optional. Sender, empty for messages sent to channels
	From *User `json:"from,omitempty"`

	// Date the message was sent in Unix time.
	Date UnixTime `json:"date"`

	// Conversation the message belongs to
	Chat Chat `json:"chat"`

	// Optional. For forwarded messages, sender of the original message
	ForwardFrom *User `json:"forward_from,omitempty"`

	// Optional. For messages forwarded from channels, information about the original channel
	ForwardFromChat *Chat `json:"forward_from_chat,omitempty"`

	// Optional. For messages forwarded from channels, identifier of the original message in the channel
	ForwardFromMessageID MessageID `json:"forward_from_message_id,omitempty"`

	// Optional. For messages forwarded from channels, signature of the post author if present
	ForwardSignature string `json:"forward_signature,omitempty"`

	// Optional. Sender's name for messages forwarded from users
	// who disallow adding a link to their account in forwarded messages
	ForwardSenderName string `json:"forward_sender_name,omitempty"`

	// Optional. For forwarded messages, date the original message was sent in Unix time
	ForwardDate UnixTime `json:"forward_date,omitempty"`

	// Optional. For replies, the original message.
	// Note that the Message object in this field will not contain further reply_to_message
	// fields even if it itself is a reply.
	ReplyToMessage *Message `json:"reply_to_message,omitempty"`

	// Optional. Date the message was last edited in Unix time
	EditDate UnixTime `json:"edit_date,omitempty"`

	// Optional. The unique identifier of a media message group this message belongs to
	MediaGroupID string `json:"media_group_id,omitempty"`

	// Optional. Signature of the post author for messages in channels
	AuthorSignature string `json:"author_signature,omitempty"`

	// Optional. For text messages, the actual UTF-8 text of the message, 0-4096 characters
	Text string `json:"text,omitempty"`

	// Optional. For text messages, special entities like usernames, URLs, bot commands, etc. that appear in the text
	Entities []MessageEntity `json:"entities,omitempty"`

	// Optional. Message is an audio file, information about the file.
	Audio *Audio `json:"audio,omitempty"`

	// Optional. Message is a general file, information about the file
	Document *Document `json:"document,omitempty"`

	// Optional. Message is an animation, information about the animation. For backward compatibility,
	// when this field is set, the document field will also be set.
	Animation *Animation `json:"animation,omitempty"`

	// Optional. Message is a game, information about the game.
	Game *Game `json:"game,omitempty"`

	// Optional. Message is a photo, available sizes of the photo
	Photo []PhotoSize `json:"photo,omitempty"`

	// Optional. Message is a sticker, information about the sticker
	Sticker *Sticker `json:"sticker,omitempty"`

	// Optional. Message is a video, information about the video
	Video *Video `json:"video,omitempty"`

	// Optional. Message is a voice message, information about the file
	Voice *Voice `json:"voice,omitempty"`

	// Optional. Message is a video note, information about the video message
	VideoNote *VideoNote `json:"video_note,omitempty"`

	// Optional. Caption for the animation, audio, document, photo, video or voice, 0-1024 characters
	Caption string `json:"caption,omitempty"`

	// Optional. For messages with a caption, special entities like usernames, URLs, bot commands, etc.
	// that appear in the caption
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`

	// Optional. Message is a shared contact, information about the contact
	Contact *Contact `json:"contact,omitempty"`

	// Optional. Message is a shared location, information about the location
	Location *Location `json:"location,omitempty"`

	// Optional. Message is a venue, information about the venue
	Venue *Venue `json:"venue,omitempty"`

	// Optional. Message is a native poll, information about the poll
	Poll *Poll `json:"poll,omitempty"`

	// Optional. Message is a dice with random value from 1 to 6.
	Dice *Dice `json:"dice,omitempty"`

	// Optional. New members that were added to the group or supergroup and information about them
	// (the bot itself may be one of these members).
	NewChatMembers []User `json:"new_chat_members,omitempty"`

	// Optional. A member was removed from the group, information about them (this member may be the bot itself)
	LeftChatMember *User `json:"left_chat_member,omitempty"`

	// Optional. A chat title was changed to this value
	NewChatTitle string `json:"new_chat_title,omitempty"`

	// Optional. A chat photo was change to this value
	NewChatPhoto []PhotoSize `json:"new_chat_photo,omitempty"`

	// Optional. Service message: the chat photo was deleted
	DeleteChatPhoto bool `json:"delete_chat_photo,omitempty"`

	// Optional. Service message: the group has been created
	GroupChatCreated bool `json:"group_chat_created,omitempty"`

	// Optional. Service message: the supergroup has been created. This field can‘t be received in a message coming through updates,
	// because bot can’t be a member of a supergroup when it is created.
	// It can only be found in reply_to_message if someone replies to a very first message
	// in a directly created supergroup.
	SupergroupChatCreated bool `json:"supergroup_chat_created,omitempty"`

	// Optional. Service message: the channel has been created. This field can‘t be received in a message coming through updates,
	// because bot can’t be a member of a channel when it is created.
	// It can only be found in reply_to_message if someone replies to a very first message in a channel.
	ChannelChatCreated bool `json:"channel_chat_created,omitempty"`

	// Optional. The group has been migrated to a supergroup with the specified identifier.
	MigrateToChatID ChatID `json:"migrate_to_chat_id,omitempty"`

	// Optional. The supergroup has been migrated from a group with the specified identifier.
	MigrateFromChatID ChatID `json:"migrate_from_chat_id,omitempty"`

	// Optional. Specified message was pinned. Note that the Message object in this field will not contain further
	// reply_to_message fields even if it is itself a reply.
	PinnedMessage *Message `json:"pinned_message,omitempty"`

	// Optional. Message is an invoice for a payment, information about the invoice.
	Invoice json.RawMessage `json:"invoice,omitempty"`

	// Optional. Message is a service message about a successful payment, information about the payment.
	SuccessfulPayment json.RawMessage `json:"successful_payment,omitempty"`

	// Optional. The domain name of the website on which the user has logged in.
	ConnectedWebsite string `json:"connected_website,omitempty"`

	// Optional. Telegram Passport data
	PassportData json.RawMessage `json:"passport_data,omitempty"`

	// Optional. Optional. Inline keyboard attached to the message.
	// login_url buttons are represented as ordinary url buttons.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`

	client *Client
}

func (msg *Message) IsCommand() bool {
	return strings.HasPrefix(msg.Text, "/")
}

func (msg *Message) CommandArgs() string {
	if !msg.IsCommand() {
		return ""
	}

	result := strings.SplitN(msg.Text, " ", 2)
	if len(result) > 1 {
		return result[1]
	}
	return ""
}

func (msg *Message) ensureClientBind() {
	if msg.client == nil {
		panic("client does not bind to message")
	}
}

// AnswerText send text answer to this message.
func (msg *Message) AnswerText(ctx context.Context, text string, opts *TextOpts) (*Message, error) {
	msg.ensureClientBind()
	return msg.client.SendText(ctx, msg.Chat.ID, text, opts)
}

// ReplyText send text reply to this message. This message was quoted.
func (msg *Message) ReplyText(ctx context.Context, text string, opts *TextOpts) (*Message, error) {
	if opts == nil {
		opts = &TextOpts{}
	}

	opts.ReplyToMessageID = msg.ID

	return msg.AnswerText(ctx, text, opts)
}

// AnswerPhoto send photo answer to this message.
func (msg *Message) AnswerPhoto(ctx context.Context, photo *InputFile, opts *PhotoOpts) (*Message, error) {
	msg.ensureClientBind()
	return msg.client.SendPhoto(ctx, msg.Chat.ID, photo, opts)
}

// ReplyText send photo reply to this message. This message was quoted.
func (msg *Message) ReplyPhoto(ctx context.Context, photo *InputFile, opts *PhotoOpts) (*Message, error) {
	if opts == nil {
		opts = &PhotoOpts{}
	}

	opts.ReplyToMessageID = msg.ID

	return msg.AnswerPhoto(ctx, photo, opts)
}

// AnswerLocation send location answer to this message.
func (msg *Message) AnswerLocation(ctx context.Context, location Location, opts *LocationOpts) (*Message, error) {
	msg.ensureClientBind()
	return msg.client.SendLocation(ctx, msg.Chat.ID, location, opts)
}

// EditLiveLocation this method edit live location messages sent by the bot or via the bot (for inline bots).
// A location can be edited until its live_period expires or
// editing is explicitly disabled by a call to stopMessageLiveLocation.
//
// From options, only ReplyMarkup is usable
func (msg *Message) EditLiveLocation(ctx context.Context, location Location, opts *LocationOpts) (*Message, error) {
	msg.ensureClientBind()

	return msg.client.EditLiveLocation(
		ctx,
		msg.Chat.ID,
		msg.ID,
		location,
		opts,
	)
}

// ReplyLocation send photo reply to this message. This message was quoted.
func (msg *Message) ReplyLocation(ctx context.Context, location Location, opts *LocationOpts) (*Message, error) {
	if opts == nil {
		opts = &LocationOpts{}
	}

	opts.ReplyToMessageID = msg.ID

	return msg.AnswerLocation(ctx, location, opts)
}

// Audio object represents an audio file to be treated as music by the Telegram clients.
type Audio struct {
	// Identifier for this file, which can be used to download or reuse the file
	FileID FileID `json:"file_id"`

	// Unique identifier for this file, which is supposed to be the same over time and for different bots.
	// Can't be used to download or reuse the file.
	FileUniqueID FileID `json:"file_unique_id"`

	// Duration of the audio in seconds as defined by sender
	Duration Duration `json:"duration"`

	// Optional. Performer of the audio as defined by sender or by audio tags
	Performer string `json:"performer,omitempty"`

	// Optional. Title of the audio as defined by sender or by audio tags
	Title string `json:"title,omitempty"`

	// Optional. MIME type of the file as defined by sender
	MIMEType string `json:"mime_type,omitempty"`

	// Optional. File size
	FileSize int `json:"file_size,omitempty"`

	// Optional. Thumbnail of the album cover to which the music file belongs
	Thumb *PhotoSize `json:"thumb,omitempty"`
}

// Document object represents a general file (as opposed to photos, voice messages and audio files).
type Document struct {
	// Identifier for this file, which can be used to download or reuse the file
	FileID FileID `json:"file_id"`

	// Unique identifier for this file, which is supposed to be the same over time and for different bots.
	// Can't be used to download or reuse the file.
	FileUniqueID FileID `json:"file_unique_id"`

	// Optional. Document thumbnail as defined by sender
	Thumb *PhotoSize `json:"thumb,omitempty"`

	// Optional. Original filename as defined by sender
	FileName string `json:"file_name,omitempty"`

	// Optional. MIME type of the file as defined by sender
	MIMEType string `json:"mime_type,omitempty"`

	// Optional. File size
	FileSize int `json:"file_size,omitempty"`
}

// Animation object represents an animation file (GIF or H.264/MPEG-4 AVC video without sound).
type Animation struct {
	// Identifier for this file, which can be used to download or reuse the file
	FileID FileID `json:"file_id"`

	// Unique identifier for this file, which is supposed to be the same over time and for different bots.
	// Can't be used to download or reuse the file.
	FileUniqueID FileID `json:"file_unique_id"`

	// Video width as defined by sender
	Width int `json:"width"`

	// Video height as defined by sender
	Height int `json:"height"`

	// Duration of the video in seconds as defined by sender
	Duration Duration `json:"duration"`

	// Optional. Animation thumbnail as defined by sender
	Thumb *PhotoSize `json:"thumb,omitempty"`

	// Optional. Original animation filename as defined by sender
	FileName string `json:"file_name,omitempty"`

	// Optional. MIME type of the file as defined by sender
	MIMEType string `json:"mime_type,omitempty"`

	// Optional. File size
	FileSize int `json:"file_size,omitempty"`
}

// Game object represents a game.
// Use BotFather to create and edit games, their short names will act as unique identifiers.
type Game struct {
	// Title of the game
	Title string `json:"title"`

	// Description of the game
	Description string `json:"description"`

	// Photo that will be displayed in the game message in chats.
	Photo []PhotoSize `json:"photo"`

	// Optional. Brief description of the game or high scores included in the game message.
	// Can be automatically edited to include current high scores for the game when the bot calls setGameScore,
	// or manually edited using editMessageText. 0-4096 characters.
	Text string `json:"text,omitempty"`

	// Optional. Special entities that appear in text, such as usernames, URLs, bot commands, etc.
	TextEntities []MessageEntity `json:"text_entities,omitempty"`

	// Optional. Animation that will be displayed in the game message in chats. Upload via BotFather.
	Animation *Animation `json:"animation,omitempty"`
}

// MaskPosition object describes the position on faces where a mask should be placed by default.
type MaskPosition struct {
	// The part of the face relative to which the mask should be placed. One of “forehead”, “eyes”, “mouth”, or “chin”.
	Point string `json:"point"`

	// Shift by X-axis measured in widths of the mask scaled to the face size, from left to right.
	// For example, choosing -1.0 will place mask just to the left of the default mask position.
	XShift float64 `json:"x_shift"`

	// Shift by Y-axis measured in heights of the mask scaled to the face size, from top to bottom.
	// For example, 1.0 will place the mask just below the default mask position.
	YShift float64 `json:"y_shift"`

	// Mask scaling coefficient. For example, 2.0 means double size.
	Scale float64 `json:"scale"`
}

// Sticker object represents a sticker.
type Sticker struct {
	// Identifier for this file, which can be used to download or reuse the file
	FileID FileID `json:"file_id"`

	// Unique identifier for this file, which is supposed to be the same over time and for different bots.
	// Can't be used to download or reuse the file.
	FileUniqueID FileID `json:"file_unique_id"`

	// Sticker width
	Width int `json:"width"`

	// Sticker height
	Height int `json:"height"`

	// True, if the sticker is animated
	IsAnimated int `json:"is_animated"`

	// Optional. Sticker thumbnail in the .WEBP or .JPG format
	Thumb *PhotoSize `json:"thumb,omitempty"`

	// Optional. Emoji associated with the sticker
	Emoji string `json:"emoji,omitempty"`

	// Optional. Name of the sticker set to which the sticker belongs
	SetName string `json:"set_name,omitempty"`

	// Optional. For mask stickers, the position where the mask should be placed
	MaskPosition *MaskPosition `json:"mask_position,omitempty"`

	// Optional. File size
	FileSize int `json:"file_size,omitempty"`
}

// Video object represents a video file.
type Video struct {
	// Identifier for this file, which can be used to download or reuse the file
	FileID FileID `json:"file_id"`

	// Unique identifier for this file, which is supposed to be the same over time and for different bots.
	// Can't be used to download or reuse the file.
	FileUniqueID FileID `json:"file_unique_id"`

	// Video width as defined by sender
	Width int `json:"width"`

	// Video height as defined by sender
	Height int `json:"height"`

	// Duration of the video in seconds as defined by sender
	Duration Duration `json:"duration"`

	// Optional. Video thumbnail
	Thumb *PhotoSize `json:"thumb,omitempty"`

	// Optional. Mime type of a file as defined by sender
	MIMEType string `json:"mime_type,omitempty"`

	// Optional. File size
	FileSize int `json:"file_size,omitempty"`
}

// Voice object represents a voice note.
type Voice struct {
	// Identifier for this file, which can be used to download or reuse the file
	FileID FileID `json:"file_id"`

	// Unique identifier for this file, which is supposed to be the same over time and for different bots.
	// Can't be used to download or reuse the file.
	FileUniqueID FileID `json:"file_unique_id"`

	// Duration of the audio in seconds as defined by sender
	Duration Duration `json:"duration"`

	// Optional. Mime type of a file as defined by sender
	MIMEType string `json:"mime_type,omitempty"`

	// Optional. File size
	FileSize int `json:"file_size,omitempty"`
}

// VideNote object represents a video message (available in Telegram apps as of v.4.0).
type VideoNote struct {
	// Identifier for this file, which can be used to download or reuse the file
	FileID FileID `json:"file_id"`

	// Unique identifier for this file, which is supposed to be the same over time and for different bots.
	// Can't be used to download or reuse the file.
	FileUniqueID FileID `json:"file_unique_id"`

	// Video width and height (diameter of the video message) as defined by sender
	Length int `json:"length"`

	// Duration of the video in seconds as defined by sender
	Duration Duration `json:"duration"`

	// Optional. Video thumbnail
	Thumb *PhotoSize `json:"thumb,omitempty"`

	// Optional. File size
	FileSize int `json:"file_size,omitempty"`
}

// Contact object represents a phone contact.
type Contact struct {
	// Contact's phone number
	PhoneNumber string `json:"phone_number"`

	// Contact's first name
	FirstName string `json:"first_name"`

	// Optional. Contact's last name
	LastName string `json:"last_name,omitempty"`

	// Optional. Contact's user identifier in Telegram
	UserID UserID `json:"user_id,omitempty"`

	// Optional. Additional data about the contact in the form of a vCard
	VCard string `json:"v_card,omitempty"`
}

// Location object represents a point on the map.
type Location struct {
	// Longitude as defined by sender
	Longitude float64 `json:"longitude"`

	// Latitude as defined by sender
	Latitude float64 `json:"latitude"`
}

// Venue object represents a venue.
type Venue struct {
	// Venue location
	Location Location `json:"location"`

	// Name of the venue
	Title string `json:"title"`

	// Address of the venue
	Address string `json:"address"`

	// Optional. Foursquare identifier of the venue
	FoursquareID string `json:"foursquare_id,omitempty"`

	// Optional. Foursquare type of the venue.
	// For example, “arts_entertainment/default”, “arts_entertainment/aquarium” or “food/icecream”.
	FoursquareType string `json:"foursquare_type,omitempty"`
}

// PollOption object contains information about one answer option in a poll.
type PollOption struct {
	// Option text, 1-100 characters
	Text string `json:"text"`

	// Number of users that voted for this option
	VoterCount int `json:"voter_count"`
}

// PollAnswer object represents an answer of a user in a non-anonymous poll.
type PollAnswer struct {
	// Unique poll identifier
	PollID string `json:"poll_id"`

	// The user, who changed the answer to the poll
	User User `json:"user"`

	// 0-based identifiers of answer options, chosen by the user. May be empty if the user retracted their vote.
	OptionIDs []int `json:"option_ids"`
}

// PollID it's unique poll identifier
type PollID string

// Poll object contains information about a poll.
type Poll struct {
	// Unique poll identifier
	ID PollID `json:"id"`

	// Poll question, 1-255 characters
	Question string `json:"question"`

	// List of poll options
	Options []PollOption

	// Total number of users that voted in the poll
	TotalVoterCount int `json:"total_voter_count"`

	// True, if the poll is closed
	IsClosed bool `json:"is_closed"`

	// True, if the poll is anonymous
	IsAnonymous bool `json:"is_anonymous"`

	// Poll type, currently can be “regular” or “quiz”
	Type string `json:"type"`

	// True, if the poll allows multiple answers
	AllowsMultipleAnswers bool `json:"allows_multiple_answers"`

	// Optional. 0-based identifier of the correct answer option.
	// Available only for polls in the quiz mode, which are closed, or was sent (not forwarded) by the bot or
	// to the private chat with the bot.
	CorrectOptionID int `json:"correct_option_id,omitempty"`

	// Optional. Text that is shown when a user chooses an incorrect answer or taps
	// on the lamp icon in a quiz-style poll, 0-200 characters
	Explanation string `json:"explanation,omitempty"`

	// Optional. Special entities like usernames, URLs, bot commands, etc. that appear in the explanation
	ExplanationEntities []MessageEntity `json:"explanation_entities,omitempty"`

	// Optional. Amount of time in seconds the poll will be active after creation.
	OpenPeriod Duration `json:"open_period,omitempty"`

	// Optional. Point in time (Unix timestamp) when the poll will be automatically closed
	CloseDate UnixTime `json:"close_date,omitempty"`
}

// Dice object represents a dice with a random value from 1 to 6 for currently supported base emoji.
// Yes, we're aware of the “proper” singular of die.
// But it's awkward, and we decided to help it change.
// One dice at a time!
type Dice struct {
	// Emoji on which the dice throw animation is based
	Emoji string `json:"emoji"`

	// Value of the dice, 1-6 for currently supported base emoji
	Value int `json:"value"`
}

// KeyboardButton object represents one button of the reply keyboard.
// For simple text buttons String can be used instead of this object to specify text of the button.
// Optional fields RequestContact, RequestLocation, and RequestPoll are mutually exclusive.
type KeyboardButton struct {
	// Text of the button. If none of the optional fields are used,
	// it will be sent as a message when the button is pressed.
	Text string `json:"text"`

	// Optional. If True, the user's phone number will be sent as a contact when the button is pressed.
	// Available in private chats only.
	RequestContact bool `json:"request_contact,omitempty"`

	// Optional. If True, the user's current location will be sent when the button is pressed.
	// Available in private chats only.
	RequestLocation bool `json:"request_location,omitempty"`

	// Optional. If specified, the user will be asked to create a poll and send it to the bot when the button is pressed.
	// Available in private chats only.
	RequestPoll *KeyboardButtonPollType `json:"request_poll,omitempty"`
}

// KeyboardButtonPollType object represents type of a poll, which is allowed to be created
// and sent when the corresponding button is pressed.
type KeyboardButtonPollType struct {
	// Optional. If quiz is passed, the user will be allowed to create only polls in the quiz mode. If regular is passed,
	// only regular polls will be allowed.
	// Otherwise, the user will be allowed to create a poll of any type.
	Type string `json:"type,omitempty"`
}

// ReplyKeyboardMarkup object represents a custom keyboard with reply options.
type ReplyKeyboardMarkup struct {
	// Array of button rows, each represented by an Array of KeyboardButton objects
	Keyboard []KeyboardRow `json:"keyboard"`

	// Optional. Requests clients to resize the keyboard vertically for optimal fit
	// (e.g., make the keyboard smaller if there are just two rows of buttons).
	// Defaults to false, in which case the custom keyboard is always of the same height as the app's standard keyboard.
	ResizeKeyboard bool `json:"resize_keyboard,omitempty"`

	// Optional. Requests clients to hide the keyboard as soon as it's been used.
	// The keyboard will still be available, but clients will automatically display the usual
	// letter-keyboard in the chat – the user can press a special button in the input field to see
	// the custom keyboard again. Defaults to false.
	OneTimeKeyboard bool `json:"one_time_keyboard,omitempty"`

	// Optional. Use this parameter if you want to show the keyboard to specific users only.
	// Targets:
	//  1) users that are @mentioned in the text of the Message object;
	//  2) if the bot's message is a reply (has reply_to_message_id), sender of the original message.
	//
	// Example: A user requests to change the bot‘s language,
	// bot replies to the request with a keyboard to select the new language.
	// Other users in the group don’t see the keyboard.
	Selective bool `json:"selective,omitempty"`
}

func (markup ReplyKeyboardMarkup) isReplyMarkup() {}

type KeyboardRow []KeyboardButton

func NewKeyboardRow(buttons ...KeyboardButton) KeyboardRow {
	return buttons
}

func NewReplyKeyboardMarkup(rows ...KeyboardRow) ReplyKeyboardMarkup {
	return ReplyKeyboardMarkup{
		Keyboard: rows,
	}
}

// ReplyKeyboardRemove Upon receiving a message with this object,
// Telegram clients will remove the current custom keyboard and display the default letter-keyboard.
// By default, custom keyboards are displayed until a new keyboard is sent by a bot.
// An exception is made for one-time keyboards that are hidden immediately
// after the user presses a button (see ReplyKeyboardMarkup).
type ReplyKeyboardRemove struct {
	// Requests clients to remove the custom keyboard. User will not be able to summon this keyboard;
	// if you want to hide the keyboard from sight but keep it accessible,
	// use one_time_keyboard in ReplyKeyboardMarkup.
	RemoveKeyboard bool `json:"remove_keyboard"`

	// Optional. Use this parameter if you want to remove the keyboard for specific users only. Targets:
	// 1) users that are @mentioned in the text of the Message object;
	// 2) if the bot's message is a reply (has reply_to_message_id), sender of the original message.
	//
	// Example: A user votes in a poll, bot returns confirmation message in reply to the vote
	// and removes the keyboard for that user, while still showing the keyboard with poll options
	// to users who haven't voted yet.
	Selective bool `json:"selective,omitempty"`
}

// InlineKeyboardMarkup object represents an inline keyboard that appears right next to the message it belongs to.
type InlineKeyboardMarkup struct {
	// Array of button rows, each represented by an Array of InlineKeyboardButton objects
	InlineKeyboard []InlineKeyboardRow `json:"inline_keyboard"`
}

type InlineKeyboardRow []InlineKeyboardButton

func NewInlineKeyboardRow(buttons ...InlineKeyboardButton) InlineKeyboardRow {
	return buttons
}

func NewInlineKeyboardMarkup(rows ...InlineKeyboardRow) InlineKeyboardMarkup {
	return InlineKeyboardMarkup{
		InlineKeyboard: rows,
	}
}

func (markup InlineKeyboardMarkup) isReplyMarkup() {}

// LoginURL object represents a parameter of the inline keyboard button
// used to automatically authorize a user.
// Serves as a great replacement for the Telegram Login Widget when the user is coming from Telegram.
// All the user needs to do is tap/click a button and confirm that they want to log in.
type LoginURL struct {
	// An HTTP URL to be opened with user authorization data
	// added to the query string when the button is pressed.
	//
	// If the user refuses to provide authorization data,
	// the original URL without information about the user will be opened.
	// The data added is the same as described in Receiving authorization data.
	//
	// NOTE: You must always check the hash of the received data to verify the authentication
	// and the integrity of the data.
	URL string `json:"url"`

	// Optional. New text of the button in forwarded messages.
	ForwardText string `json:"forward_text,omitempty"`

	// Optional. Username of a bot, which will be used for user authorization. See Setting up a bot for more details. If not specified, the current bot's username will be assumed.
	// The url's domain must be the same as the domain linked with the bot.
	BotUsername Username `json:"bot_username,omitempty"`

	// Optional. Pass True to request the permission for your bot to send messages to the user.
	RequestWriteAccess bool `json:"request_write_access,omitempty"`
}

// InlineKeyboardButton object represents one button of an inline keyboard.
// You must use exactly one of the optional fields.
type InlineKeyboardButton struct {
	// Label text on the button
	Text string `json:"text"`

	// Optional. HTTP or tg:// url to be opened when button is pressed
	URL string `json:"url"`

	// Optional. An HTTP URL used to automatically authorize the user.
	// Can be used as a replacement for the Telegram Login Widget.
	LoginURL *LoginURL `json:"login_url,omitempty"`

	// Optional. Data to be sent in a callback query to the bot when button is pressed, 1-64 bytes
	CallbackData string `json:"callback_data,omitempty"`

	// Optional. If set, pressing the button will prompt the user to select one of their chats,
	// open that chat and insert the bot‘s username and the specified inline query in the input field.
	// Can be empty, in which case just the bot’s username will be inserted.
	//
	// Note: This offers an easy way for users to start using your bot
	// in inline mode when they are currently in a private chat with it.
	// Especially useful when combined with switch_pm… actions – in this case
	// the user will be automatically returned to the chat they switched from,
	// skipping the chat selection screen.
	SwitchInlineQuery string `json:"switch_inline_query,omitempty"`

	// Optional. If set, pressing the button will insert the bot‘s username
	// and the specified inline query in the current chat’s input field.
	// Can be empty, in which case only the bot's username will be inserted.
	//
	// This offers a quick way for the user to open your bot in inline mode in the same chat
	// good for selecting something from multiple options.
	SwitchInlineQueryCurrentChat string `json:"switch_inline_query_current_chat,omitempty"`

	// Optional. Description of the game that will be launched when the user presses the button.
	// NOTE: This type of button must always be the first button in the first row.
	CallbackGame *CallbackGame `json:"callback_game,omitempty"`

	// Optional. Specify True, to send a Pay button.
	// NOTE: This type of button must always be the first button in the first row.
	Pay bool `json:"pay,omitempty"`
}

func NewInlineKeyboardButtonLoginURL(text string, login *LoginURL) InlineKeyboardButton {
	return InlineKeyboardButton{
		Text:     text,
		LoginURL: login,
	}
}

func NewInlineKeyboardButtonURL(text, url string) InlineKeyboardButton {
	return InlineKeyboardButton{
		Text: text,
		URL:  url,
	}
}

func NewInlineKeyboardButtonCallbackData(text, cd string) InlineKeyboardButton {
	return InlineKeyboardButton{
		Text:         text,
		CallbackData: cd,
	}
}

// Optional. Description of the game that will be launched when the user presses the button.
type CallbackGame struct{}

// CallbackQueryID unique ID.
type CallbackQueryID string

// CallbackQuery This object represents an incoming callback query
// from a callback button in an inline keyboard.
// If the button that originated the query
// was attached to a message sent by the bot, the field message will be present.
// If the button was attached to a message sent via the bot (in inline mode),
// the field inline_message_id will be present.
// Exactly one of the fields data or game_short_name will be present.
type CallbackQuery struct {
	// Unique identifier for this query
	ID CallbackQueryID `json:"id"`

	// Sender
	From *User `json:"from"`

	// Optional. Message with the callback button that originated the query.
	// Note that message content and message date will not be available if the message is too old.
	Message *Message `json:"message,omitempty"`

	// Optional. Identifier of the message sent via the bot in inline mode, that originated the query.
	InlineMessageID string `json:"inline_message_id,omitempty"`

	// Global identifier, uniquely corresponding to the chat
	// to which the message with the callback button was sent.
	// Useful for high scores in games.
	ChatInstance string `json:"chat_instance,omitempty"`

	// Optional. Data associated with the callback button.
	// Be aware that a bad client can send arbitrary data in this field.
	Data string `json:"data,omitempty"`

	// Optional. Short name of a Game to be returned, serves as the unique identifier for the game
	GameShortName string `json:"game_short_name,omitempty"`
}

// Contains information about the current status of a webhook.
type WebhookInfo struct {
	// Webhook URL, may be empty if webhook is not set up
	URL string `json:"url"`

	// True, if a custom certificate was provided for webhook certificate checks
	HasCustomCertificate bool `json:"has_custom_certificate"`

	// Number of updates awaiting delivery
	PendingUpdateCount int `json:"pending_update_count"`

	// Optional. Unix time for the most recent error that happened when trying to deliver an update via webhook
	LastErrorDate UnixTime `json:"last_error_date,omitempty"`

	// Optional. Error message in human-readable format
	// for the most recent error that happened
	// when trying to deliver an update via webhook
	LastErrorMessage string `json:"last_error_message,omitempty"`

	// Optional. Maximum allowed number of simultaneous HTTPS connections to the webhook for update delivery.
	MaxConnections int `json:"max_connections,omitempty"`

	// Optional. A list of update types the bot is subscribed to. Defaults to all update types.
	AllowedUpdates []string `json:"allowed_updates,omitempty"`
}

type File struct {
	// Identifier for this file, which can be used to download or reuse the file.
	FileID FileID `json:"file_id"`

	// Unique identifier for this file, which is supposed to be the same over time and for different bots.
	// Can't be used to download or reuse the file.
	FileUniqueID string `json:"file_unique_id"`

	// Optional. File size, if known.
	FileSize int `json:"file_size,omitempty"`

	// Optional. File path. Use https://api.telegram.org/file/bot<token>/<file_path> to get the file.
	FilePath string `json:"file_path,omitempty"`

	client *Client
}

func (f File) URL() string {
	return f.client.getFileURL(f.FilePath)
}

func (f File) NewReader(ctx context.Context) (io.ReadCloser, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, f.URL(), nil)
	if err != nil {
		return nil, errors.Wrap(err, "build request")
	}

	res, err := f.client.doer.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "execute request")
	}
	return res.Body, nil
}

type ChatMember struct {
	// Information about the user
	User *User `json:"user,omitempty"`

	// The member's status in the chat. Can be “creator”, “administrator”, “member”, “restricted”, “left” or “kicked”
	Status string `json:"status,omitempty"`

	// Optional. Owner and administrators only. Custom title for this user
	CustomTitle string `json:"custom_title,omitempty"`

	// Optional. Owner and administrators only. True, if the user's presence in the chat is hidden
	IsAnonymous bool `json:"is_anonymous,omitempty"`

	// Optional. Administrators only. True, if the bot is allowed to edit administrator privileges of that user
	CanBeEdited bool `json:"can_be_edited,omitempty"`

	// Optional. Administrators only. True, if the administrator can post in the channel; channels only
	CanPostMessages bool `json:"can_post_messages,omitempty"`

	// Optional. Administrators only. True, if the administrator can edit messages of other users and can pin messages; channels only
	CanEditMessages bool `json:"can_edit_messages,omitempty"`

	// Optional. Administrators only. True, if the administrator can delete messages of other users
	CanDeleteMessages bool `json:"can_delete_messages,omitempty"`

	// Optional. Administrators only. True, if the administrator can restrict, ban or unban chat members
	CanRestrictMembers bool `json:"can_restrict_members,omitempty"`

	// Optional. Administrators only. True, if the administrator can add new administrators
	// with a subset of their own privileges or demote administrators that he has promoted,
	// directly or indirectly (promoted by administrators that were appointed by the user)
	CanPromoteMembers bool `json:"can_promote_members,omitempty"`

	// Optional. Administrators and restricted only. True, if the user is allowed to change the chat title, photo and other settings
	CanChangeInfo bool `json:"can_change_info,omitempty"`

	// Optional. Administrators and restricted only. True, if the user is allowed to invite new users to the chat
	CanInviteUsers bool `json:"can_invite_users,omitempty"`

	// Optional. Administrators and restricted only. True, if the user is allowed to pin messages; groups and supergroups only
	CanPinMessages bool `json:"can_pin_messages,omitempty"`

	// Optional. Restricted only. True, if the user is a member of the chat at the moment of the request
	IsMember bool `json:"is_member,omitempty"`

	// Optional. Restricted only. True, if the user is allowed to send text messages, contacts, locations and venues
	CanSendMessages bool `json:"can_send_messages,omitempty"`

	// Optional. Restricted only. True, if the user is allowed to send audios, documents, photos, videos, video notes and voice notes
	CanSendMediaMessages bool `json:"can_send_media_messages,omitempty"`

	// Optional. Restricted only. True, if the user is allowed to send polls
	CanSendPolls bool `json:"can_send_polls,omitempty"`

	// Optional. Restricted only. True, if the user is allowed to send animations, games, stickers and use inline bots
	CanSendOtherMessages bool `json:"can_send_other_messages,omitempty"`

	// Optional. Restricted only. True, if the user is allowed to add web page previews to their messages
	CanAddWebPagePreviews bool `json:"can_add_web_page_previews,omitempty"`

	// Optional. Restricted and kicked only. Date when restrictions will be lifted for this user; unix time
	UnitDate int `json:"unit_date,omitempty"`
}
