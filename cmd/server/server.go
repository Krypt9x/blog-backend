package server

import (
	"github.com/Krypt9x/blog-backend/internal/api/controller"
	"github.com/Krypt9x/blog-backend/internal/api/repository"
	"github.com/Krypt9x/blog-backend/internal/api/routes"
	services "github.com/Krypt9x/blog-backend/internal/api/service"
	concurrentservice "github.com/Krypt9x/blog-backend/internal/api/service/concurrent_service"
	"github.com/Krypt9x/blog-backend/internal/database"
	"github.com/gofiber/fiber/v2/middleware/pprof"
)

func RunServer() {

	dbObj := database.InitDBService{
		DataSource: "postgres://root:secret@localhost:5432/blog?sslmode=disable",
	}

	db := dbObj.InitDB()

	amountRepo := repository.NewAmountRepository()
	amountService := services.NewAmountService(db, amountRepo)

	commentRepo := repository.NewCommentRepository()
	commentService := services.NewCommentService(db, commentRepo)

	userRepo := repository.NewUsersRepository()
	userService := services.NewUsersService(db, userRepo)
	userController := controller.NewUserController(userService)

	mainRepo := repository.NewMainRepository()
	mainService := services.NewMainService(db, mainRepo)

	amountConcurrentService := concurrentservice.NewAmountService(amountService)
	commentConcurrentService := concurrentservice.NewCommentService(commentService)

	concurrentService := concurrentservice.NewConcurrentService(amountConcurrentService, commentConcurrentService)

	mainController := controller.NewMainController(mainService, *concurrentService)

	objController := routes.ControllerObj{
		UserController: userController,
		MainController: mainController,
	}

	app := routes.InitRouter(objController)
	app.Listen(":3000")
	app.Use(pprof.New())
}
