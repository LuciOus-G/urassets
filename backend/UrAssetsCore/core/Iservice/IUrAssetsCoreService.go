package Iservice

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lucious/urassets/UrAssetsCore/core/request"
)

type IUserJourneyService interface {
	UserDetails(c *fiber.Ctx) error
	UserRegister(c *fiber.Ctx, MasterUser *request.MasterUserRequest) error
}
