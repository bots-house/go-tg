package tg

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_SendText(t *testing.T) {
	ctx := context.Background()

	text := "This is *test* of [bots-house/go-tg](https://github.com/bots-house/go-tg)"

	kb := [][]InlineKeyboardButton{
		{
			{Text: "🔥", URL: "https://github.com/bots-house/go-tg"},
		},
	}

	msg, err := testClient.SendText(ctx, testChannelID, text, &TextOpts{
		ParseMode:             "markdown",
		DisableWebPagePreview: true,
		DisableNotification:   true,
		ReplyToMessageID:      MessageID(1),
		ReplyMarkup:           InlineKeyboardMarkup{kb},
	})

	if assert.NoError(t, err) && assert.NotNil(t, msg) {
		assert.NotZero(t, msg.ID)
		assert.NotNil(t, msg.Chat)
	}
}

func TestClient_SendPhoto(t *testing.T) {
	ctx := context.Background()

	photo, err := NewInputFileLocal("testdata/gopher.png")
	assert.NoError(t, err)
	defer photo.Close()

	caption := "The [Go](https://golang.org) _gopher_ is an iconic mascot and one of the most distinctive features of the Go project."

	kb := [][]InlineKeyboardButton{
		{
			{Text: "🔥", URL: "https://github.com/bots-house/go-tg"},
		},
	}

	msg, err := testClient.SendPhoto(ctx, testChannelID, photo, &PhotoOpts{
		Caption:             caption,
		ParseMode:           "markdown",
		DisableNotification: true,
		ReplyToMessageID:    MessageID(1),
		ReplyMarkup:         InlineKeyboardMarkup{kb},
	})

	if assert.NoError(t, err) && assert.NotNil(t, msg) {
		assert.NotZero(t, msg.ID)
		assert.NotNil(t, msg.Chat)
	}
}

func TestInputMediaMarshal(t *testing.T) {
	photo1, err := NewInputFileLocal("testdata/gopher.png")
	assert.NoError(t, err)
	defer photo1.Close()

	video1, err := NewInputFileLocal("testdata/robot.mp4")
	assert.NoError(t, err)
	defer video1.Close()

	video1Thumb, err := NewInputFileLocal("testdata/gopher.png")
	assert.NoError(t, err)
	defer video1Thumb.Close()

	photo2, err := NewInputFileLocal("testdata/gopher.png")
	assert.NoError(t, err)
	defer photo2.Close()

	data := []InputMedia{
		InputMediaPhoto{
			Media:     photo1,
			Caption:   "this is **gopher**",
			ParseMode: "markdown",
		},
		InputMediaVideo{
			Media:             video1,
			Thumb:             video1Thumb,
			Caption:           "this robot video, but it was **gopher**!?",
			ParseMode:         "markdown",
			Width:             560,
			Height:            320,
			Duration:          6,
			SupportsStreaming: true,
		},
		InputMediaPhoto{
			Media:     photo2,
			Caption:   "mm, this is **gopher** again, really?",
			ParseMode: "markdown",
		},
	}

	v, err := json.Marshal(data)
	t.Log(string(v))
}

func TestClient_SendMediaGroup(t *testing.T) {
	ctx := context.Background()

	photo1, err := NewInputFileLocal("testdata/gopher.png")
	assert.NoError(t, err)
	defer photo1.Close()

	video1, err := NewInputFileLocal("testdata/robot.mp4")
	assert.NoError(t, err)
	defer video1.Close()

	video1Thumb, err := NewInputFileLocal("testdata/gopher.png")
	assert.NoError(t, err)
	defer video1Thumb.Close()

	photo2, err := NewInputFileLocal("testdata/gopher.png")
	assert.NoError(t, err)
	defer photo2.Close()

	group := []InputMedia{
		InputMediaPhoto{
			Media:     photo1,
			Caption:   "this is **gopher**",
			ParseMode: "markdown",
		},
		InputMediaVideo{
			Media:             video1,
			Thumb:             video1Thumb,
			Caption:           "this robot video, but it was **gopher**!?",
			ParseMode:         "markdown",
			Width:             560,
			Height:            320,
			Duration:          6,
			SupportsStreaming: true,
		},
		InputMediaPhoto{
			Media:     photo2,
			Caption:   "mm, this is **gopher** again, really?",
			ParseMode: "markdown",
		},
	}

	msg, err := testClient.SendMediaGroup(ctx, testChannelID, group, &MediaGroupOpts{
		DisableNotification: true,
		ReplyToMessageID:    1,
	})

	if assert.NoError(t, err) {
		assert.Len(t, msg, 3)
	}

}
