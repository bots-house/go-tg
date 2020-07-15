package models

import (
	"github.com/go-openapi/swag"

	"github.com/bots-house/birzzha/api/gen/models"
	"github.com/bots-house/birzzha/core"
	"github.com/bots-house/birzzha/pkg/storage"
	"github.com/bots-house/birzzha/service/catalog"
	"github.com/bots-house/birzzha/service/personal"
)

func newLotPrice(in core.LotPrice) *models.LotPrice {
	return &models.LotPrice{
		Current:   swag.Int64(int64(in.Current)),
		Previous:  nullIntToInt64(in.Previous),
		IsBargain: swag.Bool(in.IsBargain),
	}
}

func newLotMetrics(in core.LotMetrics) *models.LotMetrics {
	return &models.LotMetrics{
		MembersCount:   swag.Int64(int64(in.MembersCount)),
		DailyCoverage:  swag.Int64(int64(in.DailyCoverage)),
		MonthlyIncome:  nullIntToInt64(in.MonthlyIncome),
		PricePerMember: swag.Float64(in.PricePerMember),
		PricePerView:   swag.Float64(in.PricePerView),
		PaybackPeriod:  nullFloat64ToFloat64(in.PaybackPeriod),
	}
}

func newLotExtraResourceSlice(in []core.LotExtraResource) []*models.LotExtraResource {
	result := make([]*models.LotExtraResource, len(in))

	for i, v := range in {
		result[i] = &models.LotExtraResource{URL: v.URL}
	}

	return result
}

func NewOwnedLotSlice(s storage.Storage, lots []*personal.OwnedLot) []*models.OwnedLot {
	items := make([]*models.OwnedLot, len(lots))

	for i, lot := range lots {
		items[i] = NewOwnedLot(s, lot)
	}

	return items
}

func NewOwnedLot(s storage.Storage, in *personal.OwnedLot) *models.OwnedLot {
	ownedLot := &models.OwnedLot{
		ID:          swag.Int64(int64(in.ID)),
		ExternalID:  swag.Int64(in.ExternalID),
		Name:        swag.String(in.Name),
		Username:    nullStringToString(in.Username),
		Link:        swag.String(in.Link()),
		Bio:         nullStringToString(in.Bio),
		Price:       newLotPrice(in.Price),
		Status:      swag.String(in.Status.String()),
		Comment:     swag.String(in.Comment),
		Metrics:     newLotMetrics(in.Metrics),
		Topics:      NewTopicIDSlice(in.TopicIDs),
		CreatedAt:   timeToUnix(in.CreatedAt),
		PaidAt:      nullTimeToUnix(in.PaidAt),
		ApprovedAt:  nullTimeToUnix(in.ApprovedAt),
		PublishedAt: nullTimeToUnix(in.PublishedAt),
		Extra:       newLotExtraResourceSlice(in.ExtraResources),
		Actions: &models.OwnedLotActions{
			CanChangePricePaid: swag.Bool(in.CanChangePricePaid()),
			CanChangePriceFree: swag.Bool(in.CanChangePriceFree()),
			CanCancel:          swag.Bool(in.CanCancel()),
		},
		Files: NewOwnedLotUploadedFileSlice(s, in.Files),
	}

	if in.CanceledReasonID != 0 {
		ownedLot.CanceledReasonID = swag.Int64(int64(in.CanceledReasonID))
	}

	if in.Avatar.Valid {
		ownedLot.Avatar = swag.String(s.PublicURL(in.Avatar.String))
	}

	return ownedLot
}

func NewItemLot(s storage.Storage, in *catalog.ItemLot) *models.LotListItem {
	lotListItem := &models.LotListItem{
		ID:          swag.Int64(int64(in.ID)),
		Name:        swag.String(in.Name),
		Username:    nullStringToString(in.Username),
		Link:        swag.String(in.Link()),
		Price:       newLotPrice(in.Price),
		Comment:     swag.String(in.Comment),
		Metrics:     newLotMetrics(in.Metrics),
		InFavorites: nullBoolToBool(in.InFavorites),
		Topics:      NewTopicIDSlice(in.TopicIDs),
		CreatedAt:   timeToUnix(in.CreatedAt),
	}
	if in.Avatar.Valid {
		lotListItem.Avatar = swag.String(s.PublicURL(in.Avatar.String))
	}

	return lotListItem
}

func NewLotList(s storage.Storage, list *catalog.LotList) *models.LotList {
	return &models.LotList{
		Items: NewItemLotSlice(s, list.Items),
		Total: swag.Int64(int64(list.Total)),
	}
}

func NewItemLotSlice(s storage.Storage, in []*catalog.ItemLot) []*models.LotListItem {
	result := make([]*models.LotListItem, len(in))

	for i, v := range in {
		result[i] = NewItemLot(s, v)
	}

	return result
}

func newLotOwner(s storage.Storage, user *core.User) *models.LotOwner {
	usr := &models.LotOwner{
		FirstName: swag.String(user.FirstName),
		LastName:  nullStringToString(user.LastName),
		Username:  nullStringToString(user.Telegram.Username),
		Link:      nullStringToString(user.Telegram.TelegramLink()),
	}
	if user.Avatar.Valid {
		usr.Avatar = swag.String(s.PublicURL(user.Avatar.String))
	}

	return usr
}

func NewFullLot(s storage.Storage, in *catalog.FullLot) *models.FullLot {
	lot := &models.FullLot{
		ID:           swag.Int64(int64(in.ID)),
		Name:         swag.String(in.Name),
		Username:     nullStringToString(in.Username),
		Link:         swag.String(in.Link()),
		Price:        newLotPrice(in.Price),
		Comment:      swag.String(in.Comment),
		Metrics:      newLotMetrics(in.Metrics),
		InFavorites:  nullBoolToBool(in.InFavorites),
		Topics:       NewTopicIDSlice(in.TopicIDs),
		CreatedAt:    timeToUnix(in.CreatedAt),
		Bio:          nullStringToString(in.Bio),
		User:         newLotOwner(s, in.User),
		TgstatLink:   swag.String(in.TgstatLink()),
		TelemetrLink: swag.String(in.TelemetrLink()),
		Views:        swag.Int64(int64(in.Views)),
		Extra:        newLotExtraResourceSlice(in.ExtraResources),
		Files:        NewOwnedLotUploadedFileSlice(s, in.Files),
	}

	if in.Avatar.Valid {
		lot.Avatar = swag.String(s.PublicURL(in.Avatar.String))
	}
	return lot
}

func NewUploadedLotFile(s storage.Storage, in *personal.LotUploadedFile) *models.LotUploadedFile {
	return &models.LotUploadedFile{
		ID:   swag.Int64(int64(in.ID)),
		URL:  swag.String(s.PublicURL(in.Path)),
		Name: swag.String(in.Name),
		Size: swag.Int64(int64(in.Size)),
	}
}

func newOwnedLotUploadedFile(s storage.Storage, in *personal.OwnedLotUploadedFile) *models.OwnedLotUploadedFile {
	return &models.OwnedLotUploadedFile{
		URL:  swag.String(s.PublicURL(in.Path)),
		Name: swag.String(in.Name),
		Size: swag.Int64(int64(in.Size)),
	}
}

func NewOwnedLotUploadedFileSlice(s storage.Storage, in []*personal.OwnedLotUploadedFile) []*models.OwnedLotUploadedFile {
	result := make([]*models.OwnedLotUploadedFile, len(in))
	for i, v := range in {
		result[i] = newOwnedLotUploadedFile(s, v)
	}
	return result
}
