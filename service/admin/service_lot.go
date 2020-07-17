package admin

import (
	"context"

	"github.com/bots-house/birzzha/core"
	"github.com/pkg/errors"
)

type LotStatusesCount struct {
	Created   int
	Paid      int
	Published int
	Declined  int
	Canceled  int
}

type LotItem struct {
	*core.Lot
	Owner          *core.User
	CanceledReason *core.LotCanceledReason
	Files          core.LotFileSlice
}

type FullLot struct {
	Total int
	Items []*LotItem
}

func (srv *Service) newLotStatusesCount(statuses core.LotsCountByStatusSlice) *LotStatusesCount {
	lotStatusesCount := &LotStatusesCount{}

	for _, s := range statuses {
		switch s.Status {
		case "created":
			lotStatusesCount.Created = s.Count
		case "paid":
			lotStatusesCount.Paid = s.Count
		case "published":
			lotStatusesCount.Published = s.Count
		case "declined":
			lotStatusesCount.Declined = s.Count
		case "canceled":
			lotStatusesCount.Canceled = s.Count
		}
	}
	return lotStatusesCount
}

func (srv *Service) GetLotStatusesCount(ctx context.Context, user *core.User, id core.UserID) (*LotStatusesCount, error) {
	var err error
	if err = srv.IsAdmin(user); err != nil {
		return nil, err
	}

	var statuses core.LotsCountByStatusSlice
	if id == 0 {
		statuses, err = srv.Lot.LotsCountByStatus(ctx, nil)
		if err != nil {
			return nil, errors.Wrap(err, "get lots statuses")
		}
	} else {
		statuses, err = srv.Lot.LotsCountByStatus(ctx, &core.LotsCountByStatusFilter{UserID: id})
		if err != nil {
			return nil, errors.Wrap(err, "get lots statuses")
		}
	}

	return srv.newLotStatusesCount(statuses), nil
}

func (srv *Service) newLotItem(
	lot *core.Lot,
	owners core.UserSlice,
	files core.LotFileSlice,
	canceledReasons core.LotCanceledReasonSlice,
) *LotItem {

	return &LotItem{
		Lot:            lot,
		Owner:          owners.Find(lot.OwnerID),
		Files:          files.FindByLotID(lot.ID),
		CanceledReason: canceledReasons.Find(lot.CanceledReasonID),
	}
}

func (srv *Service) newLotItemSlice(
	lots core.LotSlice,
	owners core.UserSlice,
	files core.LotFileSlice,
	canceledReasons core.LotCanceledReasonSlice,
) []*LotItem {
	lotItems := make([]*LotItem, len(lots))
	for i, lot := range lots {
		lotItem := srv.newLotItem(lot, owners, files, canceledReasons)
		lotItems[i] = lotItem
	}
	return lotItems
}

func (srv *Service) lotsQuery(id core.UserID, status string) (core.LotStoreQuery, error) {
	query := srv.Lot.Query()

	if id != 0 {
		query = query.OwnerID(id)
	}

	if status != "" {
		parsedStatus, err := core.ParseLotStatus(status)
		if err != nil {
			return nil, errors.Wrap(err, "parse status")
		}
		query = query.Statuses(parsedStatus)
	}
	return query, nil
}

func (srv *Service) getLotsCount(ctx context.Context, id core.UserID, status string) (int, error) {
	query, err := srv.lotsQuery(id, status)
	if err != nil {
		return 0, err
	}
	return query.Count(ctx)
}

func (srv *Service) GetLots(
	ctx context.Context,
	user *core.User,
	id core.UserID,
	status string,
	limit int,
	offset int,
) (*FullLot, error) {
	if err := srv.IsAdmin(user); err != nil {
		return nil, err
	}

	query, err := srv.lotsQuery(id, status)
	if err != nil {
		return nil, err
	}

	lots, err := query.Limit(limit).Offset(offset).All(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get lots")
	}

	ownerIDs := lots.OwnerIDs()
	owners, err := srv.User.Query().ID(ownerIDs...).All(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get users")
	}

	files, err := srv.LotFile.Query().LotID(lots.IDs()...).All(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get files")
	}

	canceledReasonIDs := lots.CanceledReasonIDs()

	canceledReasons, err := srv.LotCanceledReason.Query().ID(canceledReasonIDs...).All(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get canceled reasons")
	}

	lotItems := srv.newLotItemSlice(lots, owners, files, canceledReasons)

	total, err := srv.getLotsCount(ctx, id, status)
	if err != nil {
		return nil, errors.Wrap(err, "get lots count")
	}

	return &FullLot{
		Total: total,
		Items: lotItems,
	}, nil
}
