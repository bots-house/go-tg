package tg

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testToken = os.Getenv("GO_TG_BOT_TOKEN")
	testClient = NewClient(testToken)
)

func skipIfNeed(t *testing.T) {
	noToken := testToken == ""
	isShort := testing.Short()
	if  isShort || noToken {
		t.Skipf("skip test becouse isShort:%v, noToken:%v", isShort, noToken)
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