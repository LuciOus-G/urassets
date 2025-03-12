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
	request := new(request2.MasterUserRequest)
	if err := Utilities.BodyParse(ctx, request, h.Response); err != nil {
		fmt.Println(err)
		h.Response.Err = err
		return h.Response.BadRequest()
	}

	return h.Service.UserRegister(ctx, request)
}

func (h *UrAssetsHandler) UserDetail(ctx *fiber.Ctx) error {

	return h.Service.UserDetails(ctx)
}
