package admin

import (
	"github.com/bots-house/birzzha/core"
	"github.com/bots-house/birzzha/pkg/storage"
	"github.com/bots-house/birzzha/pkg/tg"
)

type Service struct {
	Review            core.ReviewStore
	User              core.UserStore
	Lot               core.LotStore
	Settings          core.SettingsStore
	Topic             core.TopicStore
	LotTopic          core.LotTopicStore
	LotFile           core.LotFileStore
	LotCanceledReason core.LotCanceledReasonStore
	BotLinkBuilder    *core.BotLinkBuilder

	Storage        storage.Storage
	AvatarResolver tg.AvatarResolver
}

var (
	ErrUserIsNotAdmin = core.NewError("user_is_not_admin", "user is not admin")
)

func (srv *Service) IsAdmin(user *core.User) error {
	if !user.IsAdmin {
		return ErrUserIsNotAdmin
	}
	return nil
}
