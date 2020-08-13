package postgres

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/bots-house/birzzha/core"
	"github.com/bots-house/birzzha/store"
	"github.com/bots-house/birzzha/store/postgres/dal"
	"github.com/bots-house/birzzha/store/postgres/shared"
)

type SettingsStore struct {
	db    *sql.DB
	txier store.Txier
}

func (store *SettingsStore) get(ctx context.Context) (*dal.Setting, error) {
	executor := shared.GetExecutorOrDefault(ctx, store.db)

	return dal.Settings().One(ctx, executor)
}

func (store *SettingsStore) fromRow(row *dal.Setting) (*core.Settings, error) {
	result := &core.Settings{}

	if err := row.PricesApplication.Unmarshal(&result.Prices.Application); err != nil {
		return nil, errors.Wrap(err, "unmarshal application price")
	}

	if err := row.PricesChange.Unmarshal(&result.Prices.Change); err != nil {
		return nil, errors.Wrap(err, "unmarshal price change")
	}

	result.Channel.PublicUsername = row.ChannelPublicUsername
	result.Channel.PrivateLink = row.ChannelPrivateLink
	result.Channel.PrivateID = row.ChannelPrivateID
	result.CashierUsername = row.CashierUsername
	result.UpdatedAt = row.UpdatedAt
	//garant
	result.Garant.Name = row.GarantName
	result.Garant.Username = row.GarantUsername
	result.Garant.AvatarURL = row.GarantAvatarURL
	result.Garant.ReviewsChannel = row.GarantReviewsChannel
	result.Garant.PercentageDeal = row.GarantPercentageDeal
	result.Garant.PercentageDealDiscountPeriod = row.GarantPercentageDealDiscountPeriod

	return result, nil
}

func (store *SettingsStore) toRow(settings *core.Settings) (*dal.Setting, error) {
	result := &dal.Setting{
		ID: 1,
	}

	if err := result.PricesApplication.Marshal(settings.Prices.Application); err != nil {
		return nil, errors.Wrap(err, "marshal application price")
	}

	if err := result.PricesChange.Marshal(settings.Prices.Change); err != nil {
		return nil, errors.Wrap(err, "marshal price change")
	}

	result.ChannelPublicUsername = settings.Channel.PublicUsername
	result.ChannelPrivateLink = settings.Channel.PrivateLink
	result.ChannelPrivateID = settings.Channel.PrivateID
	result.CashierUsername = settings.CashierUsername
	result.GarantAvatarURL = settings.Garant.AvatarURL
	result.GarantName = settings.Garant.Name
	result.GarantUsername = settings.Garant.Username
	result.GarantReviewsChannel = settings.Garant.ReviewsChannel
	result.GarantPercentageDeal = settings.Garant.PercentageDeal
	result.GarantPercentageDealDiscountPeriod = settings.Garant.PercentageDealDiscountPeriod

	result.UpdatedAt = settings.UpdatedAt

	return result, nil
}

func (store *SettingsStore) upsert(ctx context.Context, settings *core.Settings) error {
	executor := shared.GetExecutorOrDefault(ctx, store.db)

	row, err := store.toRow(settings)
	if err != nil {
		return errors.Wrap(err, "to row")
	}

	if err := row.Upsert(ctx, executor, true, []string{"id"}, boil.Infer(), boil.Infer()); err != nil {
		return errors.Wrap(err, "upsert")
	}

	return nil
}

func (store *SettingsStore) Get(ctx context.Context) (*core.Settings, error) {
	var settings *core.Settings

	if err := store.txier(ctx, func(ctx context.Context) error {
		row, err := store.get(ctx)
		if err == sql.ErrNoRows {
			if err := store.upsert(ctx, core.DefaultSettings); err != nil {
				return errors.Wrap(err, "insert default settings")
			}
			settings = core.DefaultSettings
			return nil
		} else if err != nil {
			return errors.Wrap(err, "query row")
		}

		settings, err = store.fromRow(row)
		if err != nil {
			return errors.Wrap(err, "from row")
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return settings, nil
}

func (store *SettingsStore) Update(ctx context.Context, settings *core.Settings) error {
	if err := store.upsert(ctx, settings); err != nil {
		return errors.Wrap(err, "upsert")
	}

	return nil
}
