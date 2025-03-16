package controllers

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"github.com/lucious/urassets/UrAssetsCore/core/Iservice"
	request2 "github.com/lucious/urassets/UrAssetsCore/core/request"
	"github.com/lucious/urassets/Utilities"
)

type UrAssetsHandler struct {
	Ctx      context.Context
	Service  Iservice.IUserJourneyService
	DB       *sql.DB
	Response *Utilities.BaseResponse
}

func (h *UrAssetsHandler) UserRegister(ctx *fiber.Ctx) error {
	// initialize handler
	h.Response.Ctx = ctx // REQUIRED FOR EVERY HANDLER
	// end initialize handler
	request := new(request2.UserRegisterRequest)
	if err := Utilities.BodyParse(ctx, request, h.Response); err != nil {
		h.Response.Err = err
		return h.Response.BadRequest()
	}

	return h.Service.UserRegister(ctx, request)
}

func (h *UrAssetsHandler) UserLogin(ctx *fiber.Ctx) error {
	h.Response.Ctx = ctx

	request := new(request2.UserLoginRequest)
	if err := Utilities.BodyParse(ctx, request, h.Response); err != nil {
		h.Response.Err = err
		return h.Response.BadRequest()
	}

	return h.Service.UserLogin(ctx, request)
}

func (h *UrAssetsHandler) UserDetail(ctx *fiber.Ctx) error {
	h.Response.Ctx = ctx

	if ctx.Params("id") == "" {
		h.Response.Err = fmt.Errorf("id is required")
		return h.Response.BadRequest()
	}

	return h.Service.UserDetail(ctx)
}
