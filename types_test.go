package tg

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserBot_MarshalJSON(t *testing.T) {
	data := `{
	  "id": 12345,
	  "is_bot": true,
	  "first_name": "Test Bot",
	  "username": "test_bot",
	  "can_join_groups": true,
	  "can_read_all_group_messages": true,
	  "supports_inline_queries": true
	}`

	botUser := UserBot{}

	err := json.Unmarshal([]byte(data), &botUser)

	if assert.NoError(t, err) {
		assert.Equal(t, UserBot{
			User: User{
				ID:        12345,
				IsBot:     true,
				FirstName: "Test Bot",
				LastName:  "",
				Username:  "test_bot",
			},
			CanJoinGroups:           true,
			CanReadAllGroupMessages: true,
			SupportsInlineQueries:   true,
		}, botUser)
	}
}

func TestMessage_IsCommand(t *testing.T) {
	for _, tt := range []struct {
		Msg    Message
		Result bool
	}{
		{
			Message{
				Text: "/start",
			},
			true,
		},
		{
			Message{
				Text: "/command fsdklfdgd",
			},
			true,
		},
		{
			Message{
				Text: "fsdfsdfsd",
			},
			false,
		},
	} {
		r := tt.Msg.IsCommand()
		assert.Equal(t, tt.Result, r)
	}
}

func TestMessage_CommandArgs(t *testing.T) {
	for _, tt := range []struct {
		Msg    Message
		Result string
	}{
		{
			Message{
				Text: "/start",
			},
			"",
		},
		{
			Message{
				Text: "/start arg1 arg2",
			},
			"arg1 arg2",
		},
		{
			Message{
				Text: "fasdfsdfsd",
			},
			"",
		},
	} {
		args := tt.Msg.CommandArgs()
		assert.Equal(t, tt.Result, args)
	}
}
