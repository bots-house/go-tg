package models

import (
	"github.com/bots-house/birzzha/api/gen/models"
	"github.com/bots-house/birzzha/pkg/storage"
	authsrv "github.com/bots-house/birzzha/service/auth"
	"github.com/go-openapi/swag"
)

func NewToken(s storage.Storage, token *authsrv.Token) *models.Token {
	return &models.Token{
		User:  NewUser(s, token.User),
		Token: swag.String(token.Token),
	}
}
