package services

import (
	"context"

	model "github.com/famesensor/playground-go-fiber-todonotes/internal/core/domain"
	"github.com/famesensor/playground-go-fiber-todonotes/internal/core/ports"
)

type todoService struct {
	todoRepository ports.TodoRepository
}

func NewTodoService(todoRepository ports.TodoRepository) ports.TodoService {
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

func (srv *todoService) UpdateTodo(ctx context.Context, id string, todo *model.Todo) error {
	return srv.todoRepository.UpdateTodo(ctx, id, todo)
}

func (srv *todoService) UpdateStatus(ctx context.Context, id string) error {
	todoDoc, err := srv.FindById(ctx, id)
	if err != nil {
		return err
	}

	if todoDoc.Status != "success" {
		return srv.todoRepository.UpdateStatus(ctx, id)
	}

	return nil
}
