package postgres

import (
	"context"
	"database/sql"

	"github.com/bots-house/birzzha/core"
	"github.com/bots-house/birzzha/store/postgres/dal"
	"github.com/bots-house/birzzha/store/postgres/shared"
	"github.com/pkg/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type UserStore struct {
	boil.ContextExecutor
}

func (store *UserStore) toRow(user *core.User) *dal.User {
	return &dal.User{
		ID:                   int(user.ID),
		TelegramID:           user.Telegram.ID,
		TelegramUsername:     user.Telegram.Username,
		TelegramLanguageCode: user.Telegram.LanguageCode,
		Avatar:               user.Avatar,
		FirstName:            user.FirstName,
		LastName:             user.LastName,
		IsNameEdited:         user.IsNameEdited,
		JoinedFrom:           null.NewString(string(user.JoinedFrom), user.JoinedFrom != ""),
		IsAdmin:              user.IsAdmin,
		JoinedAt:             user.JoinedAt,
		UpdatedAt:            user.UpdatedAt,
	}
}

func (store *UserStore) fromRow(row *dal.User) *core.User {
	return &core.User{
		ID: core.UserID(row.ID),
		Telegram: core.UserTelegram{
			ID:           row.TelegramID,
			Username:     row.TelegramUsername,
			LanguageCode: row.TelegramLanguageCode,
		},
		Avatar:       row.Avatar,
		FirstName:    row.FirstName,
		LastName:     row.LastName,
		IsNameEdited: row.IsNameEdited,
		JoinedFrom:   core.JoinedFrom(row.JoinedFrom.String),
		IsAdmin:      row.IsAdmin,
		JoinedAt:     row.JoinedAt,
		UpdatedAt:    row.UpdatedAt,
	}
}

func (store *UserStore) fromRowSlice(rows dal.UserSlice) core.UserSlice {
	result := make(core.UserSlice, len(rows))
	for i, row := range rows {
		result[i] = store.fromRow(row)
	}
	return result
}

func (store *UserStore) Add(ctx context.Context, user *core.User) error {
	row := store.toRow(user)
	if err := row.Insert(ctx, shared.GetExecutorOrDefault(ctx, store.ContextExecutor), boil.Infer()); err != nil {
		return errors.Wrap(err, "insert query")
	}
	*user = *store.fromRow(row)
	return nil
}

func (store *UserStore) Update(ctx context.Context, user *core.User) error {
	row := store.toRow(user)
	n, err := row.Update(ctx, shared.GetExecutorOrDefault(ctx, store.ContextExecutor), boil.Infer())
	if err != nil {
		return errors.Wrap(err, "update query")
	}
	if n == 0 {
		return core.ErrUserNotFound
	}
	return nil
}

func (store *UserStore) Query() core.UserStoreQuery {
	return &userStoreQuery{store: store}
}

type userStoreQuery struct {
	mods  []qm.QueryMod
	store *UserStore
}

func (usq *userStoreQuery) TelegramID(id int) core.UserStoreQuery {
	usq.mods = append(usq.mods, dal.UserWhere.TelegramID.EQ(id))
	return usq
}

func (usq *userStoreQuery) Limit(limit int) core.UserStoreQuery {
	usq.mods = append(usq.mods, qm.Limit(limit))
	return usq
}

func (usq *userStoreQuery) Offset(offset int) core.UserStoreQuery {
	usq.mods = append(usq.mods, qm.Offset(offset))
	return usq
}

func (usq *userStoreQuery) OrderByJoinedAt() core.UserStoreQuery {
	usq.mods = append(usq.mods, qm.OrderBy("joined_at DESC"))
	return usq
}

func (usq *userStoreQuery) ID(ids ...core.UserID) core.UserStoreQuery {
	idsInt := make([]int, len(ids))
	for i, id := range ids {
		idsInt[i] = int(id)
	}

	usq.mods = append(usq.mods, dal.UserWhere.ID.IN(idsInt))
	return usq
}

func (usq *userStoreQuery) One(ctx context.Context) (*core.User, error) {
	executor := shared.GetExecutorOrDefault(ctx, usq.store.ContextExecutor)

	row, err := dal.Users(usq.mods...).One(ctx, executor)
	if err == sql.ErrNoRows {
		return nil, core.ErrUserNotFound
	} else if err != nil {
		return nil, err
	}

	return usq.store.fromRow(row), nil
}

func (usq *userStoreQuery) Count(ctx context.Context) (int, error) {
	executor := shared.GetExecutorOrDefault(ctx, usq.store.ContextExecutor)

	count, err := dal.Users(usq.mods...).Count(ctx, executor)

	return int(count), err
}

func (usq *userStoreQuery) All(ctx context.Context) (core.UserSlice, error) {
	executor := shared.GetExecutorOrDefault(ctx, usq.store.ContextExecutor)
	rows, err := dal.Users(usq.mods...).All(ctx, executor)
	if err != nil {
		return nil, err
	}

	return usq.store.fromRowSlice(rows), nil
}
