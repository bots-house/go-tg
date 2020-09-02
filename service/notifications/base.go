package notifications

import "github.com/bots-house/birzzha/core"

type Notification interface {
	NotificationTemplate() string
}

type userNotificationWrapper struct {
	Notification
	userID core.UserID
}

func (unw *userNotificationWrapper) UserID() core.UserID {
	return unw.userID
}
