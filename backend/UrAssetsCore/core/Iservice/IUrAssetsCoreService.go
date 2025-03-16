package Iservice

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lucious/urassets/UrAssetsCore/core/request"
)

type IUserJourneyService interface {
	UserLogin(c *fiber.Ctx, userRequest *request.UserLoginRequest) error
	UserRegister(c *fiber.Ctx, MasterUser *request.UserRegisterRequest) error
	UserDetail(c *fiber.Ctx) error
}
