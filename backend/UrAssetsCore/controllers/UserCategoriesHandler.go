package controllers

import (
	"github.com/gofiber/fiber/v2"
	request2 "github.com/lucious/urassets/UrAssetsCore/core/request"
	"github.com/lucious/urassets/Utilities"
)

func (h *UrAssetsHandler) PostUserCategoriesIncome(ctx *fiber.Ctx) error {
	h.Response.Ctx = ctx

	request := new(request2.UserCategoriesRequests)
	if err := Utilities.BodyParse(ctx, request, h.Response); err != nil {
		h.Response.Err = err
		return h.Response.BadRequest()
	}

	return h.Service.PostUserCategories(ctx, request, 1)
}

func (h *UrAssetsHandler) PostUserCategoriesExpenses(ctx *fiber.Ctx) error {
	h.Response.Ctx = ctx

	request := new(request2.UserCategoriesRequests)
	if err := Utilities.BodyParse(ctx, request, h.Response); err != nil {
		h.Response.Err = err
		return h.Response.BadRequest()
	}

	return h.Service.PostUserCategories(ctx, request, 2)
}
