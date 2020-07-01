package postgres

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/bots-house/birzzha/core"
	"github.com/bots-house/birzzha/store"
	"github.com/bots-house/birzzha/store/postgres/dal"
	"github.com/bots-house/birzzha/store/postgres/shared"
)

type LotStore struct {
	boil.ContextExecutor
	txier         store.Txier
	lotTopicStore *LotTopicStore
}

type filterBoundaries struct {
	PriceMin null.Int
	PriceMax null.Int

	MembersCountMin null.Int
	MembersCountMax null.Int

	DailyCoverageMin null.Int
	DailyCoverageMax null.Int

	MonthlyIncomeMin null.Int
	MonthlyIncomeMax null.Int

	PricePerMemberMin null.Float64
	PricePerMemberMax null.Float64

	PricePerViewMin null.Float64
	PricePerViewMax null.Float64

	PaybackPeriodMin null.Float64
	PaybackPeriodMax null.Float64
}

const queryFilterBoundariesAll = `
	select
		min(price_current) as price_min,
		max(price_current) as price_max,
		min(metrics_members_count) as members_count_min,
		max(metrics_members_count) as members_count_max,
		min(metrics_daily_coverage) as daily_coverage_min,
		max(metrics_daily_coverage) as daily_coverage_max,
		min(metrics_monthly_income) as monthly_income_min,
		max(metrics_monthly_income) as monthly_income_max,
		min(metrics_price_per_member) as price_per_member_min,
		max(metrics_price_per_member) as price_per_member_max,
		min(metrics_price_per_view) as price_per_view_min,
		max(metrics_price_per_view) as price_per_view_max,
		min(metrics_payback_period) as payback_period_min,
		max(metrics_payback_period) as payback_period_max
	from
		lot
`

const queryFilterBoudariesTopics = `
	select
		min(price_current) as price_min,
		max(price_current) as price_max,
		min(metrics_members_count) as members_count_min,
		max(metrics_members_count) as members_count_max,
		min(metrics_daily_coverage) as daily_coverage_min,
		max(metrics_daily_coverage) as daily_coverage_max,
		min(metrics_monthly_income) as monthly_income_min,
		max(metrics_monthly_income) as monthly_income_max,
		min(metrics_price_per_member) as price_per_member_min,
		max(metrics_price_per_member) as price_per_member_max,
		min(metrics_price_per_view) as price_per_view_min,
		max(metrics_price_per_view) as price_per_view_max,
		min(metrics_payback_period) as payback_period_min,
		max(metrics_payback_period) as payback_period_max
	from
		lot
	inner join lot_topic
		lt on lot.id = lt.lot_id
	where lt.topic_id = any($1)
`

func (store *LotStore) FilterBoundaries(ctx context.Context, query *core.LotFilterBoundariesQuery) (*core.LotFilterBoundaries, error) {
	exec := shared.GetExecutorOrDefault(ctx, store.ContextExecutor)

	result := &filterBoundaries{}

	var (
		args     = []interface{}{}
		sqlQuery = queryFilterBoundariesAll
	)

	if query != nil && len(query.Topics) > 0 {
		sqlQuery = queryFilterBoudariesTopics

		ids := make(pq.Int64Array, len(query.Topics))
		for i, v := range query.Topics {
			ids[i] = int64(v)
		}

		args = []interface{}{ids}
	}

	err := exec.QueryRowContext(ctx, sqlQuery, args...).Scan(
		&result.PriceMin,
		&result.PriceMax,
		&result.MembersCountMin,
		&result.MembersCountMax,
		&result.DailyCoverageMin,
		&result.DailyCoverageMax,
		&result.MonthlyIncomeMin,
		&result.MonthlyIncomeMax,
		&result.PricePerMemberMin,
		&result.PricePerMemberMax,
		&result.PricePerViewMin,
		&result.PricePerViewMax,
		&result.PaybackPeriodMin,
		&result.PaybackPeriodMax,
	)

	if err != nil {
		return nil, errors.Wrap(err, "query & scan")
	}

	return &core.LotFilterBoundaries{
		PriceMin:          result.PriceMin.Int,
		PriceMax:          result.PriceMax.Int,
		MembersCountMin:   result.MembersCountMin.Int,
		MembersCountMax:   result.MembersCountMax.Int,
		DailyCoverageMin:  result.DailyCoverageMin.Int,
		DailyCoverageMax:  result.DailyCoverageMax.Int,
		MonthlyIncomeMin:  result.MonthlyIncomeMin.Int,
		MonthlyIncomeMax:  result.MonthlyIncomeMax.Int,
		PricePerMemberMin: result.PricePerMemberMin.Float64,
		PricePerMemberMax: result.PricePerMemberMax.Float64,
		PricePerViewMin:   result.PricePerViewMin.Float64,
		PricePerViewMax:   result.PricePerViewMax.Float64,
		PaybackPeriodMin:  result.PaybackPeriodMin.Float64,
		PaybackPeriodMax:  result.PaybackPeriodMax.Float64,
	}, nil
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
		Bio:                   lot.Bio,
		Status:                lot.Status.String(),
		CanceledReasonID:      null.NewInt(int(lot.CanceledReasonID), lot.CanceledReasonID != 0),
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

	topics := make([]core.TopicID, len(row.R.LotTopics))

	for i, v := range row.R.LotTopics {
		topics[i] = core.TopicID(v.TopicID)
	}

	status, err := core.ParseLotStatus(row.Status)
	if err != nil {
		return nil, errors.Wrap(err, "parse status")
	}

	return &core.Lot{
		ID:               core.LotID(row.ID),
		OwnerID:          core.UserID(row.OwnerID),
		ExternalID:       row.ExternalID,
		Name:             row.Name,
		Avatar:           row.Avatar,
		Username:         row.Username,
		JoinLink:         row.JoinLink,
		Status:           status,
		CanceledReasonID: core.LotCanceledReasonID(row.CanceledReasonID.Int),
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
		Bio:            row.Bio,
		ExtraResources: extra,
		TopicIDs:       topics,
		CreatedAt:      row.CreatedAt,
		PaidAt:         row.PaidAt,
		ApprovedAt:     row.ApprovedAt,
		PublishedAt:    row.PublishedAt,
	}, nil
}

func (store *LotStore) Add(ctx context.Context, lot *core.Lot) error {
	return store.txier(ctx, func(ctx context.Context) error {
		return store.add(ctx, lot)
	})
}

func (store *LotStore) add(ctx context.Context, lot *core.Lot) error {
	row, err := store.toRow(lot)
	if err != nil {
		return errors.Wrap(err, "to row")
	}

	if err := row.Insert(ctx, shared.GetExecutorOrDefault(ctx, store.ContextExecutor), boil.Infer()); err != nil {
		return errors.Wrap(err, "insert query")
	}

	lot.ID = core.LotID(row.ID)

	if err := store.lotTopicStore.Set(ctx, lot.ID, lot.TopicIDs); err != nil {
		return errors.Wrap(err, "create lot topics")
	}

	return nil
}

func (store *LotStore) Update(ctx context.Context, lot *core.Lot) error {
	return store.txier(ctx, func(ctx context.Context) error {
		return store.update(ctx, lot)
	})
}

func (store *LotStore) update(ctx context.Context, lot *core.Lot) error {
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

	if err := store.lotTopicStore.Set(ctx, lot.ID, lot.TopicIDs); err != nil {
		return errors.Wrap(err, "create lot topics")
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

func (lsq *LotStoreQuery) OwnerID(id core.UserID) core.LotStoreQuery {
	lsq.mods = append(lsq.mods, dal.LotWhere.OwnerID.EQ(int(id)))
	return lsq
}

func (lsq *LotStoreQuery) SortBy(field core.LotField, typ store.SortType) core.LotStoreQuery {
	var orderBy string

	switch field {
	case core.LotFieldMembersCount:
		orderBy = dal.LotColumns.MetricsMembersCount
	case core.LotFieldPrice:
		orderBy = dal.LotColumns.PriceCurrent
	case core.LotFieldPricePerMember:
		orderBy = dal.LotColumns.MetricsPricePerMember
	case core.LotFieldDailyCoverage:
		orderBy = dal.LotColumns.MetricsDailyCoverage
	case core.LotFieldPricePerView:
		orderBy = dal.LotColumns.MetricsPricePerView
	case core.LotFieldMonthlyIncome:
		orderBy = dal.LotColumns.MetricsMonthlyIncome
	case core.LotFieldPaybackPeriod:
		orderBy = dal.LotColumns.MetricsPaybackPeriod
	case core.LotFieldCreatedAt:
		orderBy = dal.LotColumns.CreatedAt
	}

	switch typ {
	case store.SortTypeAsc:
		orderBy += " ASC"
	case store.SortTypeDesc:
		orderBy += " DESC"
	}

	lsq.mods = append(lsq.mods, qm.OrderBy(orderBy))

	return lsq
}

func (lsq *LotStoreQuery) PaybackPeriodFrom(v float64) core.LotStoreQuery {
	lsq.mods = append(lsq.mods, dal.LotWhere.MetricsPaybackPeriod.GTE(null.Float64From(v)))
	return lsq
}

func (lsq *LotStoreQuery) PaybackPeriodTo(v float64) core.LotStoreQuery {
	lsq.mods = append(lsq.mods, dal.LotWhere.MetricsPaybackPeriod.LTE(null.Float64From(v)))
	return lsq
}

func (lsq *LotStoreQuery) MonthlyIncomeFrom(v int) core.LotStoreQuery {
	lsq.mods = append(lsq.mods, dal.LotWhere.MetricsMonthlyIncome.GTE(null.IntFrom(v)))
	return lsq
}

func (lsq *LotStoreQuery) MonthlyIncomeTo(v int) core.LotStoreQuery {
	lsq.mods = append(lsq.mods, dal.LotWhere.MetricsMonthlyIncome.LTE(null.IntFrom(v)))
	return lsq
}

func (lsq *LotStoreQuery) PricePerViewFrom(v float64) core.LotStoreQuery {
	lsq.mods = append(lsq.mods, dal.LotWhere.MetricsPricePerView.GTE(v))
	return lsq
}

func (lsq *LotStoreQuery) PricePerViewTo(v float64) core.LotStoreQuery {
	lsq.mods = append(lsq.mods, dal.LotWhere.MetricsPricePerView.LTE(v))
	return lsq
}

func (lsq *LotStoreQuery) DailyCoverageFrom(v int) core.LotStoreQuery {
	lsq.mods = append(lsq.mods, dal.LotWhere.MetricsDailyCoverage.GTE(v))
	return lsq
}

func (lsq *LotStoreQuery) DailyCoverageTo(v int) core.LotStoreQuery {
	lsq.mods = append(lsq.mods, dal.LotWhere.MetricsDailyCoverage.LTE(v))
	return lsq
}

func (lsq *LotStoreQuery) PricePerMemberFrom(v float64) core.LotStoreQuery {
	lsq.mods = append(lsq.mods, dal.LotWhere.MetricsPricePerMember.GTE(v))
	return lsq
}

func (lsq *LotStoreQuery) PricePerMemberTo(v float64) core.LotStoreQuery {
	lsq.mods = append(lsq.mods, dal.LotWhere.MetricsPricePerMember.LTE(v))
	return lsq
}

func (lsq *LotStoreQuery) PriceFrom(v int) core.LotStoreQuery {
	lsq.mods = append(lsq.mods, dal.LotWhere.PriceCurrent.GTE(v))
	return lsq
}

func (lsq *LotStoreQuery) PriceTo(v int) core.LotStoreQuery {
	lsq.mods = append(lsq.mods, dal.LotWhere.PriceCurrent.LTE(v))
	return lsq
}

func (lsq *LotStoreQuery) MembersCountFrom(v int) core.LotStoreQuery {
	lsq.mods = append(lsq.mods, dal.LotWhere.MetricsMembersCount.GTE(v))
	return lsq
}

func (lsq *LotStoreQuery) MembersCountTo(v int) core.LotStoreQuery {
	lsq.mods = append(lsq.mods, dal.LotWhere.MetricsMembersCount.LTE(v))
	return lsq
}

func (lsq *LotStoreQuery) Statuses(statuses ...core.LotStatus) core.LotStoreQuery {
	vs := make([]string, len(statuses))

	for i, status := range statuses {
		vs[i] = status.String()
	}

	lsq.mods = append(lsq.mods, dal.LotWhere.Status.IN(vs))

	return lsq
}

func (lsq *LotStoreQuery) ID(ids ...core.LotID) core.LotStoreQuery {
	idsInt := make([]int, len(ids))
	for i, id := range ids {
		idsInt[i] = int(id)
	}

	lsq.mods = append(lsq.mods, dal.LotWhere.ID.IN(idsInt))
	return lsq
}

func (lsq *LotStoreQuery) TopicIDs(ids ...core.TopicID) core.LotStoreQuery {
	idsInt := make([]int64, len(ids))
	for i, id := range ids {
		idsInt[i] = int64(id)
	}

	lsq.mods = append(lsq.mods,
		qm.InnerJoin("lot_topic on lot.id = lot_topic.id"),
		qm.WhereIn("lot_topic.topic_id = any(?)", pq.Array(idsInt)),
	)
	return lsq
}

func (lsq *LotStoreQuery) Offset(v int) core.LotStoreQuery {
	lsq.mods = append(lsq.mods, qm.Offset(v))
	return lsq
}

func (lsq *LotStoreQuery) Limit(v int) core.LotStoreQuery {
	lsq.mods = append(lsq.mods, qm.Limit(v))
	return lsq
}

func (lsq *LotStoreQuery) eager() {
	lsq.mods = append(lsq.mods,
		qm.Load(dal.LotRels.LotTopics),
	)
}

func (lsq *LotStoreQuery) One(ctx context.Context) (*core.Lot, error) {
	lsq.eager()

	executor := shared.GetExecutorOrDefault(ctx, lsq.store.ContextExecutor)

	row, err := dal.Lots(lsq.mods...).One(ctx, executor)
	if err == sql.ErrNoRows {
		return nil, core.ErrTopicNotFound
	} else if err != nil {
		return nil, err
	}

	return lsq.store.fromRow(row)
}

func (lsq *LotStoreQuery) Count(ctx context.Context) (int, error) {
	executor := shared.GetExecutorOrDefault(ctx, lsq.store.ContextExecutor)

	count, err := dal.Lots(lsq.mods...).Count(ctx, executor)

	return int(count), err
}

func (lsq *LotStoreQuery) All(ctx context.Context) (core.LotSlice, error) {
	lsq.eager()

	executor := shared.GetExecutorOrDefault(ctx, lsq.store.ContextExecutor)

	rows, err := dal.Lots(lsq.mods...).All(ctx, executor)
	if err == sql.ErrNoRows {
		return nil, core.ErrTopicNotFound
	} else if err != nil {
		return nil, err
	}

	result := make(core.LotSlice, len(rows))
	for i, row := range rows {
		result[i], err = lsq.store.fromRow(row)
		if err != nil {
			return nil, errors.Wrap(err, "from row")
		}
	}

	return result, nil
}
