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
		Service: service.NewUserJourneyService(service.UserJourneyService{
			Ctx: ctx,
			DB:  DB,
		}),
	}

	Route := app.Group("")

	ApiV1 := Route.Group("/api/ua/v1")
	ApiV1.Get("/user/:id", Handler.UserDetail)
	ApiV1.Post("/login", Handler.UserLogin)
	ApiV1.Post("/", Handler.UserRegister)

	Route.Get("/swagger/*", swagger.HandlerDefault)

	err := app.Listen(fmt.Sprintf("%s:%s", Host, Port))
	if err != nil {
		return
	}
}
