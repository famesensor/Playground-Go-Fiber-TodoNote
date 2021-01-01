package ports

import model "github.com/famesensor/playground-go-fiber-todonotes/internal/core/domain"

type TodoService interface {
	Create(todo *model.Todo) error
	FindById(id string) (*model.Todo, error)
	FindAll() ([]*model.Todo, error)
	Update(todo *model.Todo) error
	Delete(id string) error
}
