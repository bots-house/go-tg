package authz

import (
	"github.com/bots-house/birzzha/core"
	"github.com/cristalhq/jwt"
)

type Identity struct {
	Token *jwt.Token
	User  *core.User
}

func (i *Identity) IsAnonymous() bool {
	return i == nil || i.User == nil
}

func (i *Identity) GetUser() *core.User {
	if i != nil {
		return i.User
	}
	return nil
}
