package todo

import (
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

func (srv *todoService) Create(todo *model.Todo) error {
	// fmt.Print(todo)

	// err := srv.todoRepository.Create(todo)
	// if err != nil {
	// 	return err
	// }

	return nil
}

func (srv *todoService) FindById(id string) (*model.Todo, error) {
	// fmt.Print(id)

	// res, err := srv.(id)
	// if err != nil {
	// 	return nil, err
	// }

	return nil, nil
}

func (srv *todoService) FindAll() ([]*model.Todo, error) {
	return nil, nil
}

func (srv *todoService) Update(todo *model.Todo) error {
	return nil
}

func (srv *todoService) Delete(id string) error {
	return nil
}
