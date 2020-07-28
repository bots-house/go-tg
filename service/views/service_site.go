package views

import (
	"context"
	"fmt"

	"github.com/bots-house/birzzha/core"
	"github.com/bots-house/birzzha/pkg/kv"
	"github.com/pkg/errors"
)

type SiteVisitor struct {
	UserID      core.UserID
	AnonymousID string
}

func (v SiteVisitor) IsAuthorized() bool {
	return v.UserID != 0 && v.AnonymousID == ""
}

type SiteView struct {
	Visitor *SiteVisitor
	LotID   core.LotID
}

func NewAuthorizedView(lot core.LotID, user core.UserID) *SiteView {
	return &SiteView{
		Visitor: &SiteVisitor{UserID: user},
		LotID:   lot,
	}
}

func NewAnonymousView(lot core.LotID, id string) *SiteView {
	return &SiteView{
		Visitor: &SiteVisitor{AnonymousID: id},
		LotID:   lot,
	}
}

func (srv *Service) getViewKey(view *SiteView) string {
	if view.Visitor.IsAuthorized() {
		return fmt.Sprintf("users:%d:lots:%d:view", view.Visitor.UserID, view.LotID)
	} else {
		return fmt.Sprintf("anonymous:%s:lots:%d:view", view.Visitor.AnonymousID, view.LotID)
	}
}

func (srv *Service) registerSiteView(ctx context.Context, key string, lot core.LotID) error {
	return srv.Txier(ctx, func(ctx context.Context) error {
		if err := srv.Lot.IncreaseSiteViews(ctx, lot); err != nil {
			return errors.Wrap(err, "increase views count")
		}

		if err := srv.KV.Set(ctx, key, true, kv.Expiration(srv.SiteViewExpiration)); err != nil {
			return errors.Wrap(err, "register view in redis")
		}

		return nil
	})
}

func (srv *Service) RegisterSiteView(ctx context.Context, view *SiteView) error {
	key := srv.getViewKey(view)

	err := srv.KV.Get(ctx, key, nil)
	if err == kv.ErrKeyNotFound {
		if err := srv.registerSiteView(ctx, key, view.LotID); err != nil {
			return errors.Wrap(err, "register lot view")
		}
	} else if err != nil {
		return errors.Wrap(err, "query kv")
	}

	return nil
}
