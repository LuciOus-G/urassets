package service

import (
	"context"
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
	"github.com/lucious/urassets/UrAssetsCore/core/Iservice"
	"github.com/lucious/urassets/UrAssetsCore/core/models"
	"github.com/lucious/urassets/UrAssetsCore/core/request"
	"github.com/lucious/urassets/UrAssetsCore/core/response"
	"github.com/lucious/urassets/Utilities"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type UserJourneyService struct {
	Ctx context.Context
	DB  *sql.DB
}

func NewUserJourneyService(init ...UserJourneyService) Iservice.IUserJourneyService {
	return Iservice.IUserJourneyService(init[0])
}

func (srv UserJourneyService) UserRegister(c *fiber.Ctx, MasterUser *request.MasterUserRequest) (err error) {
	srvResponse := Utilities.NewResponse(Utilities.BaseResponse{Ctx: c})
	userResponse := response.UserResponse{}

	hashPassword, err := Utilities.HashPassword(MasterUser.Password)
	if err != nil {
		srvResponse.Log.Info(err)
		srvResponse.Err = err
		return srvResponse.BadRequest()
	}

	// user creation
	NewUser := models.User{
		FullName: MasterUser.PhoneNumber,
		Email:    MasterUser.Email,
		Password: hashPassword,
	}
	err = NewUser.Insert(srv.Ctx, srv.DB, boil.Infer())
	if err != nil {
		srvResponse.Log.Info(err)
		srvResponse.Err = err
		return srvResponse.BadRequest()
	}

	err = copier.Copy(&userResponse, &NewUser)
	if err != nil {
		srvResponse.Log.Info(err)
		srvResponse.Err = err
		return srvResponse.BadRequest()
	}

	srvResponse.Data = userResponse
	srvResponse.Status = fiber.StatusCreated

	return srvResponse.OK()
}

func (srv UserJourneyService) UserDetails(c *fiber.Ctx) error {
	srvResponse := Utilities.NewResponse(Utilities.BaseResponse{Ctx: c})

	srvResponse.Data = "user"
	srvResponse.Status = 200

	return srvResponse.OK()
}
