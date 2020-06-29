package catalog

import (
	"context"

	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/bots-house/birzzha/core"
	"github.com/bots-house/birzzha/store"
)

func (srv *Service) GetFilterBoundaries(ctx context.Context, query *core.LotFilterBoundariesQuery) (*core.LotFilterBoundaries, error) {
	return srv.Lot.FilterBoundaries(ctx, query)
}

type LotsQuery struct {
	Topics             []core.TopicID
	MembersCountFrom   int
	MembersCountTo     int
	PriceFrom          int
	PriceTo            int
	PricePerMemberFrom float64
	PricePerMemberTo   float64
	DailyCoverageFrom  int
	DailyCoverageTo    int
	PricePerViewFrom   float64
	PricePerViewTo     float64
	MonthlyIncomeFrom  int
	MonthlyIncomeTo    int
	PaybackPeriodFrom  float64
	PaybackPeriodTo    float64
	SortBy             core.LotField
	SortByType         store.SortType
}

type ItemLot struct {
	*core.Lot

	Topics []core.TopicID
}

func (srv *Service) newItemLotSlice(ctx context.Context, lots core.LotSlice) ([]*ItemLot, error) {
	result := make([]*ItemLot, len(lots))

	for i, v := range lots {
		result[i] = &ItemLot{Lot: v}
	}

	return result, nil
}

func (srv *Service) newLotsQuery(ctx context.Context, query *LotsQuery) core.LotStoreQuery {
	qry := srv.Lot.Query()

	if query != nil {
		if len(query.Topics) > 0 {
			qry = qry.TopicIDs(query.Topics...)
		}

		if query.PriceFrom != 0 {
			qry = qry.PriceFrom(query.PriceFrom)
		}

		if query.PriceTo != 0 {
			qry = qry.PriceTo(query.PriceTo)
		}

		if query.MembersCountTo != 0 {
			qry = qry.MembersCountTo(query.MembersCountTo)
		}

		if query.MembersCountFrom != 0 {
			qry = qry.MembersCountFrom(query.MembersCountFrom)
		}

		if query.PricePerMemberFrom != 0 {
			qry = qry.PricePerMemberFrom(query.PricePerMemberFrom)
		}

		if query.PricePerMemberTo != 0 {
			qry = qry.PricePerMemberTo(query.PricePerMemberTo)
		}

		if query.DailyCoverageFrom != 0 {
			qry = qry.DailyCoverageFrom(query.DailyCoverageFrom)
		}

		if query.DailyCoverageTo != 0 {
			qry = qry.DailyCoverageTo(query.DailyCoverageTo)
		}

		if query.PricePerViewFrom != 0 {
			qry = qry.PricePerViewFrom(query.PricePerViewFrom)
		}

		if query.PricePerViewTo != 0 {
			qry = qry.PricePerViewTo(query.PricePerViewTo)
		}

		if query.MonthlyIncomeFrom != 0 {
			qry = qry.MonthlyIncomeFrom(query.MonthlyIncomeFrom)
		}

		if query.MonthlyIncomeTo != 0 {
			qry = qry.MonthlyIncomeTo(query.MonthlyIncomeTo)
		}

		if query.PaybackPeriodFrom != 0 {
			qry = qry.PaybackPeriodFrom(query.PaybackPeriodFrom)
		}

		if query.PaybackPeriodTo != 0 {
			qry = qry.PaybackPeriodTo(query.PaybackPeriodTo)
		}

		if query.SortBy != 0 {
			qry = qry.SortBy(query.SortBy, query.SortByType)
		}
	}

	return qry
}

func (srv *Service) GetLots(ctx context.Context, query *LotsQuery) ([]*ItemLot, error) {
	qry := srv.
		newLotsQuery(ctx, query).
		Statuses(core.ShowLotStatus...)

	ctx = boil.WithDebug(ctx, true)

	lots, err := qry.All(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get lots")
	}

	return srv.newItemLotSlice(ctx, lots)
}
