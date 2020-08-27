package models

import (
	"github.com/go-openapi/swag"

	"github.com/bots-house/birzzha/api/gen/models"
	"github.com/bots-house/birzzha/core"
	"github.com/bots-house/birzzha/pkg/storage"
	"github.com/bots-house/birzzha/service/admin"
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
		MonthlyIncome:  swag.Int64(int64(in.MonthlyIncome)),
		PricePerMember: swag.Float64(in.PricePerMember),
		PricePerView:   swag.Float64(in.PricePerView),
		PaybackPeriod:  nullFloat64ToFloat64(in.PaybackPeriod),
	}
}

func newLotExtraResource(in *core.LotExtraResource) *models.LotExtraResource {
	return &models.LotExtraResource{
		URL:         swag.String(in.URL),
		Title:       swag.String(in.Title),
		Image:       swag.String(in.Image),
		Description: swag.String(in.Description),
		Domain:      swag.String(in.Domain),
	}
}

func newLotExtraResourceSlice(in []*core.LotExtraResource) []*models.LotExtraResource {
	result := make([]*models.LotExtraResource, len(in))

	for i, v := range in {
		result[i] = newLotExtraResource(v)
	}

	return result
}

func toLotExtraResource(in *models.LotExtraResource) *core.LotExtraResource {
	return &core.LotExtraResource{
		URL:         swag.StringValue(in.URL),
		Title:       swag.StringValue(in.Title),
		Image:       swag.StringValue(in.Image),
		Description: swag.StringValue(in.Description),
		Domain:      swag.StringValue(in.Domain),
	}
}

func ToLotExtraResourceSlice(in []*models.LotExtraResource) []*core.LotExtraResource {
	result := make([]*core.LotExtraResource, len(in))

	for i, v := range in {
		result[i] = toLotExtraResource(v)
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
		ID:            swag.Int64(int64(in.ID)),
		ExternalID:    swag.Int64(in.ExternalID),
		Name:          swag.String(in.Name),
		Username:      nullStringToString(in.Username),
		Link:          swag.String(in.Link()),
		Bio:           nullStringToString(in.Bio),
		Price:         newLotPrice(in.Price),
		DeclineReason: in.DeclineReason.Ptr(),
		Status:        swag.String(in.Status.String()),
		Comment:       swag.String(in.Comment),
		Metrics:       newLotMetrics(in.Metrics),
		Topics:        NewTopicIDSlice(in.TopicIDs),
		CreatedAt:     timeToUnix(in.CreatedAt),
		PaidAt:        nullTimeToUnix(in.PaidAt),
		ApprovedAt:    nullTimeToUnix(in.ApprovedAt),
		PublishedAt:   nullTimeToUnix(in.PublishedAt),
		ScheduledAt:   nullTimeToUnix(in.ScheduledAt),
		Extra:         newLotExtraResourceSlice(in.ExtraResources),
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

func NewLotStatusesCount(in *admin.LotStatusesCount) *models.LotStatusesCount {
	return &models.LotStatusesCount{
		Created:   swag.Int64(int64(in.Created)),
		Paid:      swag.Int64(int64(in.Paid)),
		Published: swag.Int64(int64(in.Published)),
		Declined:  swag.Int64(int64(in.Declined)),
		Canceled:  swag.Int64(int64(in.Canceled)),
		Scheduled: swag.Int64(int64(in.Scheduled)),
	}
}

func newAdminLotUploadedFile(s storage.Storage, in *core.LotFile) *models.AdminLotUploadedFile {
	return &models.AdminLotUploadedFile{
		URL:  swag.String(s.PublicURL(in.Path)),
		Name: swag.String(in.Name),
		Size: swag.Int64(int64(in.Size)),
	}
}

func newAdminLotUploadedFileSlice(s storage.Storage, in core.LotFileSlice) []*models.AdminLotUploadedFile {
	result := make([]*models.AdminLotUploadedFile, len(in))
	for i, v := range in {
		result[i] = newAdminLotUploadedFile(s, v)
	}
	return result
}

func newAdminLot(s storage.Storage, in *admin.LotItem) *models.AdminLot {
	adminLot := &models.AdminLot{
		ID:             swag.Int64(int64(in.ID)),
		ExternalID:     swag.Int64(int64(in.ID)),
		Name:           swag.String(in.Name),
		Status:         swag.String(in.Status.String()),
		Price:          newLotPrice(in.Price),
		Topics:         NewTopicIDSlice(in.TopicIDs),
		CreatedAt:      timeToUnix(in.CreatedAt),
		PaidAt:         nullTimeToUnix(in.PaidAt),
		ApprovedAt:     nullTimeToUnix(in.ApprovedAt),
		ScheduledAt:    nullTimeToUnix(in.ScheduledAt),
		PublishedAt:    nullTimeToUnix(in.PublishedAt),
		Files:          newAdminLotUploadedFileSlice(s, in.Files),
		User:           NewUser(s, in.Owner),
		Username:       nullStringToString(in.Username),
		JoinLink:       swag.String(in.Link()),
		DeclinedReason: nullStringToString(in.DeclineReason),
	}
	if in.Avatar.Valid {
		adminLot.Avatar = swag.String(s.PublicURL(in.Avatar.String))
	}

	if in.CanceledReason != nil {
		adminLot.CanceledReason = swag.String(in.CanceledReason.Why)
	}

	return adminLot
}

func newAdminLotSlice(s storage.Storage, in []*admin.LotItem) []*models.AdminLot {
	out := make([]*models.AdminLot, len(in))
	for i, v := range in {
		out[i] = newAdminLot(s, v)
	}
	return out
}

func NewAdminLotItemList(s storage.Storage, in *admin.LotItemList) *models.AdminLotItemList {
	return &models.AdminLotItemList{
		Total: swag.Int64(int64(in.Total)),
		Items: newAdminLotSlice(s, in.Items),
	}
}

func NewPersonalItemLot(s storage.Storage, in *core.Lot) *models.LotListItem {
	lotListItem := &models.LotListItem{
		ID:          swag.Int64(int64(in.ID)),
		Name:        swag.String(in.Name),
		Username:    nullStringToString(in.Username),
		Link:        swag.String(in.Link()),
		Price:       newLotPrice(in.Price),
		Comment:     swag.String(in.Comment),
		Metrics:     newLotMetrics(in.Metrics),
		InFavorites: swag.Bool((true)),
		Topics:      NewTopicIDSlice(in.TopicIDs),
		CreatedAt:   timeToUnix(in.CreatedAt),
	}
	if in.Avatar.Valid {
		lotListItem.Avatar = swag.String(s.PublicURL(in.Avatar.String))
	}

	return lotListItem
}

func NewPersonalLotList(s storage.Storage, list *personal.LotList) *models.LotList {
	return &models.LotList{
		Items: NewPersonalItemLotSlice(s, list.Items),
		Total: swag.Int64(int64(list.Total)),
	}
}

func NewPersonalItemLotSlice(s storage.Storage, in core.LotSlice) []*models.LotListItem {
	result := make([]*models.LotListItem, len(in))

	for i, v := range in {
		result[i] = NewPersonalItemLot(s, v)
	}

	return result
}

func NewPostText(in string) *models.AdminPostText {
	return &models.AdminPostText{
		Text: swag.String(in),
	}
}

func ToTopicIDs(in []int64) []core.TopicID {
	out := make([]core.TopicID, len(in))
	for i, id := range in {
		out[i] = core.TopicID(id)
	}
	return out
}
func NewAdminLotUploadedFile(s storage.Storage, in *admin.LotUploadedFile) *models.LotUploadedFile {
	return &models.LotUploadedFile{
		ID:   swag.Int64(int64(in.ID)),
		URL:  swag.String(s.PublicURL(in.Path)),
		Name: swag.String(in.Name),
		Size: swag.Int64(int64(in.Size)),
	}
}

func NewAdminLotUploadedFileSlice(s storage.Storage, in []*admin.LotUploadedFile) []*models.LotUploadedFile {
	out := make([]*models.LotUploadedFile, len(in))
	for i, v := range in {
		out[i] = NewAdminLotUploadedFile(s, v)
	}
	return out
}

func NewAdminFullLot(s storage.Storage, in *admin.FullLot) *models.AdminFullLot {
	lot := &models.AdminFullLot{
		ID:             swag.Int64(int64(in.ID)),
		Name:           swag.String(in.Name),
		Username:       nullStringToString(in.Username),
		Link:           swag.String(in.Link()),
		Price:          newLotPrice(in.Price),
		Comment:        swag.String(in.Comment),
		Metrics:        newLotMetrics(in.Metrics),
		Topics:         NewTopicIDSlice(in.TopicIDs),
		CreatedAt:      timeToUnix(in.CreatedAt),
		Bio:            nullStringToString(in.Bio),
		User:           newLotOwner(s, in.User),
		TgstatLink:     swag.String(in.TgstatLink()),
		TelemetrLink:   swag.String(in.TelemetrLink()),
		Views:          swag.Int64(int64(in.Views)),
		Extra:          newLotExtraResourceSlice(in.ExtraResources),
		Files:          NewAdminLotUploadedFileSlice(s, in.Files),
		PaidAt:         nullTimeToUnix(in.PaidAt),
		ApprovedAt:     nullTimeToUnix(in.ApprovedAt),
		PublishedAt:    nullTimeToUnix(in.PublishedAt),
		ScheduledAt:    nullTimeToUnix(in.ScheduledAt),
		DeclinedReason: nullStringToString(in.DeclineReason),
	}

	if in.Avatar.Valid {
		lot.Avatar = swag.String(s.PublicURL(in.Avatar.String))
	}

	if in.CanceledReason != nil {
		lot.CanceledReason = swag.String(in.CanceledReason.Why)
	}
	return lot
}
