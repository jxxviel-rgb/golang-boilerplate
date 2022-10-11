package main

import (
	"mangojek-backend/config"
	"mangojek-backend/controller"
	"mangojek-backend/exception"
	"mangojek-backend/repository"
	"mangojek-backend/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	configuration := config.NewConfiguration()
	database := config.NewPostgresDatabase(configuration)
	config.NewRunMigration(database)

	// Setup Repository
	userRepository := repository.NewUserRepositoryImpl(database)

	// Setup Service
	userService := service.NewUserServiceImpl(userRepository)

	// Setup Controller
	userController := controller.NewUserControllerImpl(userService)

	// Setup Fiber
	// app := fiber.New(config.NewFiberConfig())
	app := fiber.New(config.NewFiberConfig())
	app.Use(recover.New())

	// Setup Routing
	userController.Route(app)

	// Start App
	err := app.Listen(":3000")
	exception.PanicIfNeeded(err)
}
