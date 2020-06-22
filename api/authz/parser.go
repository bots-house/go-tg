package authz

import (
	"github.com/cristalhq/jwt"
	"github.com/pkg/errors"
)

func Parse(tkn string) (*Identity, error) {
	token, err := jwt.ParseString(tkn)
	if err != nil {
		return nil, errors.Wrap(err, "parse token")
	}
	return &Identity{
		Token: token,
	}, nil
}
