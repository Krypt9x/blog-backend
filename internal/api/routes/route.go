package routes

import (
	"encoding/json"

	"github.com/Krypt9x/blog-backend/internal/api/controller"
	"github.com/gofiber/fiber/v2"
)

type (
	Routes interface {
		InitRouter() *fiber.App
	}

	ControllerObj struct {
		UserController controller.UserController
		MainController controller.MainController
	}
)

func InitRouter(controller ControllerObj) *fiber.App {
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})
	app.Post("/blog/register", controller.UserController.Register)
	app.Post("/blog/login", controller.UserController.Login)
	app.Post("/blog/post", controller.MainController.Create)
	app.Get("/blog/post", controller.MainController.GetAll)
	app.Get("/blog/post/:id", controller.MainController.GetById)
	app.Get("/blog/post/get/:username", controller.MainController.GetByUsername)
	app.Delete("/blog/post", controller.MainController.Delete)

	return app
}
