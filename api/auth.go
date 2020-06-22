package api

import (
	"github.com/bots-house/birzzha/api/authz"
	authops "github.com/bots-house/birzzha/api/gen/restapi/operations/auth"
	"github.com/bots-house/birzzha/api/models"
	"github.com/bots-house/birzzha/core"
	"github.com/bots-house/birzzha/service/auth"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/pkg/errors"
)

func (h *Handler) createToken(params authops.CreateTokenParams) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	info := &auth.TelegramWidgetInfo{
		ID:        int(swag.Int64Value(params.Payload.ID)),
		FirstName: swag.StringValue(params.Payload.FirstName),
		LastName:  params.Payload.LastName,
		Username:  params.Payload.Username,
		PhotoURL:  params.Payload.PhotoURL,
		AuthDate:  swag.Int64Value(params.Payload.AuthDate),
		Hash:      swag.StringValue(params.Payload.Hash),
	}

	token, err := h.Auth.CreateToken(ctx, info)
	if err != nil {
		if err2, ok := errors.Cause(err).(*core.Error); ok {
			return authops.NewCreateTokenBadRequest().WithPayload(models.NewError(err2))
		}
		return authops.NewCreateTokenInternalServerError().WithPayload(models.NewInternalServerError(err))
	}

	return authops.NewCreateTokenCreated().WithPayload(models.NewToken(token))
}

func (h *Handler) getUser(params authops.GetUserParams, identity *authz.Identity) middleware.Responder {
	return authops.NewGetUserOK().WithPayload(models.NewUser(identity.User))
}

func (h *Handler) loginViaBot(params authops.LoginViaBotParams) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	url, err := h.Auth.LoginViaBot(ctx, &auth.LoginViaBotInfo{
		CallbackURL: params.CallbackURL,
	})
	if err != nil {
		return authops.NewLoginViaBotInternalServerError().WithPayload(models.NewInternalServerError(err))
	}

	return authops.NewLoginViaBotFound().WithLocation(url)
}
