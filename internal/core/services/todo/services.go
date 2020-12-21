package todo

import (
	"fmt"

	"github.com/famesensor/playground-go-fiber-todonotes/internal/core/domain"
	"github.com/famesensor/playground-go-fiber-todonotes/internal/core/ports"
)

type service struct {
	todoRepository ports.TodoRepository
}

func New(todoRepository ports.TodoRepository) *service {
	return &service{
		todoRepository,
	}
}

func (srv *service) CreateTodo(todo *domain.Todo) error {
	fmt.Print(todo)

	err := srv.todoRepository.Create(todo)
	if err != nil {
		return err
	}

	return nil
}

func (srv *service) FindTodoId(id string) (*domain.Todo, error) {
	fmt.Print(id)

	res, err := srv.FindTodoId(id)
	if err != nil {
		return nil, err
	}

	return res, nil
}
