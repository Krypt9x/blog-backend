package server

import (
	"github.com/Krypt9x/blog-backend/internal/api/routes"
	"github.com/gofiber/fiber/v2/middleware/pprof"
)

func RunServer() {
	route := routes.MainRoute{}
	app := route.NewRouter()
	app.Listen(":3000")
	app.Use(pprof.New())
}
