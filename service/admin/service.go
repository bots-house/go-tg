package admin

import (
	"github.com/bots-house/birzzha/core"
	"github.com/bots-house/birzzha/pkg/storage"
	"github.com/bots-house/birzzha/pkg/tg"
	"github.com/bots-house/birzzha/service/notifications"
	"github.com/bots-house/birzzha/service/posting"
	"github.com/bots-house/birzzha/store"
	tgbotapi "github.com/bots-house/telegram-bot-api"
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
	Landing           core.LandingStore
	Favorite          core.FavoriteStore
	Post              core.PostStore
	Coupon            core.CouponStore
	CouponApply       core.CouponApplyStore
	BotLinkBuilder    *core.BotLinkBuilder
	Posting           *posting.Service

	Notify *notifications.Notifications

	Txier store.Txier

	Storage        storage.Storage
	AvatarResolver tg.AvatarResolver
	TgClient       *tgbotapi.BotAPI
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
