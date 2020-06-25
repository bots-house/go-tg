package catalog

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

var (
	ErrLotIsNotChannel = core.NewError(
		"lot_is_not_channel",
		"lot is not channel, only channels is supported",
	)
)

type OwnedLot struct {
	*core.Lot

	Topics core.TopicSlice
}

func (srv *Service) newOwnedLot(ctx context.Context, lot *core.Lot) (*OwnedLot, error) {
	topics, err := srv.LotTopic.Get(ctx, lot.ID)
	if err != nil {
		return nil, errors.Wrap(err, "get topics")
	}

	return &OwnedLot{
		Lot:    lot,
		Topics: topics,
	}, nil
}

func (srv *Service) AddLot(ctx context.Context, user *core.User, in *LotInput) (*OwnedLot, error) {
	var lot *core.Lot

	if err := srv.Txier(ctx, func(ctx context.Context) error {
		var err error
		lot, err = srv.addLot(ctx, user, in)
		return err
	}); err != nil {
		return nil, err
	}

	return srv.newOwnedLot(ctx, lot)
}

func (srv *Service) addLot(ctx context.Context, user *core.User, in *LotInput) (*core.Lot, error) {
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

	if err := srv.Lot.Add(ctx, lot); err != nil {
		return nil, errors.Wrap(err, "add lot to store")
	}

	if err := srv.LotTopic.Set(ctx, lot.ID, in.TopicIDs); err != nil {
		return nil, errors.Wrap(err, "set lot topic")
	}

	return lot, nil
}
