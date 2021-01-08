package handler

import (
	model "github.com/famesensor/playground-go-fiber-todonotes/internal/core/domain"
	interfaces "github.com/famesensor/playground-go-fiber-todonotes/internal/core/ports"
	"github.com/famesensor/playground-go-fiber-todonotes/pkg/reponseHandler"
	"github.com/gofiber/fiber/v2"
)

type todoHandler struct {
	todoService interfaces.TodoService
}

func NewTodoHandler(todoService interfaces.TodoService) interfaces.TodoHandler {
	return &todoHandler{
		todoService,
	}
}

func (h *todoHandler) Create(c *fiber.Ctx) error {
	todo := new(model.Todo)

	if err := c.BodyParser(&todo); err != nil {
		return c.JSON("Error na ka")
	}

	if err := h.todoService.Create(c.Context(), todo); err != nil {
		return c.JSON("Error na ka")
	}

	return reponseHandler.ReponseMsg(c, 201, "success", "", nil)
}

func (h *todoHandler) FindById(c *fiber.Ctx) error {
	id := c.Params("id")

	res, err := h.todoService.FindById(c.Context(), id)
	if err != nil {
		return c.JSON("Error na ka")
	}

	return reponseHandler.ReponseMsg(c, 200, "success", "", res)
}

func (h *todoHandler) FindAll(c *fiber.Ctx) error {
	return c.JSON("Error na ka")
}

func (h *todoHandler) Update(c *fiber.Ctx) error {
	return nil
}

func (h *todoHandler) Delete(c *fiber.Ctx) error {
	return nil
}
