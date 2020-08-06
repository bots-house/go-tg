package admin

import (
	"context"

	"fmt"

	"github.com/bots-house/birzzha/core"

	"github.com/bots-house/birzzha/pkg/tg"
	"github.com/bots-house/birzzha/store"
	"github.com/pkg/errors"
	"github.com/volatiletech/null/v8"
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

type LotItemList struct {
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

func (srv *Service) lotsSortQuery(status string, query core.LotStoreQuery) (core.LotStoreQuery, error) {
	if status != "" {
		parsedStatus, err := core.ParseLotStatus(status)
		if err != nil {
			return nil, errors.Wrap(err, "parse status")
		}
		switch parsedStatus {
		case core.LotStatusPaid:
			query = query.SortBy(core.LotFieldPaidAt, store.SortTypeDesc)
		case core.LotStatusPublished:
			query = query.SortBy(core.LotFieldPublishedAt, store.SortTypeDesc)
		default:
			query = query.SortBy(core.LotFieldCreatedAt, store.SortTypeDesc)
		}
	} else {
		query = query.SortBy(core.LotFieldCreatedAt, store.SortTypeDesc)
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
) (*LotItemList, error) {
	if err := srv.IsAdmin(user); err != nil {
		return nil, err
	}

	query, err := srv.lotsQuery(id, status)
	if err != nil {
		return nil, err
	}

	query, err = srv.lotsSortQuery(status, query)
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

	return &LotItemList{
		Total: total,
		Items: lotItems,
	}, nil
}

type InputAdminLot struct {
	Comment string
	Price   int
	Extra   []*core.LotExtraResource
	Topics  []core.TopicID
	Income  int
}

type FullLot struct {
	*core.Lot
	User           *core.User
	Views          int
	CanceledReason *core.LotCanceledReason
	Files          []*LotUploadedFile
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

type LotUploadedFile struct {
	ID   core.LotFileID
	Path string
	Name string
	Size int
}

func newLotUploadedFile(lf *core.LotFile) *LotUploadedFile {
	return &LotUploadedFile{
		Path: lf.Path,
		Name: lf.Name,
		Size: lf.Size,
	}
}

func NewLotUploadedFileSlice(files core.LotFileSlice) []*LotUploadedFile {
	result := make([]*LotUploadedFile, len(files))
	for i, file := range files {
		result[i] = newLotUploadedFile(file)
	}
	return result
}

func (srv *Service) GetPostText(ctx context.Context, user *core.User, id core.LotID) (string, error) {
	if err := srv.IsAdmin(user); err != nil {
		return "", err
	}
	return srv.Posting.GetText(ctx, id)
}

func (srv *Service) SendPostPreview(ctx context.Context, user *core.User, post string) error {
	if err := srv.IsAdmin(user); err != nil {
		return err
	}
	return srv.Posting.SendPreview(ctx, user, post)
}

func (srv *Service) DeclineLot(ctx context.Context, user *core.User, id core.LotID, reason string) error {
	if err := srv.IsAdmin(user); err != nil {
		return err
	}

	lot, err := srv.Lot.Query().ID(id).One(ctx)
	if err != nil {
		return errors.Wrap(err, "get lot")
	}

	lot.DeclineReason = null.StringFrom(reason)
	lot.Status = core.LotStatusDeclined

	if err := srv.Lot.Update(ctx, lot); err != nil {
		return errors.Wrap(err, "update lot")
	}

	return nil
}

func (srv *Service) UpdateLot(ctx context.Context, user *core.User, id core.LotID, in InputAdminLot) (*FullLot, error) {
	if err := srv.IsAdmin(user); err != nil {
		return nil, err
	}

	lot, err := srv.Lot.Query().ID(id).One(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get lot")
	}

	_, err = srv.Topic.Query().ID(in.Topics...).All(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get topics")
	}

	if err := srv.Txier(ctx, func(ctx context.Context) error {
		if err := srv.LotTopic.Set(ctx, id, in.Topics); err != nil {
			return errors.Wrap(err, "set lot topics")
		}

		lot.Comment = in.Comment
		lot.Price.Current = in.Price
		lot.Metrics.MonthlyIncome = null.IntFrom(in.Income)
		lot.ExtraResources = in.Extra

		if err := srv.Lot.Update(ctx, lot); err != nil {
			return errors.Wrap(err, "update lot")
		}
		return nil
	}); err != nil {
		return nil, err
	}

	usr, err := srv.User.Query().ID(lot.OwnerID).One(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get user")
	}

	canceledReason, err := srv.LotCanceledReason.Query().ID(lot.CanceledReasonID).One(ctx)
	if err != core.ErrLotCanceledReasonNotFound && err != nil {
		return nil, errors.Wrap(err, "get canceled reason")
	}

	files, err := srv.LotFile.Query().LotID(lot.ID).All(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get lot files")
	}

	return &FullLot{
		Lot:            lot,
		User:           usr,
		Views:          lot.Views.Total(),
		Files:          NewLotUploadedFileSlice(files),
		CanceledReason: canceledReason,
	}, nil
}

func (srv *Service) GetLot(ctx context.Context, user *core.User, id core.LotID) (*FullLot, error) {
	if err := srv.IsAdmin(user); err != nil {
		return nil, err
	}

	lot, err := srv.Lot.Query().ID(id).One(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get lot")
	}

	usr, err := srv.User.Query().ID(lot.OwnerID).One(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get user")
	}

	canceledReason, err := srv.LotCanceledReason.Query().ID(lot.CanceledReasonID).One(ctx)
	if err != core.ErrLotCanceledReasonNotFound && err != nil {
		return nil, errors.Wrap(err, "get canceled reason")
	}

	files, err := srv.LotFile.Query().LotID(lot.ID).All(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get lot files")
	}

	return &FullLot{
		Lot:            lot,
		User:           usr,
		Views:          lot.Views.Total(),
		Files:          NewLotUploadedFileSlice(files),
		CanceledReason: canceledReason,
	}, nil
}
