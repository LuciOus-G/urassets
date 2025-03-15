package request

type UserRegisterRequest struct {
	PhoneNumber string `json:"phone_number" validate:"required,min=10,max=13"`
	Email       string `json:"email" validate:"required"`
	Password    string `json:"password" validate:"required"`
}

type UserLoginRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}
