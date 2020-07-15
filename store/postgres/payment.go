package postgres

import (
	"context"
	"database/sql"

	"github.com/Rhymond/go-money"
	"github.com/pkg/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/bots-house/birzzha/core"
	"github.com/bots-house/birzzha/store/postgres/dal"
	"github.com/bots-house/birzzha/store/postgres/shared"
)

type PaymentStore struct {
	boil.ContextExecutor
}

func (store *PaymentStore) toRow(in *core.Payment) (*dal.Payment, error) {
	result := &dal.Payment{
		ID:         int(in.ID),
		ExternalID: in.ExternalID,
		Gateway:    in.Gateway,
		Purpose:    in.Purpose.String(),
		PayerID:    int(in.PayerID),
		LotID:      int(in.LotID),
		Status:     in.Status.String(),
		Requested:  nil,
		Paid:       null.JSON{},
		Received:   null.JSON{},
		Metadata:   null.JSON{},
		CreatedAt:  in.CreatedAt,
		UpdatedAt:  in.UpdatedAt,
	}

	if err := result.Requested.Marshal(in.Requested); err != nil {
		return nil, errors.Wrap(err, "marshal request")
	}

	if err := result.Paid.Marshal(in.Paid); err != nil {
		return nil, errors.Wrap(err, "marshal paid")
	}

	if err := result.Received.Marshal(in.Received); err != nil {
		return nil, errors.Wrap(err, "marshal paid")
	}

	if err := result.Metadata.Marshal(in.Metadata); err != nil {
		return nil, errors.Wrap(err, "marshal metadata")
	}

	return result, nil
}

func (store *PaymentStore) fromRow(in *dal.Payment) (*core.Payment, error) {

	purpose, err := core.ParsePaymentPurpose(in.Purpose)
	if err != nil {
		return nil, errors.Wrap(err, "parse payment purpose")
	}

	status, err := core.ParsePaymentStatus(in.Status)
	if err != nil {
		return nil, errors.Wrap(err, "parse payment status")
	}

	result := &core.Payment{
		ID:         core.PaymentID(in.ID),
		ExternalID: in.ExternalID,
		Gateway:    in.Gateway,
		Purpose:    purpose,
		PayerID:    core.UserID(in.PayerID),
		LotID:      core.LotID(in.LotID),
		Status:     status,
		Requested:  &money.Money{},
		Paid:       nil,
		Received:   nil,
		Metadata:   nil,
		CreatedAt:  in.CreatedAt,
		UpdatedAt:  in.UpdatedAt,
	}

	if err := in.Requested.Unmarshal(result.Requested); err != nil {
		return nil, errors.Wrap(err, "unmarshal requested")
	}

	if !in.Paid.IsZero() {
		result.Paid = &money.Money{}

		if err := in.Paid.Unmarshal(result.Paid); err != nil {
			return nil, errors.Wrap(err, "unmarshal paid")
		}
	}

	if !in.Received.IsZero() {
		result.Received = &money.Money{}

		if err := in.Received.Unmarshal(result.Received); err != nil {
			return nil, errors.Wrap(err, "unmarshal paid")
		}
	}

	if !in.Metadata.IsZero() {
		result.Metadata = map[string]string{}

		if err := in.Metadata.Unmarshal(&result.Metadata); err != nil {
			return nil, errors.Wrap(err, "unmarshal metadata")
		}
	}

	return result, nil
}

func (store *PaymentStore) Add(ctx context.Context, payment *core.Payment) error {
	row, err := store.toRow(payment)
	if err != nil {
		return errors.Wrap(err, "to row")
	}

	executor := shared.GetExecutorOrDefault(ctx, store.ContextExecutor)
	if err := row.Insert(ctx, executor, boil.Infer()); err != nil {
		return errors.Wrap(err, "insert query")
	}

	payment.ID = core.PaymentID(row.ID)

	return nil
}

func (store *PaymentStore) Update(ctx context.Context, payment *core.Payment) error {
	row, err := store.toRow(payment)
	if err != nil {
		return errors.Wrap(err, "to row")
	}
	executor := shared.GetExecutorOrDefault(ctx, store.ContextExecutor)
	n, err := row.Update(ctx, executor, boil.Infer())
	if err != nil {
		return errors.Wrap(err, "update query")
	}
	if n == 0 {
		return core.ErrPaymentNotFound
	}

	return nil
}

func (store *PaymentStore) Query() core.PaymentStoreQuery {
	return &PaymentStoreQuery{store: store}
}

type PaymentStoreQuery struct {
	mods  []qm.QueryMod
	store *PaymentStore
}

func (psq *PaymentStoreQuery) ID(id core.PaymentID) core.PaymentStoreQuery {
	psq.mods = append(psq.mods, dal.PaymentWhere.ID.EQ(int(id)))
	return psq
}

func (psq *PaymentStoreQuery) PayerID(id core.UserID) core.PaymentStoreQuery {
	psq.mods = append(psq.mods, dal.PaymentWhere.PayerID.EQ(int(id)))
	return psq
}

func (psq *PaymentStoreQuery) One(ctx context.Context) (*core.Payment, error) {
	executor := shared.GetExecutorOrDefault(ctx, psq.store.ContextExecutor)

	row, err := dal.Payments(psq.mods...).One(ctx, executor)
	if err == sql.ErrNoRows {
		return nil, core.ErrPaymentNotFound
	} else if err != nil {
		return nil, err
	}

	return psq.store.fromRow(row)
}

func (psq *PaymentStoreQuery) All(ctx context.Context) ([]*core.Payment, error) {
	executor := shared.GetExecutorOrDefault(ctx, psq.store.ContextExecutor)

	rows, err := dal.Payments(psq.mods...).All(ctx, executor)
	if err != nil {
		return nil, err
	}

	result := make([]*core.Payment, len(rows))
	for i, row := range rows {
		result[i], err = psq.store.fromRow(row)
		if err != nil {
			return nil, errors.Wrap(err, "from row")
		}
	}

	return result, nil
}
