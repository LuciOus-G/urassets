package request

type UserRegisterRequest struct {
	PhoneNumber string `validate:"required,min=10,max=13"`
	Email       string `validate:"required"`
	Password    string `validate:"required"`
}

type UserLoginRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}
