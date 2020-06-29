package models

import (
	"github.com/Rhymond/go-money"
	"github.com/go-openapi/swag"

	"github.com/bots-house/birzzha/api/gen/models"
	"github.com/bots-house/birzzha/pkg/storage"
	"github.com/bots-house/birzzha/pkg/tg"
	"github.com/bots-house/birzzha/service/payment"
	"github.com/bots-house/birzzha/service/personal"
)

func newMoney(m *money.Money) *models.Money {
	return &models.Money{
		Currency: swag.String(m.Currency().Code),
		Amount:   swag.Float64(m.AsMajorUnits()),
	}
}

func NewApplicationInvoice(s storage.Storage, inv *personal.ApplicationInvoice) *models.ApplicationInvoice {
	return &models.ApplicationInvoice{
		Lot:   NewOwnedLot(s, inv.Lot),
		Price: newMoney(inv.Price),
		Cashier: &models.ApplicationInvoiceCashier{
			Username: swag.String(inv.CashierUsername),
			Link:     swag.String(tg.GetLinkByUsername(inv.CashierUsername)),
		},
		Gateways: inv.Gateways,
	}
}

func NewPaymentForm(form *payment.Form) *models.PaymentForm {
	vs := make([]*models.PaymentFormValuesItems0, 0, len(form.Values))

	for k := range form.Values {
		vs = append(vs, &models.PaymentFormValuesItems0{
			Key:   swag.String(k),
			Value: swag.String(form.Values.Get(k)),
		})
	}

	return &models.PaymentForm{
		Method: swag.String(form.Method),
		Action: swag.String(form.Action),
		Values: vs,
	}
}

func NewPaymentStatus(status *personal.PaymentStatus) *models.PaymentStatus {
	return &models.PaymentStatus{
		Status:  swag.String(status.Status.String()),
		LotID:   swag.Int64(int64(status.LotID)),
		Purpose: swag.String(status.Purpose.String()),
	}
}
