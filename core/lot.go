package core

import (
	"context"
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"

	"github.com/volatiletech/null/v8"

	"github.com/bots-house/birzzha/store"
)

// LotID it's unique lot id.
type LotID int

// Extra resources like chats and bots.
type LotExtraResource struct {
	URL string
}

// LotPrice info
type LotPrice struct {
	// Current price of lot
	Current int

	// Previous price of lot
	Previous null.Int

	// True, if bargaining possible
	IsBargain bool
}

func NewLotPrice(price int, isBargain bool) LotPrice {
	return LotPrice{
		Current:   price,
		IsBargain: isBargain,
	}
}

type LotMetrics struct {
	// Members count in channel
	MembersCount int

	// DailyCoverage
	DailyCoverage int

	// Total monthly income
	MonthlyIncome null.Int

	// Price of one member of channel
	// (MembersCount / Price)
	// Computed.
	PricePerMember float64

	// Price of ne view in channel
	// (Price / DailyAverage)
	// Computed.
	PricePerView float64

	// Payback period in months
	// Computed.
	PaybackPeriod null.Float64
}

func NewLotMetrics(price int, membersCount, dailyCoverage int, monthlyIncome null.Int) LotMetrics {
	metrics := LotMetrics{
		MembersCount:  membersCount,
		DailyCoverage: dailyCoverage,
		MonthlyIncome: monthlyIncome,
	}

	metrics.Refresh(price)

	return metrics
}

func (lm *LotMetrics) Refresh(price int) {
	lm.refreshPricePerView(price)
	lm.refreshPricePerMember(price)
	lm.refreshPaybackPeriod(price)
}

func (lm *LotMetrics) refreshPricePerMember(price int) {
	v := float64(price) / float64(lm.MembersCount)
	lm.PricePerMember = math.Round(v*100) / 100
}

func (lm *LotMetrics) refreshPricePerView(price int) {
	v := float64(price) / float64(lm.DailyCoverage)
	lm.PricePerView = math.Round(v*100) / 100
}

func (lm *LotMetrics) refreshPaybackPeriod(price int) {
	if lm.MonthlyIncome.Valid {
		v := float64(price) / float64(lm.MonthlyIncome.Int)
		lm.PaybackPeriod = null.Float64From(math.Round(v*100) / 100)
	} else {
		lm.PaybackPeriod.Valid = false
	}
}

type LotStatus int8

const (
	LotStatusCreated LotStatus = iota + 1
	LotStatusPaid
	LotStatusPublished
	LotStatusDeclined
	LotStatusCanceled
)

var ShowLotStatus = []LotStatus{
	LotStatusPaid,
	LotStatusCreated,
}

var (
	lotStatusToString = map[LotStatus]string{
		LotStatusCreated:   "created",
		LotStatusPaid:      "paid",
		LotStatusPublished: "published",
		LotStatusDeclined:  "declined",
		LotStatusCanceled:  "canceled",
	}

	stringToLotStatus = func() map[string]LotStatus {
		result := make(map[string]LotStatus, len(lotStatusToString))

		for k, v := range lotStatusToString {
			result[v] = k
		}

		return result
	}()

	ErrLotStatusInvalid = NewError("lot_status_invalid", "lot status is invalid")
)

// ParseLotStatus returns lot status by string
func ParseLotStatus(v string) (LotStatus, error) {
	status, ok := stringToLotStatus[strings.ToLower(v)]
	if !ok {
		return LotStatus(-1), ErrLotStatusInvalid
	}
	return status, nil
}

// String representation of lot status.
func (ls LotStatus) String() string {
	return lotStatusToString[ls]
}

type LotViews struct {
	// Views count on Telegram.
	Telegram int

	// Views count on site.
	Site int
}

func (lv LotViews) Total() int {
	// return lv.Site + lv.Telegram
	return rand.Intn(500000)
}

// Lot for sale
type Lot struct {
	// Unique ID of lot.
	ID LotID

	// Reference to owner of lot
	OwnerID UserID

	// ID of lot in external system.
	ExternalID int64

	// Name of channel
	Name string

	// Avatar of lot
	Avatar null.String

	// Views
	Views LotViews

	// Username of channel
	Username null.String

	// Private join link of channel
	JoinLink null.String

	// Status of lot.
	// Default is created.
	Status LotStatus

	// Reference to canceled reason.
	// Optional.
	CanceledReasonID LotCanceledReasonID

	// Bio of channel
	Bio null.String

	// Price of lot
	Price LotPrice

	// Comment from owner
	Comment string

	// IDs of topics
	TopicIDs []TopicID

	// Metrics of lot
	Metrics LotMetrics

	// Extra resources of lot.
	ExtraResources []LotExtraResource

	// Time when lot was created
	CreatedAt time.Time

	// Time when lot was paid
	PaidAt null.Time

	// Time when lot was approved
	ApprovedAt null.Time

	// Admin who approve the lot
	ApprovedBy UserID

	// Time when lot was published in Telegram
	PublishedAt null.Time
}

func (lot *Lot) CanCancel() bool {
	return lot.Status == LotStatusPublished || lot.Status == LotStatusPaid
}

func (lot *Lot) CanChangePriceFree() bool {
	return lot.Status == LotStatusCreated || lot.Status == LotStatusPaid
}

func (lot *Lot) CanChangePricePaid() bool {
	return lot.Status == LotStatusPublished
}

func (lot *Lot) SetStatus(status LotStatus) {
	switch status {
	case LotStatusPaid:
		lot.Status = status
		lot.PaidAt.SetValid(time.Now())
	case LotStatusPublished:
		lot.Status = status
		lot.PublishedAt.SetValid(time.Now())
	}
}

func (lot *Lot) Link() string {
	if lot.JoinLink.Valid {
		return lot.JoinLink.String
	} else {
		return fmt.Sprintf("https://t.me/%s", lot.Username.String)
	}
}

type LotSlice []*Lot

func (lots LotSlice) SortByID(ids []LotID) LotSlice {
	result := make(LotSlice, 0, len(lots))
	lm := make(map[LotID]*Lot)

	for _, id := range ids {
		for _, lot := range lots {
			_, ok := lm[lot.ID]
			if id == lot.ID && !ok {
				result = append(result, lot)
				lm[id] = lot
			}
		}
	}
	return result
}

func (lots LotSlice) IDs() []LotID {
	ids := make([]LotID, len(lots))
	for i, lot := range lots {
		ids[i] = lot.ID
	}
	return ids
}

func NewLot(
	ownerID UserID,
	externalID int64,
	name string,
	price LotPrice,
	comment string,
	membersCount int,
	dailyCoverage int,
	monthlyIncome null.Int,
) *Lot {
	return &Lot{
		OwnerID:        ownerID,
		ExternalID:     externalID,
		Name:           name,
		Price:          price,
		Status:         LotStatusCreated,
		Comment:        comment,
		Metrics:        NewLotMetrics(price.Current, membersCount, dailyCoverage, monthlyIncome),
		ExtraResources: nil,
		CreatedAt:      time.Now(),
	}
}

var ErrLotNotFound = NewError("lot_not_found", "lot not found")

type LotFilterBoundaries struct {
	PriceMin int
	PriceMax int

	MembersCountMin int
	MembersCountMax int

	DailyCoverageMin int
	DailyCoverageMax int

	MonthlyIncomeMin int
	MonthlyIncomeMax int

	PricePerMemberMin float64
	PricePerMemberMax float64

	PricePerViewMin float64
	PricePerViewMax float64

	PaybackPeriodMin float64
	PaybackPeriodMax float64
}

type LotFilterBoundariesQuery struct {
	Topics []TopicID
}

type LotStore interface {
	// Add lot to store
	Add(ctx context.Context, lot *Lot) error

	// Update lot in store
	Update(ctx context.Context, lot *Lot) error

	// Get filters min and max values depend of topic.
	FilterBoundaries(ctx context.Context, query *LotFilterBoundariesQuery) (*LotFilterBoundaries, error)

	// Find similar lot id's.
	SimilarLotIDs(ctx context.Context, id LotID, limit int, offset int) ([]LotID, error)

	// Find similar lots count.
	SimilarLotsCount(ctx context.Context, id LotID) (int, error)

	// Complex query for lots
	Query() LotStoreQuery
}

type LotField int8

const (
	LotFieldMembersCount LotField = iota + 1
	LotFieldPrice
	LotFieldPricePerMember
	LotFieldDailyCoverage
	LotFieldPricePerView
	LotFieldMonthlyIncome
	LotFieldPaybackPeriod
	LotFieldCreatedAt
)

var (
	stringToLotField = map[string]LotField{
		"members_count":    LotFieldMembersCount,
		"price":            LotFieldPrice,
		"price_per_member": LotFieldPricePerMember,
		"daily_coverage":   LotFieldDailyCoverage,
		"price_per_view":   LotFieldPricePerView,
		"monthly_income":   LotFieldMonthlyIncome,
		"payback_period":   LotFieldPaybackPeriod,
		"created_at":       LotFieldCreatedAt,
	}

	lotFieldToString = mirrorStringToLotField(stringToLotField)
)

func mirrorStringToLotField(in map[string]LotField) map[LotField]string {
	result := make(map[LotField]string, len(in))

	for k, v := range in {
		result[v] = k
	}

	return result
}

var ErrInvalidLotField = NewError("invalid_lot_field", "invalid lot field")

func ParseLotField(v string) (LotField, error) {
	f, ok := stringToLotField[v]
	if !ok {
		return LotField(-1), ErrInvalidLotField
	}
	return f, nil
}

func (lf LotField) String() string {
	return lotFieldToString[lf]
}

type LotStoreQuery interface {
	ID(ids ...LotID) LotStoreQuery
	TopicIDs(ids ...TopicID) LotStoreQuery

	MembersCountFrom(v int) LotStoreQuery
	MembersCountTo(v int) LotStoreQuery

	PriceFrom(v int) LotStoreQuery
	PriceTo(v int) LotStoreQuery

	PricePerMemberFrom(v float64) LotStoreQuery
	PricePerMemberTo(v float64) LotStoreQuery

	DailyCoverageFrom(v int) LotStoreQuery
	DailyCoverageTo(v int) LotStoreQuery

	PricePerViewFrom(v float64) LotStoreQuery
	PricePerViewTo(v float64) LotStoreQuery

	MonthlyIncomeFrom(v int) LotStoreQuery
	MonthlyIncomeTo(v int) LotStoreQuery

	PaybackPeriodFrom(v float64) LotStoreQuery
	PaybackPeriodTo(v float64) LotStoreQuery

	SortBy(field LotField, typ store.SortType) LotStoreQuery
	Offset(v int) LotStoreQuery
	Limit(v int) LotStoreQuery

	OwnerID(id UserID) LotStoreQuery
	Statuses(statuses ...LotStatus) LotStoreQuery

	Count(ctx context.Context) (int, error)
	One(ctx context.Context) (*Lot, error)
	All(ctx context.Context) (LotSlice, error)
}
