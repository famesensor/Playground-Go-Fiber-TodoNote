package ports

import (
	"context"

	model "github.com/famesensor/playground-go-fiber-todonotes/internal/core/domain"
)

type TodoRepository interface {
	Create(ctx context.Context, todo *model.Todo) error
	FindById(ctx context.Context, id string) (*model.Todo, error)
	FindAll(ctx context.Context) ([]*model.Todo, error)
	Update(ctx context.Context, id string, todo *model.Todo) error
	Delete(ctx context.Context, id string) error
}
