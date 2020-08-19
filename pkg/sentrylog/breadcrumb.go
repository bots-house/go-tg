package sentrylog

import (
	"context"
	"strconv"

	"github.com/getsentry/sentry-go"
)

func AddUser(ctx context.Context, usrID int, usrTgUsrName string) {
	if sentry.HasHubOnContext(ctx) {
		hub := sentry.GetHubFromContext(ctx)
		hub.AddBreadcrumb(&sentry.Breadcrumb{
			Message: "User",
		}, nil)

		hub.ConfigureScope(func(scope *sentry.Scope) {
			scope.SetUser(sentry.User{
				ID:       strconv.Itoa(usrID),
				Username: usrTgUsrName,
			})
		})
	}
}
