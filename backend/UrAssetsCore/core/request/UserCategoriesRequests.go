package request

type UserCategoriesRequests struct {
	UserId string `validate:"required"`
	Name   string `validate:"required"`
}
