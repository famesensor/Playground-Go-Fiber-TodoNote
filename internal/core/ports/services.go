package ports

import "github.com/famesensor/playground-go-fiber-todonotes/internal/core/domain"

type TodoService interface {
	Create(user *domain.Todo) error
	FindById(id string) (*domain.Todo, error)
}
