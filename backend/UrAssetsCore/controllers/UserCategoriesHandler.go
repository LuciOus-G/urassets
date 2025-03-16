package controllers

import (
	"github.com/gofiber/fiber/v2"
	request2 "github.com/lucious/urassets/UrAssetsCore/core/request"
	"github.com/lucious/urassets/Utilities"
)

func (h *UrAssetsHandler) PostUserCategories(ctx *fiber.Ctx) error {
	h.Response.Ctx = ctx

	request := new(request2.UserCategoriesIncomeRequests)
	if err := Utilities.BodyParse(ctx, request, h.Response); err != nil {
		h.Response.Err = err
		return h.Response.BadRequest()
	}

	return h.Service.PostIncomeCategories(ctx, request)
}
