package routes

import (
	"mangojek-backend/controller"

	"github.com/gofiber/fiber/v2"
)

func Route(app *fiber.App, controller controller.UserController) {
	app.Get("/api/users", controller.FindAll)
	app.Post("/api/users", controller.Insert)
}
