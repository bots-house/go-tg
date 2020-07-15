package catalog

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/volatiletech/null/v8"

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
	InFavorites null.Bool
}

func (srv *Service) newItemLotSlice(lots core.LotSlice, favorites core.FavoriteSlice) []*ItemLot {
	result := make([]*ItemLot, len(lots))

	for i, v := range lots {
		result[i] = &ItemLot{
			Lot: v,
		}
		if favorites != nil {
			result[i].InFavorites = null.BoolFrom(favorites.HasLot(v.ID))
		}
	}

	return result
}

func (srv *Service) newListLotsQuery(query *LotsQuery, qry core.LotStoreQuery) core.LotStoreQuery {
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

func (srv *Service) newBaseLotsQuery(query *LotsQuery) core.LotStoreQuery {
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

func (srv *Service) newLotList(total int, lots core.LotSlice, favorites core.FavoriteSlice) (*LotList, error) {
	items := srv.newItemLotSlice(lots, favorites)

	return &LotList{
		Items: items,
		Total: total,
	}, nil
}

func (srv *Service) GetLots(ctx context.Context, user *core.User, query *LotsQuery) (*LotList, error) {
	qry := srv.
		newBaseLotsQuery(query).
		Statuses(core.ShowLotStatus...)

	total, err := qry.Count(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get total")
	}

	qry = srv.newListLotsQuery(query, qry)

	lots, err := qry.All(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get lots")
	}

	var favorites core.FavoriteSlice

	if user != nil {
		favorites, err = srv.Favorite.Query().UserID(user.ID).LotID(lots.IDs()...).All(ctx)
		if err != core.ErrFavoriteNotFound && err != nil {
			return nil, errors.Wrap(err, "get favorites")
		}
	}

	return srv.newLotList(total, lots, favorites)
}

type FullLot struct {
	*core.Lot
	InFavorites null.Bool
	User        *core.User
	Views       int
}

func (fl *FullLot) TgstatLink() string {
	if fl.Username.Valid {
		return fmt.Sprintf("https://tgstat.ru/channel/@%s", fl.Username.String)
	}
	_, value := tg.ParseResolveQuery(fl.JoinLink.String)
	return fmt.Sprintf("https://tgstat.ru/channel/%s", value)
}

func (fl *FullLot) TelemetrLink() string {
	if fl.Username.Valid {
		return fmt.Sprintf("https://telemetr.me/@%s", fl.Username.String)
	}
	_, value := tg.ParseResolveQuery(fl.JoinLink.String)
	return fmt.Sprintf("https://telemetr.me/joinchat/%s", value)
}

func (srv *Service) GetLot(ctx context.Context, user *core.User, id core.LotID) (*FullLot, error) {
	lot, err := srv.Lot.Query().ID(id).One(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get lot")
	}

	usr, err := srv.User.Query().ID(lot.OwnerID).One(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get user")
	}

	fullLot := &FullLot{
		Lot:   lot,
		User:  usr,
		Views: lot.Views.Total(),
	}

	if user != nil {

		favorite, err := srv.Favorite.Query().LotID(lot.ID).UserID(user.ID).One(ctx)
		if err != core.ErrFavoriteNotFound && err != nil {
			return nil, errors.Wrap(err, "get favorite")
		}

		if favorite != nil {
			fullLot.InFavorites = null.BoolFrom(true)
		}

	}
	return fullLot, nil

}

func (srv *Service) SimilarLots(ctx context.Context, id core.LotID, limit int, offset int) (*LotList, error) {
	var l int
	var o int

	if limit == 0 && offset == 0 {
		l = 10
		o = 0
	} else {
		l = limit
		o = offset
	}

	count, err := srv.Lot.SimilarLotsCount(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "similar lots count")
	}

	ids, err := srv.Lot.SimilarLotIDs(ctx, id, l, o)
	if err != nil {
		return nil, errors.Wrap(err, "similar lots")
	}

	lots, err := srv.Lot.Query().ID(ids...).All(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get lots")
	}

	sortedLots := lots.SortByID(ids)

	return srv.newLotList(count, sortedLots, nil)
}

func (srv *Service) ToggleLotFavorite(ctx context.Context, user *core.User, id core.LotID) (bool, error) {
	_, err := srv.Favorite.Query().LotID(id).UserID(user.ID).One(ctx)
	if err == core.ErrFavoriteNotFound {
		f := core.NewFavorite(
			id, user.ID,
		)
		if err := srv.Favorite.Add(ctx, f); err != nil {
			return false, errors.Wrap(err, "add favorite")
		}
		return true, nil
	} else if err != nil {
		return false, errors.Wrap(err, "get favorite")
	}

	if err := srv.Favorite.Query().LotID(id).UserID(user.ID).Delete(ctx); err != nil {
		return false, errors.Wrap(err, "delete favorite")
	}

	return false, nil
}
