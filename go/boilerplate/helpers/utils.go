package helpers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func CheckError(err error) error {
	if err != nil {
		return err
	}
	return nil
}

func CheckErrorBodyParse(ctx *fiber.Ctx, err error) error {
	if err != nil {
		ctx.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request failed"})
		return err
	}
	return nil
}

func CheckErrorBadRequest(ctx *fiber.Ctx, err error) error {
	if err != nil {
		ctx.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not create data"})
		return err
	}
	return nil
}
