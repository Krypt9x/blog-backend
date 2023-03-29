package controller

import (
	services "github.com/Krypt9x/blog-backend/internal/api/service"
	"github.com/Krypt9x/blog-backend/internal/model"
	"github.com/Krypt9x/blog-backend/internal/model/domain"
	"github.com/Krypt9x/blog-backend/pkg/helper"
	"github.com/gofiber/fiber/v2"
)

type UserControllerImpl struct {
	Service services.UserService
}

func NewUserController(service services.UserService) UserController {
	return &UserControllerImpl{
		Service: service,
	}
}

func (contoller *UserControllerImpl) Register(ctx *fiber.Ctx) error {
	reqData := new(domain.User)
	if err := ctx.BodyParser(reqData); err != nil {
		helper.PanicIfError(err)
	}

	err := contoller.Service.Register(ctx.Context(), *reqData)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(&model.Response{
			Code:   500,
			Status: "server error",
			Data:   "failed register",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&model.Response{
		Code:   200,
		Status: "OK",
		Data:   "berhasil register",
	})

}

func (controller *UserControllerImpl) Login(ctx *fiber.Ctx) error {
	reqData := new(domain.User)
	if err := ctx.BodyParser(reqData); err != nil {
		helper.PanicIfError(err)
	}

	err := controller.Service.Login(ctx.Context(), *reqData)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(&model.Response{
			Code:   401,
			Status: "user not found",
			Data:   "username or email or password are wrong",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&model.Response{
		Code:   200,
		Status: "OK",
		Data:   "welcome back",
	})
}
