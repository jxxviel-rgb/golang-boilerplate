package controller

import "github.com/gofiber/fiber/v2"

type UserController interface {
	Route(app *fiber.App)
	FindAll(c *fiber.Ctx) error
	Insert(c *fiber.Ctx) error
}
