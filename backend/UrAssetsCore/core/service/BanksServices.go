package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lucious/urassets/UrAssetsCore/core/models"
	request2 "github.com/lucious/urassets/UrAssetsCore/core/request"
	"github.com/lucious/urassets/Utilities"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func (srv Services) GetBanksList(c *fiber.Ctx) error {
	srvResponse := Utilities.NewResponse(Utilities.BaseResponse{Ctx: c})

	banks, err := models.Banks().All(srv.Ctx, srv.DB)
	if err != nil {
		srvResponse.Err = err
		return srvResponse.BadRequest()
	}

	srvResponse.Data = banks
	srvResponse.Status = fiber.StatusOK

	return srvResponse.OK()

}

func (srv Services) PostUserBanks(c *fiber.Ctx, request *request2.PostUserBanksRequests) error {
	srvResponse := Utilities.NewResponse(Utilities.BaseResponse{Ctx: c})

	userBanks := models.UserBank{
		BankID: request.BankID,
		UserID: request.UserID,
	}

	err := userBanks.Insert(srv.Ctx, srv.DB, boil.Infer())
	if err != nil {
		srvResponse.Err = err
		return srvResponse.BadRequest()
	}

	srvResponse.Data = "Success Add Bank"
	srvResponse.Status = fiber.StatusOK

	return srvResponse.OK()

}
