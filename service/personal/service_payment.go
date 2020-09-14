package personal

import (
	"context"
	"net/http"
	"net/url"

	"github.com/bots-house/birzzha/pkg/log"

	"github.com/Rhymond/go-money"
	"github.com/pkg/errors"
	"github.com/volatiletech/null/v8"

	"github.com/bots-house/birzzha/core"
	"github.com/bots-house/birzzha/service/payment"
)

const unitPaySuccess = "магазин успешно обработал запрос"

type ApplicationInvoice struct {
	Lot             *OwnedLot
	Price           *money.Money
	CashierUsername string
	Gateways        []string
}

type ChangeInvoice struct {
	Lot             *OwnedLot
	Price           *money.Money
	CashierUsername string
	Gateways        []string
}

var (
	ErrGatewayNotFound = core.NewError("gateway_not_found", "requested gateway not found")
)

func (srv *Service) getOwnedLot(ctx context.Context, user *core.User, id core.LotID) (*OwnedLot, error) {
	lot, err := srv.Lot.Query().OwnerID(user.ID).ID(id).One(ctx)
	if err != nil {
		return nil, err
	}

	files, err := srv.LotFile.Query().LotID(lot.ID).All(ctx)
	if err != nil {
		return nil, err
	}

	olufs := NewOwnedLotUploadedFileSlice(files)

	return &OwnedLot{Lot: lot, Files: olufs}, nil
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
		Gateways:        srv.Gateways.Names(),
	}, nil
}

func (srv *Service) GetChangeInvoice(
	ctx context.Context,
	user *core.User,
	id core.LotID,
) (*ChangeInvoice, error) {
	lot, err := srv.getOwnedLot(ctx, user, id)
	if err != nil {
		return nil, errors.Wrap(err, "get owned lot")
	}

	settings, err := srv.Settings.Get(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get settings")
	}

	return &ChangeInvoice{
		Lot:             lot,
		Price:           settings.Prices.Change,
		CashierUsername: settings.CashierUsername,
		Gateways:        srv.Gateways.Names(),
	}, nil
}

func (srv *Service) CreateApplicationPayment(
	ctx context.Context,
	user *core.User,
	id core.LotID,
	gatewayName string,
	coupon string,
) (*payment.Form, error) {
	var form *payment.Form

	err := srv.Txier(ctx, func(ctx context.Context) error {
		var err error
		form, err = srv.createApplicationPayment(ctx, user, id, gatewayName, coupon)
		if err != nil {
			return err
		}
		return nil
	})

	return form, err
}

func (srv *Service) ApplyCoupon(
	ctx context.Context,
	user *core.User,
	coupon *core.Coupon,
	payment *core.Payment,
) (*core.Payment, error) {
	if err := srv.ValidateCoupon(ctx, user, coupon, payment.Purpose); err != nil {
		return nil, err
	}

	payment.Requested = money.New(int64((payment.Requested.AsMajorUnits()-payment.Requested.AsMajorUnits()*coupon.Discount)*100.0), payment.Requested.Currency().Code)
	if err := srv.Payment.Update(ctx, payment); err != nil {
		return nil, errors.Wrap(err, "update payment")
	}

	apply := core.NewCouponApply(coupon.ID, payment.ID)
	if err := srv.CouponApply.Add(ctx, apply); err != nil {
		return nil, errors.Wrap(err, "create coupon apply")
	}

	return payment, nil
}

func (srv *Service) createApplicationPayment(
	ctx context.Context,
	user *core.User,
	id core.LotID,
	gatewayName string,
	coupon string,
) (*payment.Form, error) {
	invoice, err := srv.GetApplicationInvoice(ctx, user, id)
	if err != nil {
		return nil, errors.Wrap(err, "get application invoice")
	}

	gateway := srv.Gateways.Get(gatewayName)
	if gateway == nil {
		return nil, ErrGatewayNotFound
	}

	pm := core.NewPayment(
		core.PaymentPurposeApplication,
		user.ID,
		invoice.Lot.ID,
		gatewayName,
		invoice.Price,
	)

	if err := srv.Payment.Add(ctx, pm); err != nil {
		return nil, errors.Wrap(err, "add payment to store")
	}

	if coupon != "" {
		c, err := srv.Coupon.Query().Code(coupon).IsDeleted(false).One(ctx)
		if err != nil {
			return nil, errors.Wrap(err, "get coupon")
		}

		pm, err = srv.ApplyCoupon(ctx, user, c, pm)
		if err != nil {
			return nil, errors.Wrap(err, "apply coupon")
		}

		if c.IsFullDiscounted() {
			if err := srv.onPayment(ctx, pm); err != nil {
				return nil, err
			}

			vs := url.Values{}

			vs.Set("status", "success")

			return &payment.Form{
				ExternalID: "",
				Method:     http.MethodGet,
				Values:     vs,
				Action:     srv.Config.SuccessPaymentURL,
			}, nil
		}
	}

	form, err := gateway.NewPayment(ctx, user, pm)
	if err != nil {
		return nil, errors.Wrap(err, "new payment")
	}

	return form, nil
}

func (srv *Service) CreateChangePricePayment(
	ctx context.Context,
	user *core.User,
	id core.LotID,
	gatewayName string,
	changePrice *money.Money,
	coupon string,
) (*payment.Form, error) {
	var form *payment.Form

	err := srv.Txier(ctx, func(ctx context.Context) error {
		var err error
		form, err = srv.createChangePricePayment(ctx, user, id, gatewayName, changePrice, coupon)
		if err != nil {
			return err
		}
		return nil
	})
	return form, err
}

func (srv *Service) createChangePricePayment(
	ctx context.Context,
	user *core.User,
	id core.LotID,
	gatewayName string,
	changePrice *money.Money,
	coupon string,
) (*payment.Form, error) {
	invoice, err := srv.GetChangeInvoice(ctx, user, id)
	if err != nil {
		return nil, errors.Wrap(err, "get change invoice")
	}

	gateway := srv.Gateways.Get(gatewayName)
	if gateway == nil {
		return nil, ErrGatewayNotFound
	}

	pm := core.NewPayment(
		core.PaymentPurposeChangePrice,
		user.ID,
		invoice.Lot.ID,
		gatewayName,
		invoice.Price,
	)
	pm.Metadata = make(map[string]string)
	pm.SetChangePrice(changePrice)
	if err := srv.Payment.Add(ctx, pm); err != nil {
		return nil, errors.Wrap(err, "add payment to store")
	}

	if coupon != "" {
		c, err := srv.Coupon.Query().Code(coupon).IsDeleted(false).One(ctx)
		if err != nil {
			return nil, errors.Wrap(err, "get coupon")
		}

		pm, err = srv.ApplyCoupon(ctx, user, c, pm)
		if err != nil {
			return nil, errors.Wrap(err, "apply coupon")
		}

		if c.IsFullDiscounted() {
			if err := srv.onPayment(ctx, pm); err != nil {
				return nil, err
			}

			vs := url.Values{}

			vs.Set("status", "success")

			return &payment.Form{
				ExternalID: "",
				Method:     http.MethodGet,
				Values:     vs,
				Action:     srv.Config.SuccessPaymentURL,
			}, nil
		}
	}

	form, err := gateway.NewPayment(ctx, user, pm)
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

func (srv *Service) ProcessUnitPayNotification(ctx context.Context, notify *payment.GatewayNotification) (string, error) {
	var (
		message string
		err     error
	)
	if notify.Status.IsCheck() {
		_, err := srv.Payment.Query().ID(notify.PaymentID).One(ctx)
		if err != nil {
			return "", errors.Wrap(err, "query payment")
		}
		return unitPaySuccess, nil
	}
	err = srv.Txier(ctx, func(ctx context.Context) error {
		message, err = srv.processUnitPayNotification(ctx, notify)
		return err
	})

	return message, err
}

func (srv *Service) processUnitPayNotification(ctx context.Context, notify *payment.GatewayNotification) (string, error) {
	p, err := srv.Payment.Query().ID(notify.PaymentID).One(ctx)
	if err != nil {
		return "", errors.Wrap(err, "query payment")
	}
	if p.Status.IsSuccess() {
		return unitPaySuccess, nil
	}
	p.ExternalID.SetValid(notify.ExternalID)
	p.Status = notify.Status
	if len(p.Metadata) != 0 {
		for k, v := range notify.Metadata {
			p.Metadata[k] = v
		}
	} else {
		p.Metadata = notify.Metadata
	}

	switch {
	case p.Status.IsSuccess():
		if equal, _ := p.Requested.Equals(notify.Requested); !equal {
			return "", core.NewError("payment_requested_is_not_equal", "payment requested amount invalid")
		}
		p.Paid = notify.Paid
		p.Received = notify.Received
		if err := srv.onPayment(ctx, p); err != nil {
			return "", err
		}
	case p.Status.IsError():
		log.Warn(ctx, "unit pay payment wasn't succeeded, but it might get success in the future", "id", p.ID, "errMsg", p.Metadata["error"])
	}

	if err := srv.Payment.Update(ctx, p); err != nil {
		return "", errors.Wrap(err, "update payment")
	}

	return unitPaySuccess, nil
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
		return srv.onPayment(ctx, payment)
	}

	return nil
}

func (srv *Service) onPayment(ctx context.Context, pm *core.Payment) error {
	switch pm.Purpose {
	case core.PaymentPurposeApplication:
		return srv.onPaymentApplication(ctx, pm)
	case core.PaymentPurposeChangePrice:
		return srv.onPaymentChangePrice(ctx, pm)
	default:
		return errors.New("unknown purpose of payment")
	}
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

	go func() {
		ctx := context.Background()

		srv.Notify.SendUser(pm.PayerID, userNewPaymentNotification{
			Lot: lot,
		})

		coupon, err := srv.getCouponOfPayment(ctx, pm.ID)
		if err != nil {
			log.Error(ctx, "fetch coupon failed", "err", err)
			return
		}

		srv.Notify.Send(adminNewPaymentNotification{
			Payment: pm,
			Lot:     lot,
			Coupon:  coupon,
		})
	}()

	return nil
}

func (srv *Service) onPaymentChangePrice(ctx context.Context, pm *core.Payment) error {
	lot, err := srv.Lot.Query().ID(pm.LotID).One(ctx)
	if err != nil {
		return errors.Wrap(err, "query payment")
	}

	lot.Price.Previous = null.IntFrom(lot.Price.Current)
	price, err := pm.GetChangePrice()
	if err != nil {
		return errors.Wrap(err, "get change price")
	}

	lot.Price.Current = price
	if err := srv.Lot.Update(ctx, lot); err != nil {
		return errors.Wrap(err, "update lot")
	}

	go func() {
		ctx := context.Background()

		coupon, err := srv.getCouponOfPayment(ctx, pm.ID)
		if err != nil {
			log.Error(ctx, "fetch coupon failed", "err", err)
			return
		}

		srv.Notify.Send(adminNewPaymentNotification{
			Payment: pm,
			Lot:     lot,
			Coupon:  coupon,
		})
	}()

	return nil
}

func (srv *Service) getCouponOfPayment(ctx context.Context, id core.PaymentID) (*core.Coupon, error) {
	apply, err := srv.CouponApply.Query().PaymentID(id).One(ctx)

	if err != nil {
		if err == core.ErrCouponApplyNotFound {
			return nil, nil
		}

		return nil, errors.Wrap(err, "query apply")
	}

	coupon, err := srv.Coupon.Query().ID(apply.CouponID).One(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "query coupon")
	}

	return coupon, nil
}
