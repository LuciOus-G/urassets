package request

type UserCategoriesIncomeRequests struct {
	UserId string `validate:"required"`
	Name   string `validate:"required"`
}
