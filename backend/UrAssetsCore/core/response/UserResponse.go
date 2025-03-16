package response

import (
	"github.com/lucious/urassets/UrAssetsCore/core/models"
	"github.com/volatiletech/null/v8"
)

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

type UserDetailResponse struct {
	*models.User
	Password           string `copier:"-"`
	IncomeCategories   []*models.UserIncomeCategory
	ExpensesCategories []*models.UserExpensesCategory
}
