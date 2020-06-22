package authz

import (
	"github.com/bots-house/birzzha/core"
	"github.com/cristalhq/jwt"
)

type Identity struct {
	Token *jwt.Token
	User  *core.User
}
