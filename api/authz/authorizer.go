package authz

import (
	"net/http"

	"github.com/bots-house/birzzha/pkg/sentrylog"
	"github.com/bots-house/birzzha/service/auth"
	"github.com/pkg/errors"
)

type Authorizer struct {
	Srv *auth.Service
}

func (authz *Authorizer) Authorize(r *http.Request, id interface{}) error {
	if id == nil {
		return nil
	}

	ctx := r.Context()
	identity := id.(*Identity)

	user, err := authz.Srv.Authorize(ctx, identity.Token)
	if err != nil {
		//todo: fix (https://github.com/bots-house/birzzha/issues/139) service returns 403 instead of 500 (code generator issue)
		return errors.Wrap(err, "authorize")
	}

	identity.User = user
	sentrylog.AddUser(ctx, int(user.ID), user.Telegram.Username.String)

	return nil
}
