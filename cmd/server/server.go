package server

import "github.com/Krypt9x/blog-backend/internal/api/routes"

func RunServer() {
	route := routes.MainRoute{}
	app := route.NewRouter()
	app.Listen(":3000")
}
