package Iservice

import (
	"github.com/gofiber/fiber/v2"
	request2 "github.com/lucious/urassets/UrAssetsCore/core/request"
)

type IUserCategoryServices interface {
	PostUserCategories(c *fiber.Ctx, request *request2.UserCategoriesRequests, recordType int) error
}
