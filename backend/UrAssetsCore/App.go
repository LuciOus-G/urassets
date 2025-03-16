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

	UserAPI := ApiV1.Group("/user")
	UserAPI.Get("/:id", Handler.UserDetail)
	UserAPI.Post("/login", Handler.UserLogin)
	UserAPI.Post("/register", Handler.UserRegister)

	UserCategoryAPI := ApiV1.Group("/user-category")
	UserCategoryAPI.Post("/add/income", Handler.PostUserCategories)

	Route.Get("/swagger/*", swagger.HandlerDefault)

	err := app.Listen(fmt.Sprintf("%s:%s", Host, Port))
	if err != nil {
		return
	}
}
