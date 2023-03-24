package server

import "github.com/Krypt9x/blog-backend/internal/api/routes"

func RunServer() {
	app := routes.Routes()
	app.Listen(":3000")
}
