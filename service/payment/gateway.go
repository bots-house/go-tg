package payment

import (
	"context"
	"net/http"
	"net/url"

	"github.com/Rhymond/go-money"
	"github.com/bots-house/birzzha/core"
)

var (
	ErrInvalidNotificationSignature = core.NewError("invalid_notification_signature", "invalid notification signature")
	ErrInvalidNotificationData      = core.NewError("invalid_notification_data", "invalid notification data")
)

type Form struct {
	ExternalID string
	Method     string
	Action     string
	Values     url.Values
}

// GatewayNotification it's S2S notification about payment result.
type GatewayNotification struct {
	// Related Payment identifier.
	PaymentID core.PaymentID

	// True, if payment is success
	Status core.PaymentStatus

	ExternalID string

	// Amount requests
	Requested *money.Money

	// Amount user pay
	Paid *money.Money

	// Amount refund
	Received *money.Money

	// Additional metadata.
	Metadata map[string]string
}

// Gateway define generic interface for payment gateways.
type Gateway interface {
	// Name of gateway
	Name() string

	// Make payment form based on payment and user
	NewPayment(ctx context.Context, payer *core.User, payment *core.Payment) (*Form, error)

	// ParseNotification parse HTTP request and returns parsed gateway notification.
	// If status of payment is not final, parsing of paid, received and metadata should be skipping.
	ParseNotification(ctx context.Context, r *http.Request) (*GatewayNotification, error)

	// TestMode returns true, if test mode activated.
	TestMode() bool

	// Refund payment
	Refund(ctx context.Context, id string) error

	// SetTestMode enable or disable test mode.
	SetTestMode(enabled bool)
}
