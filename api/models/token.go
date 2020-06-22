package models

import (
	"github.com/bots-house/birzzha/api/gen/models"
	authsrv "github.com/bots-house/birzzha/service/auth"
	"github.com/go-openapi/swag"
)

func NewToken(token *authsrv.Token) *models.Token {
	return &models.Token{
		User:  NewUser(token.User),
		Token: swag.String(token.Token),
	}
}
