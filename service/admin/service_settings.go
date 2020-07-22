package admin

import (
	"context"
	"time"

	"github.com/Rhymond/go-money"
	"github.com/bots-house/birzzha/core"
	"github.com/gosimple/slug"
	"github.com/pkg/errors"
	"github.com/volatiletech/null/v8"
)

type FullSettings struct {
	*core.Settings
	Topics          core.TopicSlice
	CanceledReasons core.LotCanceledReasonSlice
}

func (srv *Service) newFullSettings(ctx context.Context, settings *core.Settings) (*FullSettings, error) {
	topics, err := srv.LotTopic.Query().All(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get topics")
	}

	canceledReasons, err := srv.LotCanceledReason.Query().All(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get canceled reasons")
	}

	return &FullSettings{
		Settings:        settings,
		Topics:          topics,
		CanceledReasons: canceledReasons,
	}, nil
}

func (srv *Service) GetSettings(ctx context.Context, user *core.User) (*FullSettings, error) {
	if err := srv.IsAdmin(user); err != nil {
		return nil, err
	}

	settings, err := srv.Settings.Get(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get settings")
	}

	return srv.newFullSettings(ctx, settings)
}

type SettingsPriceInput struct {
	Application *money.Money
	Change      *money.Money
	Cashier     string
}

func (srv *Service) UpdateSettingsPrice(ctx context.Context, user *core.User, input *SettingsPriceInput) (*FullSettings, error) {
	if err := srv.IsAdmin(user); err != nil {
		return nil, err
	}

	settings, err := srv.Settings.Get(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get settings")
	}

	settings.Prices.Application = input.Application
	settings.Prices.Change = input.Change
	settings.UpdatedAt = null.TimeFrom(time.Now())
	settings.UpdatedBy = user.ID
	settings.CashierUsername = input.Cashier

	if err := srv.Settings.Update(ctx, settings); err != nil {
		return nil, errors.Wrap(err, "update settings")
	}

	return srv.newFullSettings(ctx, settings)
}

func (srv *Service) CreateTopic(ctx context.Context, user *core.User, name string) (*core.Topic, error) {
	if err := srv.IsAdmin(user); err != nil {
		return nil, err
	}

	topic := core.NewTopic(name)
	if err := srv.LotTopic.Add(ctx, topic); err != nil {
		return nil, errors.Wrap(err, "create topic")
	}
	return topic, nil
}

func (srv *Service) UpdateTopic(ctx context.Context, user *core.User, id core.TopicID, name string) (*core.Topic, error) {
	if err := srv.IsAdmin(user); err != nil {
		return nil, err
	}

	topic, err := srv.LotTopic.Query().ID(id).One(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get topics")
	}

	topic.Name = name
	topic.Slug = slug.Make(name)
	if err := srv.LotTopic.Update(ctx, topic); err != nil {
		return nil, errors.Wrap(err, "updaet topic")
	}
	return topic, nil
}

func (srv *Service) DeleteTopic(ctx context.Context, user *core.User, id core.TopicID) error {
	if err := srv.IsAdmin(user); err != nil {
		return err
	}

	if err := srv.LotTopic.Delete(ctx, id); err != nil {
		return err
	}
	return nil
}

type LotCanceledReasonInput struct {
	Why      string
	IsPublic bool
}

func (srv *Service) CreateLotCanceledReason(ctx context.Context, user *core.User, input *LotCanceledReasonInput) (*core.LotCanceledReason, error) {
	if err := srv.IsAdmin(user); err != nil {
		return nil, err
	}

	reason := core.NewLotCanceledReason(input.Why, input.IsPublic)
	if err := srv.LotCanceledReason.Add(ctx, reason); err != nil {
		return nil, errors.Wrap(err, "add reason")
	}
	return reason, nil
}

func (srv *Service) UpdateLotCanceledReason(ctx context.Context, user *core.User, id core.LotCanceledReasonID, input *LotCanceledReasonInput) (*core.LotCanceledReason, error) {
	if err := srv.IsAdmin(user); err != nil {
		return nil, err
	}

	reason, err := srv.LotCanceledReason.Query().ID(id).One(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get lot canceled reason")
	}

	reason.Why = input.Why
	reason.IsPublic = input.IsPublic
	reason.UpdatedAt = null.TimeFrom(time.Now())

	if err := srv.LotCanceledReason.Update(ctx, reason); err != nil {
		return nil, errors.Wrap(err, "update lot canceled reason")
	}
	return reason, nil
}

func (srv *Service) DeleteLotCanceledReason(ctx context.Context, user *core.User, id core.LotCanceledReasonID) error {
	if err := srv.IsAdmin(user); err != nil {
		return err
	}

	if err := srv.LotCanceledReason.Delete(ctx, id); err != nil {
		return errors.Wrap(err, "delete lot canceled reason")
	}
	return nil
}
