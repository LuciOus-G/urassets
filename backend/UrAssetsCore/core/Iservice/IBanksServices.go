package Iservice

import (
	"github.com/gofiber/fiber/v2"
	request2 "github.com/lucious/urassets/UrAssetsCore/core/request"
)

type IBanksServices interface {
	GetBanksList(ctx *fiber.Ctx) error
	PostUserBanks(c *fiber.Ctx, request *request2.PostUserBanksRequests) error
}
