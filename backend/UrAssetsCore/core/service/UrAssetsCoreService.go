package service

import (
	"context"
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"github.com/lucious/urassets/UrAssetsCore/core/Iservice"
	"github.com/lucious/urassets/UrAssetsCore/core/models"
	"github.com/lucious/urassets/UrAssetsCore/core/request"
	"github.com/lucious/urassets/Utilities"
	"github.com/volatiletech/null/v8"
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
	hashPin, err := Utilities.HashPassword(MasterUser.SecurePin)
	if err != nil {
		srvResponse.Log.Info(err)
		srvResponse.Err = err
		return srvResponse.BadRequest()
	}

	hashPassword, err := Utilities.HashPassword(MasterUser.Password)
	if err != nil {
		srvResponse.Log.Info(err)
		srvResponse.Err = err
		return srvResponse.BadRequest()
	}

	encryptBalance, err := Utilities.EncryptAES256Rest("test", "LmrxZ2zhQkV5uRln0xTWuheqcF+CexDAf0iChA2nwBgeaAdQSn5goBV2X+s4JqNH")
	if err != nil {
		srvResponse.Log.Info(err)
		srvResponse.Err = err
		return srvResponse.BadRequest()
	}

	// user creation
	NewUser := models.MasterUser{
		PhoneNumber: MasterUser.PhoneNumber,
		SecurePin:   hashPin,
		Password:    null.String{String: hashPassword, Valid: len(MasterUser.Password) > 0},
		Email:       null.String{String: encryptBalance, Valid: len(MasterUser.Email) > 0},
	}
	err = NewUser.Insert(srv.Ctx, srv.DB, boil.Infer())
	if err != nil {
		srvResponse.Log.Info(err)
		srvResponse.Err = err
		return srvResponse.BadRequest()
	}

	// user creation should be paired with wallet creation
	NewUserWallet := models.MasterWallet{
		MasterUserID: null.String{String: NewUser.UID, Valid: len(NewUser.UID) > 0},
		Balance:      encryptBalance,
	}
	err = NewUserWallet.Insert(srv.Ctx, srv.DB, boil.Infer())
	if err != nil {
		srvResponse.Log.Info(err)
		srvResponse.Err = err
		return srvResponse.BadRequest()
	}

	srvResponse.Data = NewUser
	srvResponse.Status = fiber.StatusCreated

	return srvResponse.OK()
}

func (srv UserJourneyService) UserDetails(c *fiber.Ctx) error {
	srvResponse := Utilities.NewResponse(Utilities.BaseResponse{Ctx: c})
	user, err := models.MasterWallets().All(srv.Ctx, srv.DB)

	if err != nil {
		srvResponse.Log.Info(err)
		srvResponse.Err = err
		return srvResponse.BadRequest()
	}

	srvResponse.Data = user
	srvResponse.Status = 200

	return srvResponse.OK()
}
