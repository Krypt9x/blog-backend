package routes

import "github.com/gofiber/fiber/v2"

type (
	Routes interface {
		NewRouter() *fiber.App
	}

	Route struct {
	}
)

func (route *Route) NewRouter() *fiber.App {
	app := fiber.New()
	app.Get("/about", func(c *fiber.Ctx) error {
		return c.SendString("ini halaman about")
	})

	return app
}
