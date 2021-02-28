package services_test

import (
	"context"
	"testing"
	"time"

	"github.com/famesensor/playground-go-fiber-todonotes/internal/core/domain"
	"github.com/famesensor/playground-go-fiber-todonotes/internal/core/services"
	_mock "github.com/famesensor/playground-go-fiber-todonotes/internal/mocks"
	"github.com/tj/assert"
)

func TestCreate(t *testing.T) {
	mockRequest := &domain.Todo{
		Title:       "Unit Testing",
		Description: "Unit Testing",
		Status:      "Waiting",
		Priority:    1,
		CreatedAt:   time.Date(2021, 03, 01, 12, 00, 00, 00, time.Local),
	}

	t.Run("success", func(t *testing.T) {
		mockTodoRepo := new(_mock.TodoRepository)

		mockTodoRepo.On("Create", context.TODO(), mockRequest).Return(nil).Once()

		u := services.NewTodoService(mockTodoRepo)

		err := u.Create(context.TODO(), mockRequest)

		assert.NoError(t, err)
		// assert.Equal(t)
		mockTodoRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockTodoRepo := new(_mock.TodoRepository)
		u := services.NewTodoService(mockTodoRepo)

		err := u.Create(context.TODO(), mockRequest)

		assert.Error(t, err)
		// assert.Equal(t)
		mockTodoRepo.AssertExpectations(t)
	})
}

func TestFindById(t *testing.T) {

}

func TestFindAll(t *testing.T) {

}

func TestUpdateTodo(t *testing.T) {

}

func TestUpdateStatus(t *testing.T) {

}
