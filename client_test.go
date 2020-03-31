package tg

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	const (
		token = "12345:secret"
	)

	t.Run("WithoutOptions", func(t *testing.T) {
		client := NewClient(token)

		if assert.NotNil(t, client) {
			assert.Equal(t, token, client.token)
			assert.Equal(t, http.DefaultClient, client.doer)
		}
	})

	t.Run("WithOptions", func(t *testing.T) {
		doer := &http.Client{}
		client := NewClient(token, WithDoer(doer))

		if assert.NotNil(t, client) {
			assert.Equal(t, token, client.token)
			assert.Equal(t, doer, client.doer)
		}
	})
}
