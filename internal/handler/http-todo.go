package handler

import (
	model "github.com/famesensor/playground-go-fiber-todonotes/internal/core/domain"
	interfaces "github.com/famesensor/playground-go-fiber-todonotes/internal/core/ports"
	"github.com/famesensor/playground-go-fiber-todonotes/pkg/errs"
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
		return errs.ErrorReponse(c, errs.CannotParseData, nil)
	}

	// TODO: Validate data

	if err := h.todoService.Create(c.Context(), todo); err != nil {
		return errs.ErrorReponse(c, err, nil)
	}

	return reponseHandler.ReponseMsg(c, 201, "success", "", nil)
}

func (h *todoHandler) FindById(c *fiber.Ctx) error {
	id := c.Query("todoId")

	// TODO: Validate ID

	res, err := h.todoService.FindById(c.Context(), id)
	if err != nil {
		return errs.ErrorReponse(c, err, nil)
	}

	return reponseHandler.ReponseMsg(c, 200, "success", "", res)
}

func (h *todoHandler) FindAll(c *fiber.Ctx) error {
	todos, err := h.todoService.FindAll(c.Context())
	if err != nil {
		return errs.ErrorReponse(c, err, nil)
	}
	return reponseHandler.ReponseMsg(c, 200, "success", "", todos)
}

func (h *todoHandler) UpdateTodo(c *fiber.Ctx) error {
	todo := new(model.Todo)
	id := c.Query("id")

	if err := c.BodyParser(&todo); err != nil {
		return errs.ErrorReponse(c, errs.CannotParseData, nil)
	}

	// TODO: Validate Date

	if err := h.todoService.UpdateTodo(c.Context(), id, todo); err != nil {
		return errs.ErrorReponse(c, err, nil)
	}

	return reponseHandler.ReponseMsg(c, 200, "success", "", nil)
}

func (h *todoHandler) UpdateStatus(c *fiber.Ctx) error {
	id := c.Query("id")

	// TODO: Validate ID

	if err := h.todoService.UpdateStatus(c.Context(), id); err != nil {
		return errs.ErrorReponse(c, err, nil)
	}

	return reponseHandler.ReponseMsg(c, 200, "success", "", nil)
}
