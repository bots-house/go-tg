package catalog

import (
	"context"

	"github.com/davecgh/go-spew/spew"
	"github.com/pkg/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/bots-house/birzzha/core"
	"github.com/bots-house/birzzha/store"
)

type LotInput struct {
	Query         string
	TelegramID    int64
	TopicIDs      []core.TopicID
	Price         int
	IsBargain     bool
	MonthlyIncome int
	Comment       string
	Extra         []string
}

var (
	ErrLotIsNotChannel = core.NewError(
		"lot_is_not_channel",
		"lot is not channel, only channels is supported",
	)
)

type OwnedLot struct {
	*core.Lot
}

func (srv *Service) newOwnedLot(ctx context.Context, lot *core.Lot) (*OwnedLot, error) {
	return &OwnedLot{
		Lot:    lot,
	}, nil
}

func (srv *Service) AddLot(ctx context.Context, user *core.User, in *LotInput) (*OwnedLot, error) {
	result, err := srv.Resolver.ResolveByID(ctx, in.TelegramID)
	if err != nil {
		return nil, errors.Wrap(err, "resolve by id")
	}

	info := result.Channel

	if info == nil {
		return nil, ErrLotIsNotChannel
	}

	price := core.NewLotPrice(in.Price, in.IsBargain)

	lot := core.NewLot(
		user.ID,
		in.TelegramID,
		info.Name,
		price,
		in.Comment,
		info.MembersCount,
		info.DailyCoverage,
		null.NewInt(in.MonthlyIncome, in.MonthlyIncome != 0),
	)

	lot.Bio = null.NewString(info.Description, info.Description != "")

	if info.Username == "" {
		lot.JoinLink = null.NewString(in.Query, in.Query != "")
	} else {
		lot.Username = null.NewString(info.Username, info.Username != "")
	}

	if info.Avatar != "" {
		avatar, err := srv.Storage.AddByURL(ctx, "lot", info.Avatar)
		if err != nil {
			return nil, errors.Wrap(err, "add by url")
		}
		lot.Avatar = null.StringFrom(avatar)
	}

	lot.ExtraResources = make([]core.LotExtraResource, len(in.Extra))
	for i, v := range in.Extra {
		lot.ExtraResources[i] = core.LotExtraResource{URL: v}
	}

	lot.TopicIDs = in.TopicIDs

	if err := srv.Lot.Add(ctx, lot); err != nil {
		return nil, errors.Wrap(err, "add lot to store")
	}

	return srv.newOwnedLot(ctx, lot)
}

func (srv *Service) GetFilterBoundaries(ctx context.Context, query *core.LotFilterBoundariesQuery) (*core.LotFilterBoundaries, error) {
	return srv.Lot.FilterBoundaries(ctx, query)
}

type LotsQuery struct {
	Topics []core.TopicID
	MembersCountFrom int
	MembersCountTo int
	PriceFrom int
	PriceTo int
	PricePerMemberFrom float64
	PricePerMemberTo float64
	DailyCoverageFrom int
	DailyCoverageTo int
	PricePerViewFrom float64
	PricePerViewTo float64
	MonthlyIncomeFrom int
	MonthlyIncomeTo int
	PaybackPeriodFrom float64
	PaybackPeriodTo float64
	SortBy core.LotField
	SortByType store.SortType
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

func (srv *Service) GetLots(ctx context.Context, query *LotsQuery) ([]*ItemLot, error) {
	qry := srv.Lot.Query()

	spew.Dump(query)

	ctx = boil.WithDebug(ctx, true)

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

	lots, err := qry.All(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get lots")
	}

	return srv.newItemLotSlice(ctx, lots)
}
