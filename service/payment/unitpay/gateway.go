package unitpay

import (
	"context"
	"crypto/sha256"
	"fmt"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"

	"github.com/Rhymond/go-money"

	"github.com/bots-house/birzzha/pkg/log"

	"github.com/bots-house/birzzha/core"
	"github.com/bots-house/birzzha/service/payment"
)

const (
	methodParam = "method"
)

var (
	ErrInvalidNotificationMethod = core.NewError("invalid_notification_method", "invalid notification method (must be GET)")
)

type Gateway struct {
	PublicKey string
	SecretKey string
}

func (g *Gateway) Name() string {
	return "unitpay"
}

func (g *Gateway) NewPayment(ctx context.Context, payer *core.User, pm *core.Payment) (*payment.Form, error) {
	const (
		sumParam       = "sum"
		descParam      = "desc"
		accountParam   = "account"
		signatureParam = "signature"
		currencyParam  = "currency"
	)

	vs := url.Values{}
	paymentID := strconv.Itoa(int(pm.ID))
	currency := pm.Requested.Currency().Code
	amount := fmt.Sprintf("%.2f", pm.Requested.AsMajorUnits())

	desc := "Оплата"
	switch pm.Purpose {
	case core.PaymentPurposeApplication:
		desc = "Оплата размещения лота №" + strconv.Itoa(int(pm.LotID))
	case core.PaymentPurposeChangePrice:
		desc = "Оплата изменения цены лота №" + strconv.Itoa(int(pm.LotID))
	}

	vs.Set(accountParam, paymentID)
	vs.Set(currencyParam, currency)
	vs.Set(descParam, desc)
	vs.Set(sumParam, amount)
	signature := g.signature(vs, g.SecretKey)
	vs.Set(signatureParam, signature)

	return &payment.Form{
		ExternalID: "",
		Method:     http.MethodGet,
		Action:     "https://unitpay.money/pay/" + g.PublicKey,
		Values:     vs,
	}, nil
}

func (g *Gateway) signature(vs url.Values, secretKey string) string {
	const up = "{up}"

	keys := make([]string, 0, len(vs))
	for key := range vs {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	values := make([]string, 0, len(keys)+1)
	for _, key := range keys {
		if key == methodParam {
			//prepend method
			values = append([]string{vs.Get(key)}, values...)
			continue
		}
		values = append(values, vs.Get(key))
	}
	values = append(values, secretKey)

	hash := sha256.Sum256([]byte(strings.Join(values, up)))

	return fmt.Sprintf("%x", hash)
}

func (g *Gateway) ParseNotification(ctx context.Context, r *http.Request) (*payment.GatewayNotification, error) {
	if r.Method != http.MethodGet {
		return nil, ErrInvalidNotificationMethod
	}
	values := r.URL.Query()
	log.Debug(ctx, "unitPay notification", "query", values.Encode())

	signature := values.Get("params[signature]")
	values.Del("params[signature]")

	expectedSignature := g.signature(values, g.SecretKey)
	if expectedSignature != signature {
		return nil, payment.ErrInvalidNotificationSignature
	}

	notification := &payment.GatewayNotification{Metadata: map[string]string{
		//ID проекта в UnitPay
		"projectID": values.Get("params[projectId]"),
		//Ваш доход с данного платежа, в рублях
		"profit": values.Get("params[profit]"),
		//Телефон плательщика (передается только для мобильных платежей)
		"phone": values.Get("params[phone]"),
		//Код платежной системы (https://help.unitpay.ru/book-of-reference/payment-system-codes)
		"paymentType": values.Get("params[paymentType]"),
		//Буквенный код оператора (https://help.unitpay.ru/book-of-reference/operator-codes)
		"operator": values.Get("params[operator]"),
		//Дата платежа в формате YYYY-mm-dd HH:ii:ss (например 2012-10-01 12:32:00)
		"date": values.Get("params[date]"),
		//Признак тестового режима, если запрос делается с использованием механизма "тестового запроса", то значение будет равно 1. Для реальных платежей значение всегда 0
		"test": values.Get("params[test]"),
		//Признак наличия 3-DS для операций по карте, флаг присутствует при PAY уведомлениях
		"3ds": values.Get("params[3ds]"),
		//Детализация ошибки (только для метода error)
		"error": values.Get("params[errorMessage]"),
	}}

	notification.ExternalID = values.Get("params[unitpayId]")

	rawPaymentID := values.Get("params[account]")
	paymentID, err := strconv.Atoi(rawPaymentID)
	if err != nil {
		return nil, payment.ErrInvalidNotificationData
	}
	notification.PaymentID = core.PaymentID(paymentID)

	requestedAmount, err := strconv.ParseFloat(values.Get("params[orderSum]"), 64)
	if err != nil {
		log.Warn(ctx, "invalid order sum", "id", notification.PaymentID, "order_sum", values.Get("params[orderSum]"))
		return nil, payment.ErrInvalidNotificationData
	}
	notification.Requested = money.New(int64(requestedAmount*100), values.Get("params[orderCurrency]"))

	paidAmount, err := strconv.ParseFloat(values.Get("params[payerSum]"), 64)
	if err != nil {
		log.Warn(ctx, "invalid order sum", "id", notification.PaymentID, "order_sum", values.Get("params[payerSum]"))
		return nil, payment.ErrInvalidNotificationData
	}
	notification.Paid = money.New(int64(paidAmount*100), values.Get("params[payerCurrency]"))

	switch values.Get(methodParam) {
	case "pay":
		notification.Status = core.PaymentStatusSuccess
	case "check":
		notification.Status = core.PaymentStatusCheck
	case "error":
		notification.Status = core.PaymentStatusError
	default:
		log.Warn(ctx, "invalid unitPay payment status", "status", values.Get(methodParam))
		return nil, payment.ErrInvalidNotificationData
	}

	return notification, nil
}

func (g *Gateway) TestMode() bool {
	panic("implement me!")
}

func (g *Gateway) Refund(ctx context.Context, id string) error {
	panic("implement me!")
}

func (g *Gateway) SetTestMode(enabled bool) {
	panic("implement me!")
}
