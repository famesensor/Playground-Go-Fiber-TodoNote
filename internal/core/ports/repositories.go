package ports

import "github.com/famesensor/playground-go-fiber-todonotes/internal/core/domain"

type TodoRepository interface {
	Create(user *domain.Todo) error
	FindById(id string) (*domain.Todo, error)
	FindAll() ([]*domain.Todo, error)
}
