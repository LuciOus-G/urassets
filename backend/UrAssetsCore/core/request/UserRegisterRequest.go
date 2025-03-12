package request

type MasterUserRequest struct {
	PhoneNumber string `json:"phone_number" validate:"required,min=10,max=13"`
	SecurePin   string `json:"secure_pin" validate:"required,min=6,max=6"`
	Email       string `json:"email"`
	Password    string `json:"password"`
}
