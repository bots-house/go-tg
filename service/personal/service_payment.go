package personal

import (
	"context"

	"github.com/Rhymond/go-money"
	"github.com/pkg/errors"

	"github.com/bots-house/birzzha/core"
	"github.com/bots-house/birzzha/service/payment"
)

type ApplicationInvoice struct {
	Lot             *OwnedLot
	Price           *money.Money
	CashierUsername string
	Gateways        []string
}

func (srv *Service) getOwnedLot(ctx context.Context, user *core.User, id core.LotID) (*OwnedLot, error) {
	lot, err := srv.Lot.Query().OwnerID(user.ID).ID(id).One(ctx)
	if err != nil {
		return nil, err
	}

	return &OwnedLot{lot}, nil
}

func (srv *Service) GetApplicationInvoice(
	ctx context.Context,
	user *core.User,
	id core.LotID,
) (*ApplicationInvoice, error) {
	lot, err := srv.getOwnedLot(ctx, user, id)
	if err != nil {
		return nil, errors.Wrap(err, "get owned lot")
	}

	settings, err := srv.Settings.Get(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get settings")
	}

	return &ApplicationInvoice{
		Lot:             lot,
		Price:           settings.Prices.Application,
		CashierUsername: settings.CashierUsername,
		Gateways:        []string{"interkassa", "direct"},
	}, nil
}

func (srv *Service) CreateApplicationPayment(
	ctx context.Context,
	user *core.User,
	id core.LotID,
	gatewayName string,
) (*payment.Form, error) {
	var form *payment.Form

	err := srv.Txier(ctx, func(ctx context.Context) error {
		var err error
		form, err = srv.createApplicationPayment(ctx, user, id, gatewayName)
		if err != nil {
			return err
		}
		return nil
	})

	return form, err
}

func (srv *Service) createApplicationPayment(
	ctx context.Context,
	user *core.User,
	id core.LotID,
	gatewayName string,
) (*payment.Form, error) {
	invoice, err := srv.GetApplicationInvoice(ctx, user, id)
	if err != nil {
		return nil, errors.Wrap(err, "get application invoice")
	}

	gateway := srv.Gateways.Get(gatewayName)
	if gateway == nil {
		return nil, core.NewError("gateway_not_found", "requested gateway not found")
	}

	payment := core.NewPayment(
		core.PaymentPurposeApplication,
		user.ID,
		invoice.Lot.ID,
		gatewayName,
		invoice.Price,
	)

	if err := srv.Payment.Add(ctx, payment); err != nil {
		return nil, errors.Wrap(err, "add payment to store")
	}

	form, err := gateway.NewPayment(ctx, user, payment)
	if err != nil {
		return nil, errors.Wrap(err, "new payment")
	}

	return form, nil
}

type PaymentStatus struct {
	Purpose core.PaymentPurpose
	Status  core.PaymentStatus
	LotID   core.LotID
}

func (srv *Service) GetPaymentStatus(ctx context.Context, user *core.User, id core.PaymentID) (*PaymentStatus, error) {
	pm, err := srv.Payment.Query().PayerID(user.ID).ID(id).One(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "query payment")
	}
	return &PaymentStatus{
		Purpose: pm.Purpose,
		Status:  pm.Status,
		LotID:   pm.LotID,
	}, nil
}

func (srv *Service) ProcessGatewayNotification(ctx context.Context, notify *payment.GatewayNotification) error {
	return srv.Txier(ctx, func(ctx context.Context) error {
		return srv.processGatewayNotification(ctx, notify)
	})
}

func (srv *Service) processGatewayNotification(ctx context.Context, notify *payment.GatewayNotification) error {
	payment, err := srv.Payment.Query().ID(notify.PaymentID).One(ctx)
	if err != nil {
		return errors.Wrap(err, "query payment")
	}

	if payment.Status.IsFinal() {
		return core.NewError("payment_status_is_final", "try to modify payment with status final")
	}

	if equal, _ := payment.Requested.Equals(notify.Requested); !equal {
		return core.NewError("payment_requested_is_not_equal", "payment requested amount invalid")
	}

	payment.Status = notify.Status
	payment.ExternalID.SetValid(notify.ExternalID)

	if payment.Status.IsFinal() {
		payment.Paid = notify.Paid
		payment.Received = notify.Received
		payment.Metadata = notify.Metadata
	}

	if err := srv.Payment.Update(ctx, payment); err != nil {
		return errors.Wrap(err, "update payment")
	}

	if payment.Status == core.PaymentStatusSuccess {
		switch payment.Purpose {
		case core.PaymentPurposeApplication:
			return srv.onPaymentApplication(ctx, payment)
		}
	}

	return nil
}

func (srv *Service) onPaymentApplication(ctx context.Context, pm *core.Payment) error {
	lot, err := srv.Lot.Query().ID(pm.LotID).One(ctx)
	if err != nil {
		return errors.Wrap(err, "query payment")
	}

	lot.SetStatus(core.LotStatusPaid)

	if err := srv.Lot.Update(ctx, lot); err != nil {
		return errors.Wrap(err, "update lot")
	}

	srv.AdminNotify.Send(NewPaymentNotification{
		Payment: pm,
		Lot:     lot,
	})

	return nil
}
