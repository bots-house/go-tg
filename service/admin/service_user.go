package admin

import (
	"context"

	"github.com/bots-house/birzzha/core"
	"github.com/pkg/errors"
)

type FullUser struct {
	*core.User
	Lots int
}

type FullUserList struct {
	Total int
	Items []*FullUser
}

var (
	ErrAdminUserCantToggleSelf = core.NewError("admin_user_cant_toggle_self", "admin user cant toggle self")
)

func (srv *Service) newFullUser(user *core.User, lotsCountByUser *core.LotsCountByUser) *FullUser {
	fullUser := &FullUser{
		User: user,
	}
	if lotsCountByUser != nil {
		fullUser.Lots = lotsCountByUser.Lots
	} else {
		fullUser.Lots = 0
	}
	return fullUser
}

func (srv *Service) newFullUserList(users core.UserSlice, total int, lotsCountByUser core.LotsCountByUserSlice) *FullUserList {
	fullUsers := make([]*FullUser, len(users))
	for i, user := range users {
		fullUsers[i] = srv.newFullUser(user, lotsCountByUser.Find(user.ID))
	}

	return &FullUserList{
		Total: total,
		Items: fullUsers,
	}
}

func (srv *Service) GetUsers(ctx context.Context, user *core.User, offset int, limit int) (*FullUserList, error) {
	if err := srv.IsAdmin(user); err != nil {
		return nil, err
	}

	total, err := srv.User.Query().Count(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get users count")
	}

	users, err := srv.User.Query().Offset(offset).Limit(limit).OrderByJoinedAt().All(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get users")
	}

	userIDs := make([]core.UserID, len(users))
	for i, user := range users {
		userIDs[i] = user.ID
	}

	lotsCountByUser, err := srv.Lot.CountByUser(ctx, userIDs...)
	if err != nil {
		return nil, errors.Wrap(err, "get lots by user id's")
	}

	return srv.newFullUserList(users, total, lotsCountByUser), nil
}

func (srv *Service) ToggleUserAdmin(ctx context.Context, user *core.User, id core.UserID) (*FullUser, error) {
	if err := srv.IsAdmin(user); err != nil {
		return nil, err
	}

	if user.ID == id {
		return nil, ErrAdminUserCantToggleSelf
	}

	usr, err := srv.User.Query().ID(id).One(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get user")
	}

	usr.IsAdmin = !usr.IsAdmin

	if err := srv.User.Update(ctx, usr); err != nil {
		return nil, errors.Wrap(err, "update user")
	}

	lotsCountByUser, err := srv.Lot.CountByUser(ctx, usr.ID)
	if err != nil {
		return nil, errors.Wrap(err, "get lots by user id's")
	}

	return srv.newFullUser(usr, lotsCountByUser.Find(usr.ID)), nil
}
