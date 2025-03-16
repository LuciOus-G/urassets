package service

import (
	"fmt"
	"github.com/ericlagergren/decimal"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
	"github.com/lucious/urassets/UrAssetsCore/core/models"
	"github.com/lucious/urassets/UrAssetsCore/core/request"
	"github.com/lucious/urassets/UrAssetsCore/core/response"
	"github.com/lucious/urassets/Utilities"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/types"
	"time"
)

func (srv Services) UserRegister(c *fiber.Ctx, MasterUser *request.UserRegisterRequest) (err error) {
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

	// create all important relationship initialization
	newUserSteps := models.UserStep{
		UserID: NewUser.ID,
	}
	err = newUserSteps.Insert(srv.Ctx, srv.DB, boil.Infer())
	if err != nil {
		srvResponse.Log.Info(err)
		srvResponse.Err = err
		return srvResponse.BadRequest()
	}

	totalWealth := models.TotalWealth{
		UserID:  NewUser.ID,
		Balance: types.NewDecimal(decimal.New(0, 2)),
	}
	err = totalWealth.Insert(srv.Ctx, srv.DB, boil.Infer())
	if err != nil {
		srvResponse.Log.Info(err)
		srvResponse.Err = err
		return srvResponse.BadRequest()
	}

	currenMonthWealth := models.CurrentMonthWealth{
		UsersID: NewUser.ID,
		Balance: types.NewDecimal(decimal.New(0, 2)),
		PeriodMonth: null.Time{
			Time:  time.Now(),
			Valid: true,
		},
	}
	err = currenMonthWealth.Insert(srv.Ctx, srv.DB, boil.Infer())
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

func (srv Services) UserLogin(c *fiber.Ctx, userRequest *request.UserLoginRequest) error {
	srvResponse := Utilities.NewResponse(Utilities.BaseResponse{Ctx: c})
	userResponse := response.UserResponse{}

	userDetail, err := models.Users(
		qm.Where("email = ?", userRequest.Email),
		qm.Load("UserStep"),
	).One(srv.Ctx, srv.DB)
	if err != nil {
		srvResponse.Err = fmt.Errorf("email not found")
		return srvResponse.NotFound()
	}

	isPasswordValid, err := Utilities.CheckPassword(userRequest.Password, userDetail.Password)
	if err != nil {
		srvResponse.Log.Info(err)
		srvResponse.Err = err
		return srvResponse.BadRequest()
	}

	if !isPasswordValid {
		srvResponse.Err = fmt.Errorf("password is invalid")
		srvResponse.Log.Info(err)
		return srvResponse.AccessDenied()
	}

	err = copier.Copy(&userResponse, &userDetail)
	if err != nil {
		srvResponse.Log.Info(err)
		srvResponse.Err = err
		return srvResponse.BadRequest()
	}

	token, err := Utilities.GenerateJWT(userResponse.ID)
	if err != nil {
		srvResponse.Log.Info(err)
		srvResponse.Err = err
		return srvResponse.BadRequest()
	}

	// bind the relationship
	userResponse.UserJourney = userDetail.R.UserStep
	userResponse.Token = token

	srvResponse.Data = userResponse
	srvResponse.Status = fiber.StatusOK

	return srvResponse.OK()
}

func (srv Services) UserDetail(c *fiber.Ctx) error {
	srvResponse := Utilities.NewResponse(Utilities.BaseResponse{Ctx: c})
	userResponse := response.UserDetailResponse{}

	user, err := models.Users(
		qm.Where("id = ?", c.Params("id")),
		qm.Load("UserIncomeCategories"),
		qm.Load("UserExpensesCategories"),
	).One(srv.Ctx, srv.DB)
	if err != nil {
		srvResponse.Err = fmt.Errorf("user not found")
		return srvResponse.NotFound()
	}

	err = copier.Copy(&userResponse, &user)
	if err != nil {
		srvResponse.Log.Info(err)
		srvResponse.Err = err
		return srvResponse.BadRequest()
	}

	// bind the relationship into response
	userResponse.IncomeCategories = user.R.UserIncomeCategories
	userResponse.ExpensesCategories = user.R.UserExpensesCategories

	srvResponse.Data = userResponse
	srvResponse.Status = fiber.StatusOK

	return srvResponse.OK()

}
