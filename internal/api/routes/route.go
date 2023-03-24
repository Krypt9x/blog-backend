package routes

import "github.com/gofiber/fiber/v2"

func Routes() *fiber.App {
	app := fiber.New()
	app.Get("/about", func(c *fiber.Ctx) error {
		return c.SendString("ini halaman about")
	})

	return app
}
