package tg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func assertArgsEqual(t *testing.T, r *Request, args map[string]string) bool {
	return assert.Equal(t, args, r.args)
}

func TestNewRequest(t *testing.T) {
	r := NewRequest("getMe")

	if assert.NotNil(t, r) {
		assert.Equal(t, "getMe", r.Method())
	}
}

func TestRequestToken(t *testing.T) {
	r := NewRequest("getMe")

	if assert.NotNil(t, r) {
		assert.Equal(t, "", r.Token())

		r.SetToken("1234:test")

		assert.Equal(t, "1234:test", r.Token())
	}
}

func TestRequest_SetString(t *testing.T) {
	r := NewRequest("sendMessage")

	r.SetString("string", "test")

	assertArgsEqual(t, r, map[string]string{
		"string": "test",
	})
}

func TestRequest_SetOptString(t *testing.T) {
	r := NewRequest("sendMessage")

	r.SetOptString("string", "test")
	r.SetOptString("skip", "")

	assertArgsEqual(t, r, map[string]string{
		"string": "test",
	})
}
