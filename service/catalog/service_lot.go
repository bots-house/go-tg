package catalog

import (
	"context"
	"fmt"

	"github.com/pkg/errors"

	"github.com/bots-house/birzzha/core"
	"github.com/bots-house/birzzha/pkg/tg"
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
	Offset             int
	Limit              int
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

func (srv *Service) newListLotsQuery(ctx context.Context, query *LotsQuery, qry core.LotStoreQuery) core.LotStoreQuery {
	if query != nil {

		if query.SortBy != 0 {
			qry = qry.SortBy(query.SortBy, query.SortByType)
		}

		if query.Offset != 0 {
			qry = qry.Offset(query.Offset)
		}

		if query.Limit != 0 {
			qry = qry.Limit(query.Limit)
		}
	}

	return qry
}

func (srv *Service) newBaseLotsQuery(ctx context.Context, query *LotsQuery) core.LotStoreQuery {
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
	}

	return qry
}

type LotList struct {
	Total int
	Items []*ItemLot
}

func (srv *Service) newLotList(ctx context.Context, total int, lots []*core.Lot) (*LotList, error) {
	items, err := srv.newItemLotSlice(ctx, lots)
	if err != nil {
		return nil, errors.Wrap(err, "to item lot")
	}

	return &LotList{
		Items: items,
		Total: total,
	}, nil
}

func (srv *Service) GetLots(ctx context.Context, query *LotsQuery) (*LotList, error) {
	qry := srv.
		newBaseLotsQuery(ctx, query).
		Statuses(core.ShowLotStatus...)

	total, err := qry.Count(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get total")
	}

	qry = srv.newListLotsQuery(ctx, query, qry)

	lots, err := qry.All(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get lots")
	}

	return srv.newLotList(ctx, total, lots)
}

type FullLot struct {
	ItemLot *ItemLot
	User    *core.User
	Views   int
}

func (fl *FullLot) TgstatLink() string {
	if fl.ItemLot.Username.Valid {
		return fmt.Sprintf("https://tgstat.ru/channel/@%s", fl.ItemLot.Username.String)
	}
	_, value := tg.ParseResolveQuery(fl.ItemLot.JoinLink.String)
	return fmt.Sprintf("https://tgstat.ru/channel/%s", value)
}

func (fl *FullLot) TelemetrLink() string {
	if fl.ItemLot.Username.Valid {
		return fmt.Sprintf("https://telemetr.me/@%s", fl.ItemLot.Username.String)
	}
	_, value := tg.ParseResolveQuery(fl.ItemLot.JoinLink.String)
	return fmt.Sprintf("https://telemetr.me/joinchat/%s", value)
}

func (srv *Service) GetLot(ctx context.Context, id core.LotID) (*FullLot, error) {
	lot, err := srv.Lot.Query().ID(id).One(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get lot")
	}

	user, err := srv.User.Query().ID(lot.OwnerID).One(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get user")
	}
	return &FullLot{
		ItemLot: &ItemLot{
			lot,
			lot.TopicIDs,
		},
		User:  user,
		Views: 0,
	}, nil
}
