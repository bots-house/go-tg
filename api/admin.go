package api

import (
	"github.com/bots-house/birzzha/api/authz"
	adminops "github.com/bots-house/birzzha/api/gen/restapi/operations/admin"
	"github.com/bots-house/birzzha/api/models"
	"github.com/bots-house/birzzha/core"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/pkg/errors"
)

func (h *Handler) adminDeleteReview(params adminops.AdminDeleteReviewParams, identity *authz.Identity) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	if err := h.Admin.DeleteReview(ctx, identity.User, core.ReviewID(int(params.ID))); err != nil {
		if err2, ok := errors.Cause(err).(*core.Error); ok {
			return adminops.NewAdminDeleteReviewBadRequest().WithPayload(models.NewError(err2))
		}
		return adminops.NewAdminDeleteReviewInternalServerError().WithPayload(models.NewInternalServerError(err))
	}
	return adminops.NewAdminDeleteReviewNoContent()
}

func (h *Handler) adminUpdateReview(params adminops.AdminUpdateReviewParams, identity *authz.Identity) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	result, err := h.Admin.UpdateReview(ctx, identity.User, core.ReviewID(int(params.ID)), swag.StringValue(params.Body.Text))
	if err != nil {
		if err2, ok := errors.Cause(err).(*core.Error); ok {
			return adminops.NewAdminUpdateReviewBadRequest().WithPayload(models.NewError(err2))
		}
		return adminops.NewAdminUpdateReviewInternalServerError().WithPayload(models.NewInternalServerError(err))
	}
	return adminops.NewAdminUpdateReviewOK().WithPayload(models.NewReview(h.Storage, result))
}

func (h *Handler) adminGetReviews(params adminops.AdminGetReviewsParams, identity *authz.Identity) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	result, err := h.Admin.GetReviews(ctx, identity.User, int(swag.Int64Value(params.Offset)), int(swag.Int64Value(params.Limit)))
	if err != nil {
		if err2, ok := errors.Cause(err).(*core.Error); ok {
			return adminops.NewAdminGetReviewsBadRequest().WithPayload(models.NewError(err2))
		}
		return adminops.NewAdminGetReviewsInternalServerError().WithPayload(models.NewInternalServerError(err))
	}
	return adminops.NewAdminGetReviewsOK().WithPayload(models.NewAdminReviewList(h.Storage, result))
}

func (h *Handler) adminGetUsers(params adminops.AdminGetUsersParams, identity *authz.Identity) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	users, err := h.Admin.GetUsers(ctx, identity.User, int(swag.Int64Value(params.Offset)), int(swag.Int64Value(params.Limit)))
	if err != nil {
		if err2, ok := errors.Cause(err).(*core.Error); ok {
			return adminops.NewAdminGetUsersBadRequest().WithPayload(models.NewError(err2))
		}
		return adminops.NewAdminGetUsersInternalServerError().WithPayload(models.NewInternalServerError(err))
	}
	return adminops.NewAdminGetUsersOK().WithPayload(models.NewAdminFullUserList(h.Storage, users))

}

func (h *Handler) toggleUserAdmin(params adminops.ToggleUserAdminParams, identity *authz.Identity) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	user, err := h.Admin.ToggleUserAdmin(ctx, identity.User, core.UserID(int(params.ID)))
	if err != nil {
		if err2, ok := errors.Cause(err).(*core.Error); ok {
			return adminops.NewToggleUserAdminBadRequest().WithPayload(models.NewError(err2))
		}
		return adminops.NewToggleUserAdminInternalServerError().WithPayload(models.NewInternalServerError(err))
	}
	return adminops.NewToggleUserAdminOK().WithPayload(models.NewAdminFullUser(h.Storage, user))

}
