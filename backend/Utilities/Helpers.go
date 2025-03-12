package Utilities

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func BodyParse(ctx *fiber.Ctx, body interface{}, response *BaseResponse) error {
	if err := ctx.BodyParser(&body); err != nil {
		fmt.Println(err)
		response.Err = err
		return err
	}

	validate := validator.New()
	err := validate.Struct(body)
	if err != nil {
		fmt.Println(err)
		response.Err = err
		return err
	}

	return nil
}
