package interkassa

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"

	"github.com/Rhymond/go-money"
	"github.com/pkg/errors"

	"github.com/bots-house/birzzha/core"
	"github.com/bots-house/birzzha/pkg/log"
	"github.com/bots-house/birzzha/service/payment"
)

type Gateway struct {
	CheckoutID    string
	SecretKey     string
	TestSecretKey string

	NotificationURL string
	SuccessURL      string
	PendingURL      string
	FailedURL       string
}

func injectQueryParam(u, k, v string) string {
	uri, err := url.Parse(u)
	if err != nil {
		return u
	}

	query := uri.Query()
	query.Add(k, v)

	uri.RawQuery = query.Encode()

	return uri.String()
}

const testPayWay = "test_interkassa_test_xts"

func (gw *Gateway) isTest() bool {
	return gw.TestSecretKey != ""
}

func (gw *Gateway) Name() string {
	return "interkassa"
}

func (gw *Gateway) NewPayment(ctx context.Context, payer *core.User, pm *core.Payment) (*payment.Form, error) {
	// https://docs.interkassa.com/#section/3.-Protokol/3.2.-Forma-zaprosa-platezha
	vs := url.Values{}

	// set checkout id
	vs.Set("ik_co_id", gw.CheckoutID)

	// set pm id
	var paymentID string
	if gw.isTest() {
		paymentID = "TEST-" + strconv.Itoa(int(pm.ID))
	} else {
		paymentID = strconv.Itoa(int(pm.ID))
	}

	vs.Set("ik_pm_no", paymentID)

	// set currency
	vs.Set("ik_cur", pm.Requested.Currency().Code)

	// set amount
	amount := fmt.Sprintf("%.2f", pm.Requested.AsMajorUnits())
	vs.Set("ik_am", amount)

	// pm description
	var desc string

	switch pm.Purpose {
	case core.PaymentPurposeApplication:
		desc = fmt.Sprintf("Оплата размещения лота №%d", pm.LotID)
	case core.PaymentPurposeChangePrice:
		desc = fmt.Sprintf("Оплата изменения цены лота №%d", pm.LotID)
	default:
		desc = "Оплата"
	}

	vs.Set("ik_desc", desc)

	ourPaymentID := strconv.Itoa(int(pm.ID))

	// Interaction URL
	vs.Set("ik_ia_u", gw.NotificationURL)
	vs.Set("ik_ia_m", http.MethodPost)

	// Success URL
	vs.Set("ik_suc_u", injectQueryParam(gw.SuccessURL, "payment_id", ourPaymentID))
	vs.Set("ik_suc_m", http.MethodGet)

	// Pending URL
	vs.Set("ik_pnd_u", injectQueryParam(gw.PendingURL, "payment_id", ourPaymentID))
	vs.Set("ik_pnd_m", http.MethodGet)

	// Fail URL
	vs.Set("ik_fal_u", injectQueryParam(gw.FailedURL, "payment_id", ourPaymentID))
	vs.Set("ik_fal_m", http.MethodGet)

	// Meta
	vs.Set("ik_x_lot_id", strconv.Itoa(int(pm.LotID)))
	vs.Set("ik_x_payer_id", strconv.Itoa(int(pm.PayerID)))
	vs.Set("ik_x_purpose", pm.Purpose.String())

	if gw.isTest() {
		vs.Set("ik_pw_on", testPayWay)
	} else {
		vs.Set("ik_pw_off", testPayWay)
	}

	sign, err := gw.signature(vs, gw.SecretKey)
	if err != nil {
		return nil, errors.Wrap(err, "calc signature")
	}

	vs.Set("ik_sign", sign)

	return &payment.Form{
		ExternalID: "",
		Method:     http.MethodPost,
		Action:     "https://sci.interkassa.com/",
		Values:     vs,
	}, nil
}

func (gw *Gateway) signature(vs url.Values, secretKey string) (string, error) {

	// sort keys
	keys := make([]string, 0, len(vs))
	for key := range vs {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	// build ordered values list
	values := make([]string, len(keys))
	for i, key := range keys {
		values[i] = vs.Get(key)
	}

	values = append(values, secretKey)

	data := strings.Join(values, ":")

	// hashing
	hash := sha256.New()
	if _, err := io.WriteString(hash, data); err != nil {
		return "", errors.Wrap(err, "write data to hash")
	}

	hashSum := hash.Sum(nil)

	sign := base64.StdEncoding.EncodeToString(hashSum)

	return sign, nil
}

var (
	ErrInvalidNotificationMethod      = core.NewError("invalid_notification_method", "invalid notification method (should be POST)")
	ErrInvalidNotificationContentType = core.NewError("invalid_notification_content_type", "invalid notification content-type")
	ErrInvalidNotificationSignature   = core.NewError("invalid_notification_signature", "invalid notification signature")
	ErrInvalidNotificationCheckoutID  = core.NewError("invalid_notification_checkout_id", "invalid notification checkout id")
	ErrInvalidNotificationData        = core.NewError("invalid_notification_data", "invalid notification data")
)

func (gw *Gateway) parsePayWayCurrency(payway string) string {
	parts := strings.Split(payway, "_")

	currency := parts[len(parts)-1]

	return currency
}

func (gw *Gateway) ParseNotification(ctx context.Context, r *http.Request) (*payment.GatewayNotification, error) {
	if r.Method != http.MethodPost {
		return nil, ErrInvalidNotificationMethod
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, errors.Wrap(err, "fail to read body")
	}

	if r.Header.Get("Content-Type") != "application/x-www-form-urlencoded" {
		return nil, ErrInvalidNotificationContentType
	}

	// parse body
	values, err := url.ParseQuery(string(body))
	if err != nil {
		return nil, errors.Wrap(err, "parse body")
	}

	// determine secret used for signature
	payWay := values.Get("ik_pw_via")

	var secret string

	if payWay == testPayWay {
		secret = gw.TestSecretKey
	} else {
		secret = gw.SecretKey
	}

	signature := values.Get("ik_sign")
	values.Del("ik_sign")

	// calculate and compare signature
	exceptedSignature, err := gw.signature(values, secret)
	if err != nil {
		return nil, ErrInvalidNotificationSignature
	}

	if signature != exceptedSignature {
		return nil, ErrInvalidNotificationSignature
	}

	notification := &payment.GatewayNotification{Metadata: make(map[string]string)}

	notification.Metadata["ik_checkout_id"] = values.Get("ik_co_id")
	notification.Metadata["ik_checkout_purse_id"] = values.Get("ik_co_prs_id")

	notification.Metadata["ik_pay_way"] = payWay

	notification.Metadata["ik_invoice_created_at"] = values.Get("ik_inv_crt")
	notification.Metadata["ik_invoice_processed_at"] = values.Get("ik_inv_prc")

	if gw.CheckoutID != notification.Metadata["ik_checkout_id"] {
		return nil, ErrInvalidNotificationCheckoutID
	}

	notification.ExternalID = values.Get("ik_inv_id")

	rawPaymentID := values.Get("ik_pm_no")

	if gw.isTest() {
		rawPaymentID = strings.TrimPrefix(rawPaymentID, "TEST-")
	}

	paymentID, err := strconv.Atoi(rawPaymentID)
	if err != nil {
		log.Warn(ctx, "invalid payment id", "id", rawPaymentID)
		return nil, ErrInvalidNotificationData
	}

	notification.PaymentID = core.PaymentID(paymentID)

	switch values.Get("ik_inv_st") {
	case "success":
		notification.Status = core.PaymentStatusSuccess
	case "fail", "canceled":
		notification.Status = core.PaymentStatusFailed
	default:
		log.Warn(ctx, "invalid payment status", "id", rawPaymentID, "status", values.Get("ik_inv_st"))
		return nil, ErrInvalidNotificationData
	}

	// requested money
	requestedAmount, err := strconv.ParseFloat(values.Get("ik_am"), 64)
	if err != nil {
		log.Warn(ctx, "invalid requested amount", "id", rawPaymentID, "requested_amount", values.Get("ik_am"))
		return nil, ErrInvalidNotificationData
	}

	requestedCurrency := values.Get("ik_cur")

	notification.Requested = money.New(int64(requestedAmount*100), requestedCurrency)

	if !notification.Status.IsFinal() {
		return notification, nil
	}

	// Parse paid
	paidAmount, err := strconv.ParseFloat(values.Get("ik_ps_price"), 64)
	if err != nil {
		log.Warn(ctx, "invalid paid amount", "id", rawPaymentID, "paid_amount", values.Get("ik_ps_price"))
		return nil, ErrInvalidNotificationData
	}
	paidCurrency := gw.parsePayWayCurrency(payWay)

	notification.Paid = money.New(int64(paidAmount*100), paidCurrency)

	// Received money.
	// Interkassa refund money in payment currency, so we get currency from paid.
	receivedAmount, err := strconv.ParseFloat(values.Get("ik_co_rfn"), 64)
	if err != nil {
		log.Warn(ctx, "invalid received amount", "id", rawPaymentID, "received_amount", values.Get("ik_co_rfn"))

		return nil, ErrInvalidNotificationData
	}

	notification.Received = money.New(int64(receivedAmount*100), notification.Requested.Currency().Code)

	return notification, nil
}

func (gw *Gateway) Refund(ctx context.Context, id string) error {
	return nil
}

func (gw *Gateway) TestMode() bool {
	panic("implement me")
}

func (gw *Gateway) SetTestMode(enabled bool) {
	panic("implement me")
}
