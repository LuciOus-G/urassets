package UrAssetsCore

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/lucious/urassets/UrAssetsCore/controllers"
	"github.com/lucious/urassets/UrAssetsCore/core/service"
	"github.com/lucious/urassets/Utilities"
)

func RunUserServiceApp(Host string, Port string, DB *sql.DB) {
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

	ApiV1 := app.Group("/api/ua/v1")
	ApiV1.Get("/", Handler.UserDetail)
	ApiV1.Post("/", Handler.UserRegister)

	app.Get("/swagger/*", swagger.HandlerDefault)

	err := app.Listen(fmt.Sprintf("%s:%s", Host, Port))
	if err != nil {
		return
	}
}
