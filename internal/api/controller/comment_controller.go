package controller

import "github.com/gofiber/fiber/v2"

type CommentController interface {
	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}
