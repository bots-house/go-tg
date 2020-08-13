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

type FullTopic struct {
	*core.Topic
	Lots int
}

type FullSettings struct {
	*core.Settings
	Landing         *core.Landing
	Topics          []*FullTopic
	CanceledReasons core.LotCanceledReasonSlice
}

func (srv *Service) newFullSettings(ctx context.Context, settings *core.Settings) (*FullSettings, error) {
	topics, err := srv.Topic.Query().All(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get topics")
	}

	canceledReasons, err := srv.LotCanceledReason.Query().All(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get canceled reasons")
	}

	fullTopics := make([]*FullTopic, len(topics))
	for i, topic := range topics {
		lots, err := srv.LotTopic.Query().TopicID(topic.ID).Count(ctx)
		if err != nil {
			return nil, errors.Wrap(err, "get lots count")
		}
		fullTopics[i] = &FullTopic{
			Topic: topic,
			Lots:  lots,
		}
	}

	landing, err := srv.Landing.Get(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get landing")
	}

	return &FullSettings{
		Settings:        settings,
		Topics:          fullTopics,
		CanceledReasons: canceledReasons,
		Landing:         landing,
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

type SettingsPricesInput struct {
	Application *money.Money
	Change      *money.Money
	Cashier     string
}

type SettingsGarantInput struct {
	Name                           string
	Username                       string
	ReviewsChannel                 string
	Avatar                         string
	PercentageDealOfDiscountPeriod float64
	PercentageDeal                 float64
}

type SettingsChannelInput struct {
	PrivateID      int64
	PublicUsername string
	PrivateLink    string
	Landing        SettingsInputLanding
}

type SettingsInputLanding struct {
	UniqueUsersPerMonthShift int
	AvgSiteReachShift        int
	AvgChannelReachShift     int
}

func (srv *Service) UpdateSettingsLanding(ctx context.Context, user *core.User, input *SettingsInputLanding) (*core.Landing, error) {
	if err := srv.IsAdmin(user); err != nil {
		return nil, err
	}

	landing, err := srv.Landing.Get(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get landing")
	}

	landing.AvgChannelReachShift = input.AvgChannelReachShift
	landing.AvgSiteReachShift = input.AvgSiteReachShift
	landing.UniqueUsersPerMonthShift = input.UniqueUsersPerMonthShift

	if err := srv.Landing.Update(ctx, landing); err != nil {
		return nil, errors.Wrap(err, "update landing")
	}

	return landing, err
}

func (srv *Service) UpdateSettingsPrice(ctx context.Context, user *core.User, input *SettingsPricesInput) (*core.Settings, error) {
	if err := srv.IsAdmin(user); err != nil {
		return nil, err
	}

	settings, err := srv.Settings.Get(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get settings")
	}

	settings.Prices.Application = input.Application
	settings.Prices.Change = input.Change

	settings.CashierUsername = input.Cashier

	if err := srv.Settings.Update(ctx, settings); err != nil {
		return nil, errors.Wrap(err, "update settings")
	}
	return settings, nil
}

func (srv *Service) UpdateSettingsGarant(ctx context.Context, user *core.User, input *SettingsGarantInput) (*core.Settings, error) {
	if err := srv.IsAdmin(user); err != nil {
		return nil, err
	}

	settings, err := srv.Settings.Get(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get settings")
	}

	settings.Garant.ReviewsChannel = input.ReviewsChannel
	settings.Garant.Name = input.Name
	settings.Garant.AvatarURL = null.NewString(input.Avatar, input.Avatar != "")
	settings.Garant.Username = input.Username
	settings.Garant.PercentageDeal = input.PercentageDeal
	settings.Garant.PercentageDealDiscountPeriod = null.NewFloat64(input.PercentageDealOfDiscountPeriod, input.PercentageDealOfDiscountPeriod != 0)

	settings.UpdatedAt = null.TimeFrom(time.Now())
	settings.UpdatedBy = user.ID

	if err := srv.Settings.Update(ctx, settings); err != nil {
		return nil, errors.Wrap(err, "update settings")
	}
	return settings, nil
}

func (srv *Service) UpdateSettingsChannel(ctx context.Context, user *core.User, input *SettingsChannelInput) (*core.Settings, error) {
	if err := srv.IsAdmin(user); err != nil {
		return nil, err
	}

	settings, err := srv.Settings.Get(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get settings")
	}

	settings.UpdatedAt = null.TimeFrom(time.Now())
	settings.UpdatedBy = user.ID
	settings.Channel.PrivateID = input.PrivateID
	settings.Channel.PrivateLink = input.PrivateLink
	settings.Channel.PublicUsername = input.PublicUsername

	settings.UpdatedAt = null.TimeFrom(time.Now())
	settings.UpdatedBy = user.ID

	if err := srv.Settings.Update(ctx, settings); err != nil {
		return nil, errors.Wrap(err, "update settings")
	}
	return settings, nil
}

func (srv *Service) CreateTopic(ctx context.Context, user *core.User, name string) (*FullTopic, error) {
	if err := srv.IsAdmin(user); err != nil {
		return nil, err
	}

	topic := core.NewTopic(name)
	if err := srv.Topic.Add(ctx, topic); err != nil {
		return nil, errors.Wrap(err, "create topic")
	}
	return &FullTopic{
		Topic: topic,
		Lots:  0,
	}, nil
}

func (srv *Service) UpdateTopic(ctx context.Context, user *core.User, id core.TopicID, name string) (*FullTopic, error) {
	if err := srv.IsAdmin(user); err != nil {
		return nil, err
	}

	topic, err := srv.Topic.Query().ID(id).One(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get topics")
	}

	topic.Name = name
	topic.Slug = slug.Make(name)
	if err := srv.Topic.Update(ctx, topic); err != nil {
		return nil, errors.Wrap(err, "updaet topic")
	}

	lots, err := srv.LotTopic.Query().TopicID(topic.ID).Count(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get lots count")
	}

	return &FullTopic{
		Topic: topic,
		Lots:  lots,
	}, nil
}

func (srv *Service) DeleteTopic(ctx context.Context, user *core.User, id core.TopicID) error {
	if err := srv.IsAdmin(user); err != nil {
		return err
	}

	if err := srv.Topic.Delete(ctx, id); err != nil {
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
