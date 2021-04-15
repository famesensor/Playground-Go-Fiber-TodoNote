package ports

import "github.com/gofiber/fiber/v2"

type TodoHandler interface {
	// // Create create a new book data
	// // @Summary Create a new todo
	// // @Description Create book
	// // @Tags todo
	// // @Accept json
	// // @Produce json
	// // @Param todo body model.Todo true "Create book"
	// // @Success 200 {object} ResponseHTTP{}
	// // @Failure 400 {object} ResponseHTTP{}
	// // @Router  /api/v1/create [post]
	Create(c *fiber.Ctx) error

	// // FindById is a function to get a todo by ID
	// // @Summary Get todo by ID
	// // @Description Get todo by ID
	// // @Tags todos
	// // @Accept json
	// // @Produce json
	// // @Param id path int true "TODO ID"
	// // @Success 200 {object} ResponseHTTP{data=[]model.Todo}
	// // @Failure 404 {object} ResponseHTTP{}
	// // @Failure 500 {object} ResponseHTTP{}
	// // @Router /api/v1/:id [get]
	FindById(c *fiber.Ctx) error

	// FindAll is a function to getall Todo data from database
	// @Summary Get all Todos
	// @Description Get all todos
	// @Tags todos
	// @Accept json
	// @Produce json
	// @Success 200 {object} ResponseHTTP{data=[]model.Todo}
	// @Failure 404 {object} ResponseHTTP{}
	// @Failure 500 {object} ResponseHTTP{}
	// @Router /todo/todos [get]
	FindAll(c *fiber.Ctx) error

	// // UpdateTodo function removes a book by ID
	// // @Summary Remove book by ID
	// // @Description Remove book by ID
	// // @Tags books
	// // @Accept json
	// // @Produce json
	// // @Param id path int true "Book ID"
	// // @Success 200 {object} ResponseHTTP{}
	// // @Failure 404 {object} ResponseHTTP{}
	// // @Failure 503 {object} ResponseHTTP{}
	// // @Router /v1/books/{id} [delete]
	UpdateTodo(c *fiber.Ctx) error

	// // DeleteBook function removes a book by ID
	// // @Summary Remove book by ID
	// // @Description Remove book by ID
	// // @Tags books
	// // @Accept json
	// // @Produce json
	// // @Param id path int true "Book ID"
	// // @Success 200 {object} ResponseHTTP{}
	// // @Failure 404 {object} ResponseHTTP{}
	// // @Failure 503 {object} ResponseHTTP{}
	// // @Router /v1/books/{id} [delete]
	UpdateStatus(c *fiber.Ctx) error
}
