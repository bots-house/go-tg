// Code generated by SQLBoiler 4.1.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package dal

import (
	"strconv"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/strmangle"
)

// M type is for providing columns and column values to UpdateAll.
type M map[string]interface{}

// ErrSyncFail occurs during insert when the record could not be retrieved in
// order to populate default value information. This usually happens when LastInsertId
// fails or there was a primary key configuration that was not resolvable.
var ErrSyncFail = errors.New("dal: failed to synchronize data after insert")

type insertCache struct {
	query        string
	retQuery     string
	valueMapping []uint64
	retMapping   []uint64
}

type updateCache struct {
	query        string
	valueMapping []uint64
}

func makeCacheKey(cols boil.Columns, nzDefaults []string) string {
	buf := strmangle.GetBuffer()

	buf.WriteString(strconv.Itoa(cols.Kind))
	for _, w := range cols.Cols {
		buf.WriteString(w)
	}

	if len(nzDefaults) != 0 {
		buf.WriteByte('.')
	}
	for _, nz := range nzDefaults {
		buf.WriteString(nz)
	}

	str := buf.String()
	strmangle.PutBuffer(buf)
	return str
}

// Enum values for lot_status
const (
	LotStatusCreated   = "created"
	LotStatusPaid      = "paid"
	LotStatusPublished = "published"
	LotStatusDeclined  = "declined"
	LotStatusCanceled  = "canceled"
	LotStatusScheduled = "scheduled"
)

// Enum values for payment_gateway
const (
	PaymentGatewayInterkassa = "interkassa"
	PaymentGatewayDirect     = "direct"
	PaymentGatewayUnitpay    = "unitpay"
)

// Enum values for payment_purpose
const (
	PaymentPurposeApplication = "application"
	PaymentPurposeChangePrice = "change_price"
)

// Enum values for payment_status
const (
	PaymentStatusCreated = "created"
	PaymentStatusPending = "pending"
	PaymentStatusSuccess = "success"
	PaymentStatusFailed  = "failed"
)

// Enum values for post_status
const (
	PostStatusScheduled = "scheduled"
	PostStatusPublished = "published"
)

// Enum values for joined_from
const (
	JoinedFromSite = "site"
	JoinedFromBot  = "bot"
)
