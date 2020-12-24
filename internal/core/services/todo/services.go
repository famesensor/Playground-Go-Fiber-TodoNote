package todo

import (
	"github.com/famesensor/playground-go-fiber-todonotes/internal/core/domain"
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

func (srv *todoService) Create(todo *domain.Todo) error {
	// fmt.Print(todo)

	// err := srv.todoRepository.Create(todo)
	// if err != nil {
	// 	return err
	// }

	return nil
}

func (srv *todoService) FindById(id string) (*domain.Todo, error) {
	// fmt.Print(id)

	// res, err := srv.(id)
	// if err != nil {
	// 	return nil, err
	// }

	return nil, nil
}

func (srv *todoService) FindAll() ([]*domain.Todo, error) {
	return nil, nil
}
