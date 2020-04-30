package tg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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