package personal

import (
	"context"

	"github.com/pkg/errors"
	"github.com/volatiletech/null/v8"

	"github.com/bots-house/birzzha/core"
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

type OwnedLot struct {
	*core.Lot
}

var (
	ErrLotIsNotChannel = core.NewError(
		"lot_is_not_channel",
		"lot is not channel, only channels is supported",
	)
)

func (srv *Service) newOwnedLot(ctx context.Context, lot *core.Lot) (*OwnedLot, error) {
	return &OwnedLot{
		Lot: lot,
	}, nil
}

func (srv *Service) newOwnedLotSlice(ctx context.Context, lots []*core.Lot) ([]*OwnedLot, error) {
	result := make([]*OwnedLot, len(lots))
	for i, lot := range lots {
		var err error
		result[i], err = srv.newOwnedLot(ctx, lot)
		if err != nil {
			return nil, errors.Wrap(err, "new owned lot")
		}
	}
	return result, nil
}

func (srv *Service) GetLots(ctx context.Context, user *core.User) ([]*OwnedLot, error) {
	lots, err := srv.Lot.Query().OwnerID(user.ID).All(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get lots")
	}
	return srv.newOwnedLotSlice(ctx, lots)
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

	srv.AdminNotify.Send(&NewLotNotification{
		User: user,
		Lot:  lot,
	})

	return srv.newOwnedLot(ctx, lot)
}
