package handler

import (
	"github.com/famesensor/playground-go-fiber-todonotes/internal/core/domain"
	"github.com/famesensor/playground-go-fiber-todonotes/internal/core/ports"
	"github.com/gofiber/fiber/v2"
)

type TodoHandler interface {
	Create(c *fiber.Ctx) error
	FindById(c *fiber.Ctx) error
	FindAll(c *fiber.Ctx) error
}

type todoHandler struct {
	todoService ports.TodoService
}

func New(todoService ports.TodoService) TodoHandler {
	return &todoHandler{
		todoService,
	}
}

func (h *todoHandler) Create(c *fiber.Ctx) error {
	var todo domain.Todo

	if err := c.BodyParser(todo); err != nil {
		return c.JSON("Error na ka")
	}

	if err := h.todoService.Create(&todo); err != nil {
		return c.JSON("Error na ka")
	}

	return c.JSON(fiber.Map{
		"status": true,
		"data":   "success",
	})

}

func (h *todoHandler) FindById(c *fiber.Ctx) error {
	id := c.Params("id")

	res, err := h.todoService.FindById(id)
	if err != nil {
		return c.JSON("Error na ka")
	}

	return c.JSON(fiber.Map{
		"Status": true,
		"data":   res,
	})
}

func (h *todoHandler) FindAll(c *fiber.Ctx) error {
	return c.JSON("Error na ka")
}
