package response

import "github.com/volatiletech/null/v8"

type UserResponse struct {
	ID         string
	FullName   string
	Email      string
	IsActive   null.Bool
	CreatedAt  null.Time
	UpdatedAt  null.Time
	LastUpdate null.Time
	Token      string
}
