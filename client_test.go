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
		clien_t := NewClient(token, WithDoer(doer))

		if assert.NotNil(t, clien_t) {
			assert.Equal(t, token, clien_t.token)
			assert.Equal(t, doer, clien_t.doer)
		}
	})
}
