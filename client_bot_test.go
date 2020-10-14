package tg

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testToken     = os.Getenv("GO_TG_BOT_TOKEN")
	testChannelID ChatID
	testClient    = NewClient(testToken)
	testUserID    UserID
)

func init() {
	testChannelIDString := os.Getenv("GO_TG_TEST_CHANNEL_ID")
	if testChannelIDString != "" {
		parsed, err := strconv.ParseInt(testChannelIDString, 10, 64)
		if err != nil {
			panic(fmt.Sprintf("test channel id is provided but not number (%s)", testChannelIDString))
		}
		testChannelID = ChatID(parsed)
	}

	testUserIDString := os.Getenv("GO_TG_USER_ID")
	if testUserIDString != "" {
		parsed, err := strconv.ParseInt(testUserIDString, 10, 64)
		if err != nil {
			panic(fmt.Sprintf("test user id is provided but not number (%s)", testUserIDString))
		}
		testUserID = UserID(parsed)
	}

}

func skipIfNeed(t *testing.T) {
	noToken := testToken == ""
	isShort := testing.Short()
	if isShort || noToken {
		t.Skipf("skip test because isShort:%v, noToken:%v", isShort, noToken)
	}
}

func TestClient_GetMe(t *testing.T) {
	skipIfNeed(t)

	user, err := testClient.GetMe(context.Background())
	if assert.NoError(t, err) {
		assert.NotZero(t, user.ID)
		assert.True(t, user.IsBot)
		assert.NotZero(t, user.FirstName)
		assert.NotZero(t, user.Username)
		assert.True(t, user.SupportsInlineQueries)
		assert.True(t, user.CanReadAllGroupMessages)
		assert.True(t, user.CanJoinGroups)
	}
}

func TestClient_SetWebhook(t *testing.T) {
	skipIfNeed(t)
	ctx := context.Background()

	opts := &SetWebhookOptions{
		MaxConnections: 40,
		URL:            "https://bots.house",
	}

	err := testClient.SetWebhook(ctx, opts)

	if assert.NoError(t, err) {
		info, err := testClient.GetWebhookInfo(ctx)
		if assert.NoError(t, err) {
			assert.Equal(t, info.URL, opts.URL)
			assert.Equal(t, info.MaxConnections, opts.MaxConnections)
			assert.Equal(t, opts.AllowedUpdates, info.AllowedUpdates)
		}
	}
}

func TestClient_GetChat(t *testing.T) {
	skipIfNeed(t)

	ctx := context.Background()

	chat, err := testClient.GetChat(ctx, testChannelID)
	if assert.NoError(t, err) {
		assert.Equal(t, testChannelID, chat.ID)
	}
}

func TestClient_SetAndGetMyCommands(t *testing.T) {
	skipIfNeed(t)

	ctx := context.Background()

	commands := []BotCommand{
		{"start", "just start bot"},
		{"stop", "just stop bot"},
		{"restart", "just restart bot"},
		{"help", "show helps"},
	}

	err := testClient.SetMyCommands(ctx, commands)

	if assert.NoError(t, err) {
		actual, err := testClient.GetMyCommands(ctx)
		if assert.NoError(t, err) {
			assert.Equal(t, commands, actual)
		}
	}
}

func TestClient_GetUserProfilePhotos(t *testing.T) {
	skipIfNeed(t)

	ctx := context.Background()

	opts := &GetUserProfilePhotosOptions{
		UserID: testUserID,
	}

	_, err := testClient.GetUserProfilePhotos(ctx, opts)
	assert.NoError(t, err)
}
