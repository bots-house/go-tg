//+build unit

package tg

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestGetQueryType(t *testing.T) {
	for _, test := range []struct {
		Input    string
		Excepted QueryType
		Value    string
	}{
		{"https://t.me/joinchat/AAAAAES_pid_l6flZONwGQ", QueryTypeJoinLink, "AAAAAES_pid_l6flZONwGQ"},
		{"zzap.run/joinchat/AAAAAES_pid_l6flZONwGQ", QueryTypeJoinLink, "AAAAAES_pid_l6flZONwGQ"},
		{"https://t.me/channely", QueryTypeUsername, "channely"},
		{"https://zzap.run/channely", QueryTypeUsername, "channely"},
		{"t.me/channely_bot", QueryTypeUsername, "channely_bot"},
		{"channely", QueryTypeUsername, "channely"},
		{"@channely", QueryTypeUsername, "channely"},
		{},
	} {
		typ, val := ParseResolveQuery(test.Input)

		assert.Equal(t, test.Excepted, typ)
		assert.Equal(t, test.Value, val)
	}

}
