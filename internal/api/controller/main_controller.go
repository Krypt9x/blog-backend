package controller

import "github.com/gofiber/fiber/v2"

type MainController interface {
	Create(ctx *fiber.Ctx) error
	GetAll(ctx *fiber.Ctx) error
	GetByUsername(ctx *fiber.Ctx) error
	GetById(ctx *fiber.Ctx) error
	UpdateById(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}
