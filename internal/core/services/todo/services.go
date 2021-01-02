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
	// fmt.Print(todo)

	// err := srv.todoRepository.Create(todo)
	// if err != nil {
	// 	return err
	// }

	return nil
}

func (srv *todoService) FindById(ctx context.Context, id string) (*model.Todo, error) {
	// fmt.Print(id)

	// res, err := srv.(id)
	// if err != nil {
	// 	return nil, err
	// }

	return nil, nil
}

func (srv *todoService) FindAll(ctx context.Context) ([]*model.Todo, error) {
	return nil, nil
}

func (srv *todoService) Update(ctx context.Context, id string, todo *model.Todo) error {
	return nil
}

func (srv *todoService) Delete(ctx context.Context, id string) error {
	return nil
}
