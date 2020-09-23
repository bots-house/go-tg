package api

import (
	"time"

	"github.com/bots-house/birzzha/api/authz"
	adminops "github.com/bots-house/birzzha/api/gen/restapi/operations/admin"
	"github.com/bots-house/birzzha/api/models"
	"github.com/bots-house/birzzha/core"
	"github.com/bots-house/birzzha/service/admin"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/pkg/errors"
	"github.com/volatiletech/null/v8"
)

func (h *Handler) adminDeleteReview(params adminops.AdminDeleteReviewParams, identity *authz.Identity) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	if err := h.Admin.DeleteReview(ctx, identity.User, core.ReviewID(int(params.ID))); err != nil {
		if err2, ok := errors.Cause(err).(*core.Error); ok {
			return adminops.NewAdminDeleteReviewBadRequest().WithPayload(models.NewError(err2))
		}
		return adminops.NewAdminDeleteReviewInternalServerError().WithPayload(models.NewInternalServerError(ctx, err))
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
		return adminops.NewAdminUpdateReviewInternalServerError().WithPayload(models.NewInternalServerError(ctx, err))
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
		return adminops.NewAdminGetReviewsInternalServerError().WithPayload(models.NewInternalServerError(ctx, err))
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
		return adminops.NewAdminGetUsersInternalServerError().WithPayload(models.NewInternalServerError(ctx, err))
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
		return adminops.NewToggleUserAdminInternalServerError().WithPayload(models.NewInternalServerError(ctx, err))
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
		return adminops.NewAdminGetLotStatusesInternalServerError().WithPayload(models.NewInternalServerError(ctx, err))
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
		return adminops.NewAdminGetLotsInternalServerError().WithPayload(models.NewInternalServerError(ctx, err))
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
		return adminops.NewAdminGetSettingsInternalServerError().WithPayload(models.NewInternalServerError(ctx, err))
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
		return adminops.NewAdminUpdateSettingsLandingInternalServerError().WithPayload(models.NewInternalServerError(ctx, err))
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
		return adminops.NewAdminUpdateReviewInternalServerError().WithPayload(models.NewInternalServerError(ctx, err))
	}

	return adminops.NewAdminUpdateSettingsPricesOK().WithPayload(models.NewSettingsPrices(result))
}

func (h *Handler) adminUpdateSettingsGarant(params adminops.AdminUpdateSettingsGarantParams, identity *authz.Identity) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	result, err := h.Admin.UpdateSettingsGarant(ctx, identity.User, &admin.SettingsGarantInput{
		Name:                           swag.StringValue(params.Garant.Name),
		Username:                       swag.StringValue(params.Garant.Username),
		ReviewsChannel:                 swag.StringValue(params.Garant.ReviewsChannel),
		Avatar:                         params.Garant.AvatarURL,
		PercentageDealOfDiscountPeriod: params.Garant.PercentageDealOfDiscountPeriod,
		PercentageDeal:                 swag.Float64Value(params.Garant.PercentageDeal),
	})
	if err != nil {
		if err2, ok := errors.Cause(err).(*core.Error); ok {
			return adminops.NewAdminUpdateSettingsGarantBadRequest().WithPayload(models.NewError(err2))
		}
		return adminops.NewAdminUpdateSettingsGarantInternalServerError().WithPayload(models.NewInternalServerError(ctx, err))
	}
	return adminops.NewAdminUpdateSettingsGarantOK().WithPayload(models.NewSettingsGarant(result))
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
		return adminops.NewAdminUpdateReviewInternalServerError().WithPayload(models.NewInternalServerError(ctx, err))
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
		return adminops.NewAdminCreateTopicInternalServerError().WithPayload(models.NewInternalServerError(ctx, err))
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
		return adminops.NewAdminUpdateTopicInternalServerError().WithPayload(models.NewInternalServerError(ctx, err))
	}
	return adminops.NewAdminUpdateTopicOK().WithPayload(models.NewAdminFullTopic(result))
}

func (h *Handler) adminDeleteTopic(params adminops.AdminDeleteTopicParams, identity *authz.Identity) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	if err := h.Admin.DeleteTopic(ctx, identity.User, core.TopicID(int(params.ID))); err != nil {
		if err2, ok := errors.Cause(err).(*core.Error); ok {
			return adminops.NewAdminDeleteTopicBadRequest().WithPayload(models.NewError(err2))
		}
		return adminops.NewAdminDeleteTopicInternalServerError().WithPayload(models.NewInternalServerError(ctx, err))
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
		return adminops.NewAdminCreateLotCanceledReasonInternalServerError().WithPayload(models.NewInternalServerError(ctx, err))
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
		return adminops.NewAdminUpdateLotCanceledReasonInternalServerError().WithPayload(models.NewInternalServerError(ctx, err))
	}
	return adminops.NewAdminUpdateLotCanceledReasonOK().WithPayload(models.NewLotCanceledReason(result))
}

func (h *Handler) adminDeleteLotCanceledReason(params adminops.AdminDeleteLotCanceledReasonParams, identity *authz.Identity) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	if err := h.Admin.DeleteLotCanceledReason(ctx, identity.User, core.LotCanceledReasonID(int(params.ID))); err != nil {
		if err2, ok := errors.Cause(err).(*core.Error); ok {
			return adminops.NewAdminDeleteLotCanceledReasonBadRequest().WithPayload(models.NewError(err2))
		}
		return adminops.NewAdminDeleteLotCanceledReasonInternalServerError().WithPayload(models.NewInternalServerError(ctx, err))
	}
	return adminops.NewAdminDeleteLotCanceledReasonNoContent()
}

func (h *Handler) adminGetPostText(params adminops.AdminGetPostTextParams, identity *authz.Identity) middleware.Responder {
	ctx := params.HTTPRequest.Context()
	id := core.LotID(int(params.ID))

	result, err := h.Admin.GetPostText(ctx, identity.GetUser(), id)
	if err != nil {
		if err2, ok := errors.Cause(err).(*core.Error); ok {
			return adminops.NewAdminGetPostTextBadRequest().WithPayload(models.NewError(err2))
		}
		return adminops.NewAdminGetLotStatusesInternalServerError().WithPayload(models.NewInternalServerError(ctx, err))
	}
	return adminops.NewAdminGetPostTextOK().WithPayload(models.NewPostText(result))
}

func (h *Handler) adminSendPostPreview(params adminops.AdminSendPostPreviewParams, identity *authz.Identity) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	if err := h.Admin.SendPostPreview(ctx, identity.GetUser(), swag.StringValue(params.Post.Text), core.LotID(params.Post.LotID)); err != nil {
		if err2, ok := errors.Cause(err).(*core.Error); ok {
			return adminops.NewAdminSendPostPreviewBadRequest().WithPayload(models.NewError(err2))
		}
		return adminops.NewAdminSendPostPreviewInternalServerError().WithPayload(models.NewInternalServerError(ctx, err))
	}
	return adminops.NewAdminSendPostPreviewOK()
}

func (h *Handler) adminUpdateLot(params adminops.AdminUpdateLotParams, identity *authz.Identity) middleware.Responder {
	ctx := params.HTTPRequest.Context()
	id := core.LotID(int(params.ID))

	result, err := h.Admin.UpdateLot(ctx, identity.GetUser(), id, admin.InputAdminLot{
		Comment: swag.StringValue(params.Lot.Comment),
		Price:   int(swag.Int64Value(params.Lot.Price)),
		Extra:   models.ToLotExtraResourceSlice(params.Lot.Extra),
		Topics:  models.ToTopicIDs(params.Lot.Topics),
		Income:  int(swag.Int64Value(params.Lot.MonthlyIncome)),
	})
	if err != nil {
		if err2, ok := errors.Cause(err).(*core.Error); ok {
			return adminops.NewAdminUpdateLotBadRequest().WithPayload(models.NewError(err2))
		}
		return adminops.NewAdminUpdateLotInternalServerError().WithPayload(models.NewInternalServerError(ctx, err))
	}
	return adminops.NewAdminUpdateLotOK().WithPayload(models.NewAdminFullLot(h.Storage, result))
}

func (h *Handler) adminDeclineLot(params adminops.AdminDeclineLotParams, identity *authz.Identity) middleware.Responder {
	ctx := params.HTTPRequest.Context()
	lotID := core.LotID(int(params.ID))

	if err := h.Admin.DeclineLot(ctx, identity.GetUser(), lotID, swag.StringValue(params.Reason)); err != nil {
		if err2, ok := errors.Cause(err).(*core.Error); ok {
			return adminops.NewAdminDeclineLotBadRequest().WithPayload(models.NewError(err2))
		}
		return adminops.NewAdminDeclineLotInternalServerError().WithPayload(models.NewInternalServerError(ctx, err))
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
		return adminops.NewAdminGetLotInternalServerError().WithPayload(models.NewInternalServerError(ctx, err))
	}

	return adminops.NewAdminGetLotOK().WithPayload(models.NewAdminFullLot(h.Storage, result))
}

func (h *Handler) adminCreatePost(params adminops.AdminCreatePostParams, identity *authz.Identity) middleware.Responder {
	ctx := params.HTTPRequest.Context()
	in := &admin.PostInput{
		Text:                  swag.StringValue(params.Post.Text),
		DisableWebPagePreview: swag.BoolValue(params.Post.DisableWebPagePreview),
		LotID:                 core.LotID(int(params.Post.LotID)),
		LotLinkButton:         params.Post.LotLinkButton,
	}
	if params.Post.ScheduledAt != 0 {
		in.ScheduledAt = time.Unix(params.Post.ScheduledAt, 0)
	} else {
		in.ScheduledAt = time.Now()
	}
	if params.Post.Title != "" {
		in.Title = null.StringFrom(params.Post.Title)
	}

	result, err := h.Admin.CreatePost(ctx, identity.GetUser(), in)
	if err != nil {
		if err2, ok := errors.Cause(err).(*core.Error); ok {
			return adminops.NewAdminCreatePostBadRequest().WithPayload(models.NewError(err2))
		}
		return adminops.NewAdminCreatePostInternalServerError().WithPayload(models.NewInternalServerError(ctx, err))
	}
	return adminops.NewAdminCreatePostCreated().WithPayload(models.NewPost(result))
}

func (h *Handler) adminUpdatePost(params adminops.AdminUpdatePostParams, identity *authz.Identity) middleware.Responder {
	ctx := params.HTTPRequest.Context()
	in := &admin.PostInput{
		LotID:                 core.LotID(int(params.Post.LotID)),
		Text:                  swag.StringValue(params.Post.Text),
		DisableWebPagePreview: swag.BoolValue(params.Post.DisableWebPagePreview),
		LotLinkButton:         params.Post.LotLinkButton,
	}
	if params.Post.ScheduledAt != 0 {
		in.ScheduledAt = time.Unix(params.Post.ScheduledAt, 0)
	} else {
		in.ScheduledAt = time.Now()
	}
	if params.Post.Title != "" {
		in.Title = null.StringFrom(params.Post.Title)
	}

	result, err := h.Admin.UpdatePost(ctx, identity.GetUser(), core.PostID(int(params.ID)), in)
	if err != nil {
		if err2, ok := errors.Cause(err).(*core.Error); ok {
			return adminops.NewAdminUpdatePostBadRequest().WithPayload(models.NewError(err2))
		}
		return adminops.NewAdminUpdatePostInternalServerError().WithPayload(models.NewInternalServerError(ctx, err))
	}
	return adminops.NewAdminUpdatePostOK().WithPayload(models.NewPostItem(h.Storage, result))
}

func (h *Handler) adminDeletePost(params adminops.AdminDeletePostParams, identity *authz.Identity) middleware.Responder {
	ctx := params.HTTPRequest.Context()
	id := core.PostID(int(params.ID))

	if err := h.Admin.DeletePost(ctx, identity.GetUser(), id, swag.BoolValue(params.DeleteFromChannel)); err != nil {
		if err2, ok := errors.Cause(err).(*core.Error); ok {
			return adminops.NewAdminDeletePostBadRequest().WithPayload(models.NewError(err2))
		}
		return adminops.NewAdminDeletePostInternalServerError().WithPayload(models.NewInternalServerError(ctx, err))
	}
	return adminops.NewAdminDeletePostNoContent()
}

func (h *Handler) adminGetPosts(params adminops.AdminGetPostsParams, identity *authz.Identity) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	result, err := h.Admin.GetPosts(ctx, identity.GetUser(), int(swag.Int64Value(params.Limit)), int(swag.Int64Value(params.Offset)))
	if err != nil {
		if err2, ok := errors.Cause(err).(*core.Error); ok {
			return adminops.NewAdminGetPostsBadRequest().WithPayload(models.NewError(err2))
		}
		return adminops.NewAdminGetPostsInternalServerError().WithPayload(models.NewInternalServerError(ctx, err))
	}
	return adminops.NewAdminGetPostsOK().WithPayload(models.NewFullPost(h.Storage, result))
}

func (h *Handler) adminCancelLot(params adminops.AdminCancelLotParams, identity *authz.Identity) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	if err := h.Admin.CancelLot(ctx, identity.GetUser(), core.LotID(params.ID), core.LotCanceledReasonID(params.ReasonID)); err != nil {
		if err2, ok := errors.Cause(err).(*core.Error); ok {
			return adminops.NewAdminCancelLotBadRequest().WithPayload(models.NewError(err2))
		}
		return adminops.NewAdminCancelLotInternalServerError().WithPayload(models.NewInternalServerError(ctx, err))
	}
	return adminops.NewAdminCancelLotOK()
}

func (h *Handler) adminCreateCoupon(params adminops.AdminCreateCouponParams, identity *authz.Identity) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	coupon, err := h.Admin.CreateCoupon(ctx, identity.GetUser(), &admin.CouponInput{
		Code:                  swag.StringValue(params.Coupon.Code),
		Discount:              swag.Float64Value(params.Coupon.Discount),
		Purposes:              params.Coupon.Purposes,
		ExpireAt:              null.NewTime(time.Unix(params.Coupon.ExpireAt, 0), params.Coupon.ExpireAt != 0),
		MaxAppliesByUserLimit: null.NewInt(int(params.Coupon.MaxAppliesByUserLimit), params.Coupon.MaxAppliesByUserLimit != 0),
		MaxAppliesLimit:       null.NewInt(int(params.Coupon.MaxAppliesLimit), params.Coupon.MaxAppliesLimit != 0),
	})
	if err != nil {
		if err2, ok := errors.Cause(err).(*core.Error); ok {
			return adminops.NewAdminCreateCouponBadRequest().WithPayload(models.NewError(err2))
		}
		return adminops.NewAdminCreateCouponInternalServerError().WithPayload(models.NewInternalServerError(ctx, err))
	}
	return adminops.NewAdminCreateCouponCreated().WithPayload(models.NewCouponItem(coupon))
}

func (h *Handler) adminUpdateCoupon(params adminops.AdminUpdateCouponParams, identity *authz.Identity) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	coupon, err := h.Admin.UpdateCoupon(ctx, identity.GetUser(), core.CouponID(params.ID), &admin.CouponInput{
		Code:                  swag.StringValue(params.Coupon.Code),
		Discount:              swag.Float64Value(params.Coupon.Discount),
		Purposes:              params.Coupon.Purposes,
		ExpireAt:              null.NewTime(time.Unix(params.Coupon.ExpireAt, 0), params.Coupon.ExpireAt != 0),
		MaxAppliesByUserLimit: null.NewInt(int(params.Coupon.MaxAppliesByUserLimit), params.Coupon.MaxAppliesByUserLimit != 0),
		MaxAppliesLimit:       null.NewInt(int(params.Coupon.MaxAppliesLimit), params.Coupon.MaxAppliesLimit != 0),
	})
	if err != nil {
		if err2, ok := errors.Cause(err).(*core.Error); ok {
			return adminops.NewAdminUpdateCouponBadRequest().WithPayload(models.NewError(err2))
		}
		return adminops.NewAdminUpdateCouponInternalServerError().WithPayload(models.NewInternalServerError(ctx, err))
	}
	return adminops.NewAdminUpdateCouponOK().WithPayload(models.NewCouponItem(coupon))
}

func (h *Handler) adminDeleteCoupon(params adminops.AdminDeleteCouponParams, identity *authz.Identity) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	if err := h.Admin.DeleteCoupon(ctx, identity.GetUser(), core.CouponID(params.ID)); err != nil {
		if err2, ok := errors.Cause(err).(*core.Error); ok {
			return adminops.NewAdminDeleteCouponBadRequest().WithPayload(models.NewError(err2))
		}
		return adminops.NewAdminDeleteCouponInternalServerError().WithPayload(models.NewInternalServerError(ctx, err))
	}
	return adminops.NewAdminDeleteCouponNoContent()
}

func (h *Handler) adminGetCoupons(params adminops.AdminGetCouponsParams, identity *authz.Identity) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	coupons, err := h.Admin.GetCoupons(ctx, identity.GetUser(), int(swag.Int64Value(params.Limit)), int(swag.Int64Value(params.Offset)))
	if err != nil {
		if err2, ok := errors.Cause(err).(*core.Error); ok {
			return adminops.NewAdminGetCouponsBadRequest().WithPayload(models.NewError(err2))
		}
		return adminops.NewAdminGetCouponsInternalServerError().WithPayload(models.NewInternalServerError(ctx, err))
	}
	return adminops.NewAdminGetCouponsOK().WithPayload(models.NewCouponListItem(coupons))
}

func (h *Handler) adminRefreshLot(params adminops.AdminRefreshLotParams, identity *authz.Identity) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	result, err := h.Admin.RefreshLot(ctx, identity.GetUser(), core.LotID(params.ID), params.Identity)
	if err != nil {
		if err2, ok := errors.Cause(err).(*core.Error); ok {
			return adminops.NewAdminRefreshLotBadRequest().WithPayload(models.NewError(err2))
		}
		return adminops.NewAdminRefreshLotInternalServerError().WithPayload(models.NewInternalServerError(ctx, err))
	}
	return adminops.NewAdminRefreshLotOK().WithPayload(models.NewAdminFullLot(h.Storage, result))
}
