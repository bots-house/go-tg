package postgres

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/bots-house/birzzha/core"
	"github.com/bots-house/birzzha/store/postgres/dal"
	"github.com/bots-house/birzzha/store/postgres/shared"
	"github.com/pkg/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type LotStore struct {
	boil.ContextExecutor
}

func (store *LotStore) toRow(lot *core.Lot) (*dal.Lot, error) {
	extra, err := json.Marshal(lot.ExtraResources)
	if err != nil {
		return nil, errors.Wrap(err, "marshal extra resources")
	}

	return &dal.Lot{
		ID:                    int(lot.ID),
		OwnerID:               int(lot.OwnerID),
		ExternalID:            lot.ExternalID,
		Name:                  lot.Name,
		Avatar:                lot.Avatar,
		Username:              lot.Username,
		JoinLink:              lot.JoinLink,
		PriceCurrent:          lot.Price.Current,
		PricePrevious:         lot.Price.Previous,
		PriceIsBargain:        lot.Price.IsBargain,
		Comment:               lot.Comment,
		MetricsMembersCount:   lot.Metrics.MembersCount,
		MetricsDailyCoverage:  lot.Metrics.DailyCoverage,
		MetricsMonthlyIncome:  lot.Metrics.MonthlyIncome,
		MetricsPricePerMember: lot.Metrics.PricePerMember,
		MetricsPricePerView:   lot.Metrics.PricePerView,
		MetricsPaybackPeriod:  lot.Metrics.PaybackPeriod,
		ExtraResources:        null.JSONFrom(extra),
		CreatedAt:             lot.CreatedAt,
		PaidAt:                lot.PaidAt,
		ApprovedAt:            lot.ApprovedAt,
		PublishedAt:           lot.PublishedAt,
	}, nil
}

func (store *LotStore) fromRow(row *dal.Lot) (*core.Lot, error) {
	var extra []core.LotExtraResource

	if err := row.ExtraResources.Unmarshal(&extra); err != nil {
		return nil, errors.Wrap(err, "unmarshal extra resources")
	}

	return &core.Lot{
		ID:         core.LotID(row.ID),
		OwnerID:    core.UserID(row.OwnerID),
		ExternalID: row.ExternalID,
		Name:       row.Name,
		Avatar:     row.Avatar,
		Username:   row.Username,
		JoinLink:   row.JoinLink,
		Price: core.LotPrice{
			Current:   row.PriceCurrent,
			Previous:  row.PricePrevious,
			IsBargain: row.PriceIsBargain,
		},
		Comment: row.Comment,
		Metrics: core.LotMetrics{
			MembersCount:   row.MetricsMembersCount,
			DailyCoverage:  row.MetricsDailyCoverage,
			MonthlyIncome:  row.MetricsMonthlyIncome,
			PricePerMember: row.MetricsPricePerMember,
			PricePerView:   row.MetricsPricePerView,
			PaybackPeriod:  row.MetricsPaybackPeriod,
		},
		ExtraResources: extra,
		CreatedAt:      row.CreatedAt,
		PaidAt:         row.PaidAt,
		ApprovedAt:     row.ApprovedAt,
		PublishedAt:    row.PublishedAt,
	}, nil
}

func (store *LotStore) Add(ctx context.Context, lot *core.Lot) error {
	row, err := store.toRow(lot)
	if err != nil {
		return errors.Wrap(err, "to row")
	}

	if err := row.Insert(ctx, shared.GetExecutorOrDefault(ctx, store.ContextExecutor), boil.Infer()); err != nil {
		return errors.Wrap(err, "insert query")
	}

	lot.ID = core.LotID(row.ID)

	return nil
}

func (store *LotStore) Update(ctx context.Context, lot *core.Lot) error {
	row, err := store.toRow(lot)
	if err != nil {
		return errors.Wrap(err, "to row")
	}

	n, err := row.Update(ctx, shared.GetExecutorOrDefault(ctx, store.ContextExecutor), boil.Infer())
	if err != nil {
		return errors.Wrap(err, "update query")
	}
	if n == 0 {
		return core.ErrLotNotFound
	}
	return nil
}

func (store *LotStore) Query() core.LotStoreQuery {
	return &LotStoreQuery{store: store}
}

type LotStoreQuery struct {
	mods  []qm.QueryMod
	store *LotStore
}

func (usq *LotStoreQuery) ID(ids ...core.LotID) core.LotStoreQuery {
	idsInt := make([]int, len(ids))
	for i, id := range ids {
		idsInt[i] = int(id)
	}

	usq.mods = append(usq.mods, dal.LotWhere.ID.IN(idsInt))
	return usq
}

func (usq *LotStoreQuery) One(ctx context.Context) (*core.Lot, error) {
	executor := shared.GetExecutorOrDefault(ctx, usq.store.ContextExecutor)

	row, err := dal.Lots(usq.mods...).One(ctx, executor)
	if err == sql.ErrNoRows {
		return nil, core.ErrTopicNotFound
	} else if err != nil {
		return nil, err
	}

	return usq.store.fromRow(row)
}

func (usq *LotStoreQuery) All(ctx context.Context) (core.LotSlice, error) {
	executor := shared.GetExecutorOrDefault(ctx, usq.store.ContextExecutor)

	rows, err := dal.Lots(usq.mods...).All(ctx, executor)
	if err == sql.ErrNoRows {
		return nil, core.ErrTopicNotFound
	} else if err != nil {
		return nil, err
	}

	result := make(core.LotSlice, len(rows))
	for i, row := range rows {
		result[i], err = usq.store.fromRow(row)
		if err != nil {
			return nil, errors.Wrap(err, "from row")
		}
	}

	return result, nil
}
