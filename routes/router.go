package routes

import (
	"mangojek-backend/controller"

	"github.com/gofiber/fiber/v2"
)

func middleware(c *fiber.Ctx) error {
	token := c.Get("token")
	if token != "RAHASIA" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status": "UNAUTHORIZED",
		})
	}
	return c.Next()
}
func Route(app *fiber.App, controller controller.UserController) {
	app.Get("/api/users", middleware, controller.FindAll)
	app.Post("/api/users", middleware, controller.Insert)
}
