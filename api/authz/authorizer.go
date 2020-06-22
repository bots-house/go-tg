package authz

import (
	"net/http"

	"github.com/bots-house/birzzha/service/auth"
	"github.com/pkg/errors"
)

type Authorizer struct {
	Srv *auth.Service
}

func (authz *Authorizer) Authorize(r *http.Request, id interface{}) error {
	ctx := r.Context()
	identity := id.(*Identity)

	user, err := authz.Srv.Authorize(ctx, identity.Token)
	if err != nil {
		return errors.Wrap(err, "authorize")
	}

	identity.User = user

	return nil
}
