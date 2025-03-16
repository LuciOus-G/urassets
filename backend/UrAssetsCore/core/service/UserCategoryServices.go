package service

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/lucious/urassets/UrAssetsCore/core/models"
	request2 "github.com/lucious/urassets/UrAssetsCore/core/request"
	"github.com/lucious/urassets/Utilities"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func (srv Services) PostUserCategories(c *fiber.Ctx, request *request2.UserCategoriesRequests, recordType int) error {
	srvResponse := Utilities.NewResponse(Utilities.BaseResponse{Ctx: c})

	if recordType == 1 {
		userCategory := models.UserIncomeCategory{
			UserID: request.UserId,
			Name:   request.Name,
		}

		err := userCategory.Insert(srv.Ctx, srv.DB, boil.Infer())
		if err != nil {
			srvResponse.Err = err
			return srvResponse.BadRequest()
		}
	} else if recordType == 2 {
		userCategory := models.UserExpensesCategory{
			UserID: request.UserId,
			Name:   request.Name,
		}

		err := userCategory.Insert(srv.Ctx, srv.DB, boil.Infer())
		if err != nil {
			srvResponse.Err = err
			return srvResponse.BadRequest()
		}
	} else {
		srvResponse.Err = fmt.Errorf("invalid recordType")
		return srvResponse.BadRequest()
	}

	srvResponse.Data = "Success Create Income Category"
	srvResponse.Status = fiber.StatusOK

	return srvResponse.OK()

}
