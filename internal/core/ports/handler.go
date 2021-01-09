package ports

import "github.com/gofiber/fiber/v2"

type TodoHandler interface {
	Create(c *fiber.Ctx) error
	FindById(c *fiber.Ctx) error
	FindAll(c *fiber.Ctx) error
	UpdateTodo(c *fiber.Ctx) error
	UpdateStatus(c *fiber.Ctx) error
}
