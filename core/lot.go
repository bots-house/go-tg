package core

import (
	"context"
	"fmt"
	"math"
	"time"

	"github.com/volatiletech/null/v8"
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
	lm.PricePerMember = math.Round(v*100)/100
}

func (lm *LotMetrics) refreshPricePerView(price int) {
	v := float64(price) / float64(lm.DailyCoverage)
	lm.PricePerView = math.Round(v*100)/100
}

func (lm *LotMetrics) refreshPaybackPeriod(price int) {
	if lm.MonthlyIncome.Valid {
		v := float64(price) / float64(lm.MonthlyIncome.Int)
		lm.PaybackPeriod = null.Float64From(math.Round(v*100)/100)
	} else {
		lm.PaybackPeriod.Valid = false
	}
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

	// Username of channel
	Username null.String

	// Private join link of channel
	JoinLink null.String

	// Bio of channel
	Bio null.String

	// Price of lot
	Price LotPrice

	// Comment from owner
	Comment string

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

func (lot *Lot) Link() string {
	if lot.JoinLink.Valid {
		return lot.JoinLink.String
	} else {
		return fmt.Sprintf("https://t.me/%s", lot.Username.String)
	}
}

type LotSlice []*Lot

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
		Comment:        comment,
		Metrics:        NewLotMetrics(price.Current, membersCount, dailyCoverage, monthlyIncome),
		ExtraResources: nil,
		CreatedAt:      time.Now(),
	}
}

var ErrLotNotFound = NewError("lot_not_found", "lot not found")

type LotStore interface {
	// Add lot to store
	Add(ctx context.Context, lot *Lot) error

	// Update lot in store
	Update(ctx context.Context, lot *Lot) error

	// Complex query for lots
	Query() LotStoreQuery
}

type LotStoreQuery interface {
	ID(ids ...LotID) LotStoreQuery
	One(ctx context.Context) (*Lot, error)
	All(ctx context.Context) (LotSlice, error)
}
