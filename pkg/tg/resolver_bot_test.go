package tg

import (
	"context"
	"fmt"
	"testing"

	tgbotapi "github.com/bots-house/telegram-bot-api"
	"github.com/stretchr/testify/assert"
)

func TestBotResolver_Resolve(t *testing.T) {
	const testJoinLink = "https://t.me/joinchat/AAAAAFAwRhlmcW2KRm-A1g"

	t.Run("successfully resolve", func(t *testing.T) {
		//GIVEN
		joinLink := testJoinLink
		api := BotAPIMock{
			getChat: func(config tgbotapi.ChatConfig) (chat tgbotapi.Chat, err error) {
				return tgbotapi.Chat{ID: -1001345340953, Type: "channel", Title: "Test channel"}, nil
			},
			getChatMembersCount: func(config tgbotapi.ChatConfig) (i int, err error) {
				return 3, nil
			},
		}
		expResult := &ResolveResult{Channel: &Channel{ID: -1001345340953, Name: "Test channel", MembersCount: 3}}

		//WHEN
		bot := BotResolver{Client: api}
		result, err := bot.Resolve(context.Background(), joinLink)

		//THEN
		assert.NoError(t, err)
		assert.Equal(t, expResult, result)
	})

	t.Run("invalid join link", func(t *testing.T) {
		//GIVEN
		joinLink := "https://t.me/joinchat/invalid"

		//WHEN
		bot := BotResolver{}
		result, err := bot.Resolve(context.Background(), joinLink)

		//THEN
		assert.Nil(t, result)
		assert.EqualError(t, err, "failed to decode base64 of join link: illegal base64 data at input byte 8")
	})

	t.Run("'getChatMembersCount' fails", func(t *testing.T) {
		//GIVEN
		joinLink := testJoinLink
		api := BotAPIMock{
			getChat: func(config tgbotapi.ChatConfig) (chat tgbotapi.Chat, err error) {
				return tgbotapi.Chat{ID: -1001345340953, Type: "channel", Title: "Test channel"}, nil
			},
			getChatMembersCount: func(config tgbotapi.ChatConfig) (i int, err error) {
				return 0, fmt.Errorf("went wrong")
			},
		}

		//WHEN
		bot := BotResolver{Client: api}
		result, err := bot.Resolve(context.Background(), joinLink)

		//THEN
		assert.Nil(t, result)
		assert.EqualError(t, err, "can't get members count: went wrong")
	})

	t.Run("'getChat' fails", func(t *testing.T) {
		//GIVEN
		joinLink := testJoinLink
		api := BotAPIMock{
			getChat: func(config tgbotapi.ChatConfig) (chat tgbotapi.Chat, err error) {
				return tgbotapi.Chat{}, fmt.Errorf("get chat: went wrong")
			},
		}

		//WHEN
		bot := BotResolver{Client: api}
		result, err := bot.Resolve(context.Background(), joinLink)

		//THEN
		assert.Nil(t, result)
		assert.EqualError(t, err, "get chat: get chat: went wrong")
	})

	t.Run("chat not found", func(t *testing.T) {
		//GIVEN
		joinLink := testJoinLink
		api := BotAPIMock{
			getChat: func(config tgbotapi.ChatConfig) (chat tgbotapi.Chat, err error) {
				return tgbotapi.Chat{}, &tgbotapi.Error{Message: "Bad Request: chat not found"}
			},
		}

		//WHEN
		bot := BotResolver{Client: api}
		result, err := bot.Resolve(context.Background(), joinLink)

		//THEN
		assert.Nil(t, result)
		assert.EqualError(t, err, ErrEntityNotFoundOrBotIsNotAdmin.Error())
	})

}

func TestBotResolver_decodeJoinLink(t *testing.T) {
	testCases := []struct {
		username string

		expPayload joinLinkPayload
	}{
		{
			username:   "AAAAAEQErDeBcAxzcnzSAA",
			expPayload: joinLinkPayload{CreatorUserID: 0, GlobalChatID: 1141156919, RandomID: 9326968518265852416},
		},
		{
			username:   "AAAAAE-Xu0Ah6FIrMt9E9w",
			expPayload: joinLinkPayload{CreatorUserID: 0, GlobalChatID: 1335343936, RandomID: 2443293143339058423},
		},
		{
			username:   "AAAAAFemkQmj_o7kfUXTuA",
			expPayload: joinLinkPayload{CreatorUserID: 0, GlobalChatID: 1470533897, RandomID: 11817039584272176056},
		},
	}

	var resolver BotResolver
	for _, test := range testCases {
		test := test
		t.Run("decode telegram invitation username "+test.username, func(t *testing.T) {
			expPayload, err := resolver.decodeJoinLink(test.username)
			assert.NoError(t, err)
			assert.Equal(t, test.expPayload, expPayload)
		})
	}
}

type BotAPIMock struct {
	getChatMembersCount func(config tgbotapi.ChatConfig) (int, error)
	getChat             func(config tgbotapi.ChatConfig) (tgbotapi.Chat, error)
}

func (b BotAPIMock) GetChatMembersCount(config tgbotapi.ChatConfig) (int, error) {
	return b.getChatMembersCount(config)
}

func (b BotAPIMock) GetChat(config tgbotapi.ChatConfig) (tgbotapi.Chat, error) {
	return b.getChat(config)
}
