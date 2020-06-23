package tg

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestGetQueryType(t *testing.T) {
	for _, test := range []struct {
		Input    string
		Excepted queryType
		Value    string
	}{
		{"https://t.me/joinchat/AAAAAES_pid_l6flZONwGQ", queryTypeJoinLink, "AAAAAES_pid_l6flZONwGQ"},
		{"zzap.run/joinchat/AAAAAES_pid_l6flZONwGQ", queryTypeJoinLink, "AAAAAES_pid_l6flZONwGQ"},
		{"https://t.me/channely", queryTypeUsername, "channely"},
		{"https://zzap.run/channely", queryTypeUsername, "channely"},
		{"t.me/channely_bot", queryTypeUsername, "channely_bot"},
		{"channely", queryTypeUsername, "channely"},
		{"@channely", queryTypeUsername, "channely"},
		{},
	} {
		typ, val := parseResolveQuery(test.Input)

		assert.Equal(t, test.Excepted, typ)
		assert.Equal(t, test.Value, val)
	}

}
