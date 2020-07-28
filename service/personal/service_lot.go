package personal

import (
	"context"
	"io"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	"github.com/volatiletech/null/v8"

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
	Files         []core.LotFileID
}

type LotsQuery struct {
	SortBy     core.LotField
	SortByType store.SortType
}

type LotList struct {
	Total int
	Items core.LotSlice
}

type OwnedLot struct {
	*core.Lot
	Files []*OwnedLotUploadedFile
}

type OwnedLotUploadedFile struct {
	Path string
	Name string
	Size int
}

func newOwnedLotUploadedFile(lf *core.LotFile) *OwnedLotUploadedFile {
	return &OwnedLotUploadedFile{
		Path: lf.Path,
		Name: lf.Name,
		Size: lf.Size,
	}
}

func NewOwnedLotUploadedFileSlice(files core.LotFileSlice) []*OwnedLotUploadedFile {
	result := make([]*OwnedLotUploadedFile, len(files))
	for i, file := range files {
		result[i] = newOwnedLotUploadedFile(file)
	}
	return result
}

var (
	ErrLotIsNotChannel = core.NewError(
		"lot_is_not_channel",
		"lot is not channel, only channels is supported",
	)
	ErrLotFileSizeIsLarge = core.NewError(
		"lot_file_size_is_large",
		"lot file size is large (6MB max)",
	)

	ErrLotFileExtensionIsWrong = core.NewError(
		"lot_file_extension_is_wrong",
		"lot file extension is wrong (pdf, png, jpeg)",
	)
)

const (
	uploadLotFileMaxSizeInBytes = 6 * 1024 * 1024
	lotDir                      = "lot"
)

func (srv *Service) newOwnedLot(lot *core.Lot, files []*OwnedLotUploadedFile) (*OwnedLot, error) {
	olot := &OwnedLot{
		Lot:   lot,
		Files: files,
	}

	return olot, nil
}

func (srv *Service) newOwnedLotSlice(lots []*core.Lot) ([]*OwnedLot, error) {
	result := make([]*OwnedLot, len(lots))
	for i, lot := range lots {
		var err error
		result[i], err = srv.newOwnedLot(lot, nil)
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
	return srv.newOwnedLotSlice(lots)
}

func (srv *Service) newLotExtraResourceSlice(ctx context.Context, urls []string) ([]*core.LotExtraResource, error) {
	result := make([]*core.LotExtraResource, len(urls))
	var err error

	for i, url := range urls {
		result[i], err = srv.Parser.Parse(ctx, url)
		if err != nil {
			return nil, err
		}
	}
	return result, nil
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
		avatar, err := srv.Storage.AddByURL(ctx, lotDir, info.Avatar)
		if err != nil {
			return nil, errors.Wrap(err, "add by url")
		}
		lot.Avatar = null.StringFrom(avatar)
	}

	resources, err := srv.newLotExtraResourceSlice(ctx, in.Extra)
	if err != nil {
		return nil, errors.Wrap(err, "format extra resources")
	}
	lot.ExtraResources = resources

	lot.TopicIDs = in.TopicIDs

	var files core.LotFileSlice

	if err := srv.Txier(ctx, func(ctx context.Context) error {
		if err := srv.Lot.Add(ctx, lot); err != nil {
			return errors.Wrap(err, "add lot to store")
		}
		if len(in.Files) != 0 {
			files, err = srv.LotFile.Query().ID(in.Files...).All(ctx)
			if err != nil {
				return errors.Wrap(err, "find lot files")
			}

			for _, file := range files {
				file.LotID = lot.ID
				if err := srv.LotFile.Update(ctx, file); err != nil {
					return errors.Wrap(err, "update file")
				}
			}
		}
		return nil
	}); err != nil {
		return nil, err
	}

	srv.AdminNotify.Send(&NewLotNotification{
		User: user,
		Lot:  lot,
	})

	olufs := NewOwnedLotUploadedFileSlice(files)

	return srv.newOwnedLot(lot, olufs)
}

var ErrLotCantBeCanceled = core.NewError("lot_cant_be_canceled", "lot can't be canceled on current status")

func (srv *Service) CancelLot(
	ctx context.Context,
	user *core.User,
	lotID core.LotID,
	reasonID core.LotCanceledReasonID,
) error {
	lot, err := srv.getOwnedLot(ctx, user, lotID)
	if err != nil {
		return errors.Wrap(err, "get owned lot")
	}

	if !lot.CanCancel() {
		return ErrLotCantBeCanceled
	}

	reason, err := srv.LotCanceledReason.Query().ID(reasonID).One(ctx)
	if err != nil {
		return errors.Wrap(err, "get lot canceled reason")
	}

	lot.CanceledReasonID = reason.ID
	lot.Status = core.LotStatusCanceled

	if err := srv.Lot.Update(ctx, lot.Lot); err != nil {
		return errors.Wrap(err, "update lot")
	}

	srv.AdminNotify.Send(&CanceledLotNotification{
		Lot:    lot.Lot,
		Reason: reason,
	})

	return nil
}

type LotUploadedFile struct {
	ID   core.LotFileID
	Name string
	Path string
	Size int
}

func (srv *Service) UploadLotFile(
	ctx context.Context,
	body io.Reader,
	filename string,
	size int64,
	mimeType string,
) (*LotUploadedFile, error) {
	if size > uploadLotFileMaxSizeInBytes {
		return nil, ErrLotFileSizeIsLarge
	}

	ext := filepath.Ext(filename)
	ext = strings.TrimPrefix(ext, ".")

	if ext != "png" && ext != "pdf" && ext != "jpeg" {
		return nil, ErrLotFileExtensionIsWrong
	}

	path, err := srv.Storage.Add(ctx, lotDir, body, ext)
	if err != nil {
		return nil, errors.Wrap(err, "upload lot file")
	}

	lotFile := core.NewLotFile(filename, int(size), mimeType, path)

	if err := srv.LotFile.Add(ctx, lotFile); err != nil {
		return nil, errors.Wrap(err, "add lot file")
	}

	return &LotUploadedFile{
		ID:   lotFile.ID,
		Path: path,
		Name: filename,
		Size: int(size),
	}, nil
}

func (srv *Service) GetFavoriteLots(
	ctx context.Context,
	user *core.User,
	query *LotsQuery,
	limit int,
	offset int,
) (*LotList, error) {
	favorites, err := srv.LotFavorite.Query().UserID(user.ID).SortBy(core.FavoriteFieldCreatedAt, store.SortTypeDesc).All(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get favorites")
	}

	lotIDs := make([]core.LotID, len(favorites))
	for i, f := range favorites {
		lotIDs[i] = f.LotID
	}

	var lots core.LotSlice
	var finalLots core.LotSlice

	if query.SortBy != 0 {
		lots, err = srv.Lot.Query().SortBy(query.SortBy, query.SortByType).ID(lotIDs...).Limit(limit).Offset(offset).All(ctx)
		if err != nil {
			return nil, errors.Wrap(err, "get lots")
		}
		finalLots = lots
	} else {
		lots, err = srv.Lot.Query().ID(lotIDs...).All(ctx)
		if err != nil {
			return nil, errors.Wrap(err, "get lots")
		}
		for _, f := range favorites {
			finalLots = append(finalLots, lots.Find(f.LotID))
		}
	}

	return &LotList{
		Total: len(favorites),
		Items: finalLots,
	}, nil
}
