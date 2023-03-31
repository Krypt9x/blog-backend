package helper

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

type HttpResponseData struct {
	Code   int
	Status string
	Data   any
}

func SendJSONResponse(ctx *fiber.Ctx, res *HttpResponseData) error {
	if res == nil || ctx == nil {
		return errors.New("response data must be not null")
	}
	return ctx.Status(res.Code).JSON(res)
}
