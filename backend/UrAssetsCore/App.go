package UrAssetsCore

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/lucious/urassets/UrAssetsCore/controllers"
	"github.com/lucious/urassets/UrAssetsCore/core/service"
	"github.com/lucious/urassets/Utilities"
)

func UrAssetsCore(Host string, Port string, DB *sql.DB) {
	app := fiber.New(fiber.Config{
		AppName:     "UrAssets",
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	ctx := context.Background()
	Handler := controllers.UrAssetsHandler{
		DB:       DB,
		Ctx:      ctx,
		Response: Utilities.NewResponse(Utilities.BaseResponse{}),
		Service: service.Interfaces(service.Services{
			Ctx: ctx,
			DB:  DB,
		}),
	}

	Route := app.Group("")

	ApiV1 := Route.Group("/api/ua/v1")

	userAPI := ApiV1.Group("/user")
	userAPI.Get("/:id", Handler.UserDetail)
	userAPI.Post("/login", Handler.UserLogin)
	userAPI.Post("/register", Handler.UserRegister)

	userCategoryAPI := ApiV1.Group("/user-category")
	userCategoryAPI.Post("/add/income", Handler.PostUserCategoriesIncome)
	userCategoryAPI.Post("/add/expenses", Handler.PostUserCategoriesExpenses)

	banksAPI := ApiV1.Group("/banks")
	banksAPI.Get("/", Handler.GetBankList)
	banksAPI.Post("/add/user-banks", Handler.PostUserBanks)

	Route.Get("/swagger/*", swagger.HandlerDefault)

	err := app.Listen(fmt.Sprintf("%s:%s", Host, Port))
	if err != nil {
		return
	}
}
