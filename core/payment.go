package core

import (
	"context"
	"time"

	"github.com/Rhymond/go-money"
	"github.com/volatiletech/null/v8"
)

// PaymentID it's unique payment ID.
type PaymentID int

type PaymentPurpose int8

const (
	PaymentPurposeApplication PaymentPurpose = iota + 1
	PaymentPurposeChangePrice
)

var (
	paymentPurposeToString = map[PaymentPurpose]string{
		PaymentPurposeApplication: "application",
		PaymentPurposeChangePrice: "change_price",
	}

	stringToPaymentPurpose = func() map[string]PaymentPurpose {
		result := make(map[string]PaymentPurpose, len(paymentPurposeToString))
		for k, v := range paymentPurposeToString {
			result[v] = k
		}
		return result
	}()
)

var (
	ErrPaymentPurposeInvalid = NewError("payment_purpose_invalid", "payment purpose invalid")
)

func ParsePaymentPurpose(v string) (PaymentPurpose, error) {
	pp, ok := stringToPaymentPurpose[v]
	if !ok {
		return PaymentPurpose(-1), ErrPaymentPurposeInvalid
	}
	return pp, nil
}

func (pp PaymentPurpose) String() string {
	return paymentPurposeToString[pp]
}

type PaymentStatus int8

const (
	PaymentStatusCreated PaymentStatus = iota + 1
	PaymentStatusPending
	PaymentStatusSuccess
	PaymentStatusFailed
	PaymentStatusRefunded
)

func (ps PaymentStatus) IsFinal() bool {
	return ps == PaymentStatusFailed || ps == PaymentStatusSuccess || ps == PaymentStatusRefunded
}

var (
	paymentStatusToString = map[PaymentStatus]string{
		PaymentStatusCreated:  "created",
		PaymentStatusPending:  "pending",
		PaymentStatusSuccess:  "success",
		PaymentStatusFailed:   "failed",
		PaymentStatusRefunded: "refunded",
	}

	stringToPaymentStatus = func() map[string]PaymentStatus {
		result := make(map[string]PaymentStatus, len(paymentStatusToString))
		for k, v := range paymentStatusToString {
			result[v] = k
		}
		return result
	}()
)

var (
	ErrPaymentStatusInvalid = NewError("payment_status_invalid", "payment status invalid")
)

func ParsePaymentStatus(v string) (PaymentStatus, error) {
	ps, ok := stringToPaymentStatus[v]
	if !ok {
		return PaymentStatus(-1), ErrPaymentStatusInvalid
	}
	return ps, nil
}

func (ps PaymentStatus) String() string {
	return paymentStatusToString[ps]
}

type Payment struct {
	// Unique ID of payment
	ID PaymentID

	// ID of payment in gateway
	ExternalID null.String

	// Payment purpose
	Purpose PaymentPurpose

	// Who pays?
	PayerID UserID

	// Reference to lot.
	LotID LotID

	// Gateway of payment
	Gateway string

	// Status of payment
	Status PaymentStatus

	// How much we request?
	Requested *money.Money

	// How much user paid?
	Paid *money.Money

	// How much we receive?
	Received *money.Money

	Metadata map[string]string

	// When payment was created?
	CreatedAt time.Time

	// When payment was updated?
	UpdatedAt null.Time
}

func NewPayment(
	purpose PaymentPurpose,
	payer UserID,
	lot LotID,
	gateway string,
	requested *money.Money,
) *Payment {
	return &Payment{
		Purpose:   purpose,
		PayerID:   payer,
		LotID:     lot,
		Gateway:   gateway,
		Status:    PaymentStatusCreated,
		Requested: requested,
		CreatedAt: time.Now(),
	}
}

var (
	ErrPaymentNotFound = NewError("payment_not_found", "payment not found")
)

type PaymentStore interface {
	Add(ctx context.Context, payment *Payment) error
	Update(ctx context.Context, payment *Payment) error
	Query() PaymentStoreQuery
}

type PaymentStoreQuery interface {
	ID(id PaymentID) PaymentStoreQuery
	PayerID(id UserID) PaymentStoreQuery

	One(ctx context.Context) (*Payment, error)
	All(ctx context.Context) ([]*Payment, error)
}
