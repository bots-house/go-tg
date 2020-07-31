package api

import (
	"github.com/bots-house/birzzha/api/authz"
	adminops "github.com/bots-house/birzzha/api/gen/restapi/operations/admin"
	"github.com/bots-house/birzzha/api/models"
	"github.com/bots-house/birzzha/core"
	"github.com/bots-house/birzzha/service/admin"
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

func (h *Handler) adminGetLotStatuses(params adminops.AdminGetLotStatusesParams, identity *authz.Identity) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	result, err := h.Admin.GetLotStatusesCount(ctx, identity.User, core.UserID(int(swag.Int64Value(params.UserID))))
	if err != nil {
		if err2, ok := errors.Cause(err).(*core.Error); ok {
			return adminops.NewAdminGetLotStatusesBadRequest().WithPayload(models.NewError(err2))
		}
		return adminops.NewAdminGetLotStatusesInternalServerError().WithPayload(models.NewInternalServerError(err))
	}
	return adminops.NewAdminGetLotStatusesOK().WithPayload(models.NewLotStatusesCount(result))

}

func (h *Handler) adminGetLots(params adminops.AdminGetLotsParams, identity *authz.Identity) middleware.Responder {
	ctx := params.HTTPRequest.Context()
	userID := core.UserID(int(swag.Int64Value(params.UserID)))
	status := swag.StringValue(params.Status)
	limit := int(swag.Int64Value(params.Limit))
	offset := int(swag.Int64Value(params.Offset))

	result, err := h.Admin.GetLots(ctx, identity.User, userID, status, limit, offset)
	if err != nil {
		if err2, ok := errors.Cause(err).(*core.Error); ok {
			return adminops.NewAdminGetLotsBadRequest().WithPayload(models.NewError(err2))
		}
		return adminops.NewAdminGetLotsInternalServerError().WithPayload(models.NewInternalServerError(err))
	}
	return adminops.NewAdminGetLotsOK().WithPayload(models.NewAdminLotItemList(h.Storage, result))
}

func (h *Handler) adminGetSettings(params adminops.AdminGetSettingsParams, identity *authz.Identity) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	result, err := h.Admin.GetSettings(ctx, identity.User)
	if err != nil {
		if err2, ok := errors.Cause(err).(*core.Error); ok {
			return adminops.NewAdminGetSettingsBadRequest().WithPayload(models.NewError(err2))
		}
		return adminops.NewAdminGetSettingsInternalServerError().WithPayload(models.NewInternalServerError(err))
	}
	return adminops.NewAdminGetSettingsOK().WithPayload(models.NewSettings(result))
}

func (h *Handler) adminUpdateSettingsLanding(params adminops.AdminUpdateSettingsLandingParams, identity *authz.Identity) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	result, err := h.Admin.UpdateSettingsLanding(ctx, identity.User, &admin.SettingsInputLanding{
		UniqueUsersPerMonthShift: int(swag.Int64Value(params.Landing.UniqueUsersPerMonthShift)),
		AvgChannelReachShift:     int(swag.Int64Value(params.Landing.AvgChannelReachShift)),
		AvgSiteReachShift:        int(swag.Int64Value(params.Landing.AvgSiteReachShift)),
	})
	if err != nil {
		if err2, ok := errors.Cause(err).(*core.Error); ok {
			return adminops.NewAdminUpdateSettingsLandingBadRequest().WithPayload(models.NewError(err2))
		}
		return adminops.NewAdminUpdateSettingsLandingInternalServerError().WithPayload(models.NewInternalServerError(err))
	}

	return adminops.NewAdminUpdateSettingsLandingOK().WithPayload(models.NewAdminLanding(result))
}

func (h *Handler) adminUpdateSettingsPrices(params adminops.AdminUpdateSettingsPricesParams, identity *authz.Identity) middleware.Responder {
	ctx := params.HTTPRequest.Context()
	result, err := h.Admin.UpdateSettingsPrice(ctx,
		identity.User,
		&admin.SettingsPricesInput{
			Application: models.ToMoney(params.Prices.Application),
			Change:      models.ToMoney(params.Prices.Change),
			Cashier:     swag.StringValue(params.Prices.Cashier),
		},
	)
	if err != nil {
		if err2, ok := errors.Cause(err).(*core.Error); ok {
			return adminops.NewAdminUpdateSettingsPricesBadRequest().WithPayload(models.NewError(err2))
		}
		return adminops.NewAdminUpdateReviewInternalServerError().WithPayload(models.NewInternalServerError(err))
	}

	return adminops.NewAdminUpdateSettingsPricesOK().WithPayload(models.NewSettingsPrices(result))
}

func (h *Handler) adminUpdateSettingsChannel(params adminops.AdminUpdateSettingsChannelParams, identity *authz.Identity) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	result, err := h.Admin.UpdateSettingsChannel(ctx,
		identity.User,
		&admin.SettingsChannelInput{
			PrivateID:      swag.Int64Value(params.Channel.ID),
			PrivateLink:    swag.StringValue(params.Channel.JoinLink),
			PublicUsername: swag.StringValue(params.Channel.PublicUsername),
		},
	)
	if err != nil {
		if err2, ok := errors.Cause(err).(*core.Error); ok {
			return adminops.NewAdminUpdateSettingsChannelBadRequest().WithPayload(models.NewError(err2))
		}
		return adminops.NewAdminUpdateReviewInternalServerError().WithPayload(models.NewInternalServerError(err))
	}
	return adminops.NewAdminUpdateSettingsChannelOK().WithPayload(models.NewSettingsChannel(result))
}

func (h *Handler) adminCreateTopic(params adminops.AdminCreateTopicParams, identity *authz.Identity) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	result, err := h.Admin.CreateTopic(ctx, identity.User, swag.StringValue(params.Topic.Name))
	if err != nil {
		if err2, ok := errors.Cause(err).(*core.Error); ok {
			return adminops.NewAdminCreateTopicBadRequest().WithPayload(models.NewError(err2))
		}
		return adminops.NewAdminCreateTopicInternalServerError().WithPayload(models.NewInternalServerError(err))
	}
	return adminops.NewAdminCreateTopicOK().WithPayload(models.NewAdminFullTopic(result))
}

func (h *Handler) adminUpdateTopic(params adminops.AdminUpdateTopicParams, identity *authz.Identity) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	result, err := h.Admin.UpdateTopic(ctx, identity.User, core.TopicID(int(params.ID)), swag.StringValue(params.Topic.Name))
	if err != nil {
		if err2, ok := errors.Cause(err).(*core.Error); ok {
			return adminops.NewAdminUpdateTopicBadRequest().WithPayload(models.NewError(err2))
		}
		return adminops.NewAdminUpdateTopicInternalServerError().WithPayload(models.NewInternalServerError(err))
	}
	return adminops.NewAdminUpdateTopicOK().WithPayload(models.NewAdminFullTopic(result))
}

func (h *Handler) adminDeleteTopic(params adminops.AdminDeleteTopicParams, identity *authz.Identity) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	if err := h.Admin.DeleteTopic(ctx, identity.User, core.TopicID(int(params.ID))); err != nil {
		if err2, ok := errors.Cause(err).(*core.Error); ok {
			return adminops.NewAdminDeleteTopicBadRequest().WithPayload(models.NewError(err2))
		}
		return adminops.NewAdminDeleteTopicInternalServerError().WithPayload(models.NewInternalServerError(err))
	}
	return adminops.NewAdminDeleteTopicNoContent()
}

func (h *Handler) adminCreateLotCanceledReason(params adminops.AdminCreateLotCanceledReasonParams, identity *authz.Identity) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	result, err := h.Admin.CreateLotCanceledReason(ctx,
		identity.User,
		&admin.LotCanceledReasonInput{
			Why:      swag.StringValue(params.Reason.Why),
			IsPublic: swag.BoolValue(params.Reason.IsPublic),
		},
	)
	if err != nil {
		if err2, ok := errors.Cause(err).(*core.Error); ok {
			return adminops.NewAdminCreateLotCanceledReasonBadRequest().WithPayload(models.NewError(err2))
		}
		return adminops.NewAdminCreateLotCanceledReasonInternalServerError().WithPayload(models.NewInternalServerError(err))
	}
	return adminops.NewAdminCreateLotCanceledReasonOK().WithPayload(models.NewLotCanceledReason(result))
}

func (h *Handler) adminUpdateLotCanceledReason(params adminops.AdminUpdateLotCanceledReasonParams, identity *authz.Identity) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	result, err := h.Admin.UpdateLotCanceledReason(ctx,
		identity.User,
		core.LotCanceledReasonID(int(params.ID)),
		&admin.LotCanceledReasonInput{
			Why:      swag.StringValue(params.Reason.Why),
			IsPublic: swag.BoolValue(params.Reason.IsPublic),
		},
	)
	if err != nil {
		if err2, ok := errors.Cause(err).(*core.Error); ok {
			return adminops.NewAdminUpdateLotCanceledReasonBadRequest().WithPayload(models.NewError(err2))
		}
		return adminops.NewAdminUpdateLotCanceledReasonInternalServerError().WithPayload(models.NewInternalServerError(err))
	}
	return adminops.NewAdminUpdateLotCanceledReasonOK().WithPayload(models.NewLotCanceledReason(result))
}

func (h *Handler) adminDeleteLotCanceledReason(params adminops.AdminDeleteLotCanceledReasonParams, identity *authz.Identity) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	if err := h.Admin.DeleteLotCanceledReason(ctx, identity.User, core.LotCanceledReasonID(int(params.ID))); err != nil {
		if err2, ok := errors.Cause(err).(*core.Error); ok {
			return adminops.NewAdminDeleteLotCanceledReasonBadRequest().WithPayload(models.NewError(err2))
		}
		return adminops.NewAdminDeleteLotCanceledReasonInternalServerError().WithPayload(models.NewInternalServerError(err))
	}
	return adminops.NewAdminDeleteLotCanceledReasonNoContent()
}

func (h *Handler) adminDeclineLot(params adminops.AdminDeclineLotParams, identity *authz.Identity) middleware.Responder {
	ctx := params.HTTPRequest.Context()
	lotID := core.LotID(int(params.ID))

	if err := h.Admin.DeclineLot(ctx, identity.GetUser(), lotID, params.Reason); err != nil {
		if err2, ok := errors.Cause(err).(*core.Error); ok {
			return adminops.NewAdminDeclineLotBadRequest().WithPayload(models.NewError(err2))
		}
		return adminops.NewAdminDeclineLotInternalServerError().WithPayload(models.NewInternalServerError(err))
	}
	return adminops.NewAdminDeclineLotNoContent()
}

func (h *Handler) adminGetLot(params adminops.AdminGetLotParams, identity *authz.Identity) middleware.Responder {
	ctx := params.HTTPRequest.Context()
	id := core.LotID(int(params.ID))

	result, err := h.Admin.GetLot(ctx, identity.GetUser(), id)
	if err != nil {
		if err2, ok := errors.Cause(err).(*core.Error); ok {
			return adminops.NewAdminGetLotBadRequest().WithPayload(models.NewError(err2))
		}
		return adminops.NewAdminGetLotInternalServerError().WithPayload(models.NewInternalServerError(err))
	}

	return adminops.NewAdminGetLotOK().WithPayload(models.NewAdminFullLot(h.Storage, result))
}
