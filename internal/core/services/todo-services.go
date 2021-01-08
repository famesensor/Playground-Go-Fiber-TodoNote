package todo

import (
	"context"

	model "github.com/famesensor/playground-go-fiber-todonotes/internal/core/domain"
	"github.com/famesensor/playground-go-fiber-todonotes/internal/core/ports"
)

type todoService struct {
	todoRepository ports.TodoRepository
}

func New(todoRepository ports.TodoRepository) ports.TodoService {
	return &todoService{
		todoRepository,
	}
}

func (srv *todoService) Create(ctx context.Context, todo *model.Todo) error {
	return srv.todoRepository.Create(ctx, todo)
}

func (srv *todoService) FindById(ctx context.Context, id string) (*model.Todo, error) {
	return srv.todoRepository.FindById(ctx, id)
}

func (srv *todoService) FindAll(ctx context.Context) ([]*model.Todo, error) {
	return srv.todoRepository.FindAll(ctx)
}

func (srv *todoService) Update(ctx context.Context, id string, todo *model.Todo) error {
	return nil
}

func (srv *todoService) Delete(ctx context.Context, id string) error {
	return nil
}
