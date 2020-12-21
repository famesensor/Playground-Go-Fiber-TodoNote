package handler

import (
	"github.com/famesensor/playground-go-fiber-todonotes/internal/core/domain"
	"github.com/famesensor/playground-go-fiber-todonotes/internal/core/ports"
	"github.com/gofiber/fiber/v2"
)

type todoHandler struct {
	todoService ports.TodoService
}

func New(todoService ports.TodoService) *todoHandler {
	return &todoHandler{
		todoService,
	}
}

func (h *todoHandler) CreateTodo(c *fiber.Ctx) {
	var todo domain.Todo

	if err := c.BodyParser(todo); err != nil {
		return
	}

	if err := h.todoService.Create(&todo); err != nil {
		return
	}

	c.JSON(fiber.Map{
		"status": true,
		"data":   "success",
	})

}

func (h *todoHandler) FindTodoId(c *fiber.Ctx) {
	id := c.Params("id")

	res, err := h.todoService.FindById(id)
	if err != nil {
		return
	}

	c.JSON(fiber.Map{
		"Status": true,
		"data":   res,
	})
}
