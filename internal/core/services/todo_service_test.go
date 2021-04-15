package services_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/famesensor/playground-go-fiber-todonotes/internal/core/domain"
	"github.com/famesensor/playground-go-fiber-todonotes/internal/core/services"
	"github.com/famesensor/playground-go-fiber-todonotes/internal/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/tj/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var mockTodoRepo *mocks.TodoRepository
var ctx context.Context
var unexpectedError = errors.New("Unexpected Error")

func TestMain(m *testing.M) {
	mockTodoRepo = new(mocks.TodoRepository)
	ctx = context.Background()
	code := m.Run()
	os.Exit(code)
}

func TestCreate(t *testing.T) {
	mockRequest := &domain.Todo{
		Title:       "Unit Testing",
		Description: "Unit Testing",
		Status:      "Waiting",
		Priority:    1,
		CreatedAt:   time.Now(),
	}

	t.Run("success", func(t *testing.T) {
		mockTodoRepo.On("Create", context.TODO(), mock.AnythingOfType("*domain.Todo")).Return(nil).Once()
		s := services.NewTodoService(mockTodoRepo)
		err := s.Create(ctx, mockRequest)
		assert.NoError(t, err)
		mockTodoRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockTodoRepo.On("Create", context.TODO(), mock.AnythingOfType("*domain.Todo")).Return(unexpectedError).Once()
		s := services.NewTodoService(mockTodoRepo)
		err := s.Create(ctx, mockRequest)
		assert.Error(t, err)
		mockTodoRepo.AssertExpectations(t)
	})
}

func TestFindById(t *testing.T) {
	mockTodoId := "todo_uuid_id"
	mockTodoRes := &domain.Todo{
		ID:          primitive.NewObjectID(),
		Title:       "title todo",
		Description: "description todo",
		Status:      "success",
		Priority:    1,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	t.Run("success", func(t *testing.T) {
		mockTodoRepo.On("FindById", context.TODO(), mock.AnythingOfType("string")).Return(mockTodoRes, nil).Once()
		s := services.NewTodoService(mockTodoRepo)
		_, err := s.FindById(ctx, mockTodoId)
		assert.NoError(t, err)
		mockTodoRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockTodoRepo.On("FindById", context.TODO(), mock.AnythingOfType("string")).Return(mockTodoRes, unexpectedError).Once()
		s := services.NewTodoService(mockTodoRepo)
		_, err := s.FindById(ctx, mockTodoId)
		assert.Error(t, err)
		mockTodoRepo.AssertExpectations(t)
	})
}

func TestFindAll(t *testing.T) {
	mockTodoRes := []*domain.Todo{
		{
			ID:          primitive.NewObjectID(),
			Title:       "title todo",
			Description: "description todo",
			Status:      "success",
			Priority:    1,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}

	t.Run("success", func(t *testing.T) {
		mockTodoRepo.On("FindAll", context.TODO()).Return(mockTodoRes, nil).Once()
		s := services.NewTodoService(mockTodoRepo)
		_, err := s.FindAll(ctx)
		assert.NoError(t, err)
		mockTodoRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockTodoRepo.On("FindAll", context.TODO()).Return(mockTodoRes, unexpectedError).Once()
		s := services.NewTodoService(mockTodoRepo)
		_, err := s.FindAll(ctx)
		assert.Error(t, err)
		mockTodoRepo.AssertExpectations(t)
	})
}

func TestUpdateTodo(t *testing.T) {
	mockTodoId := "todo_uuid_id"
	mockTodoRes := &domain.Todo{
		Title:       "title todo",
		Description: "description todo",
		Status:      "success",
		Priority:    1,
	}

	t.Run("success", func(t *testing.T) {
		mockTodoRepo.On("UpdateTodo", context.TODO(), mock.AnythingOfType("string"), mock.AnythingOfType("*domain.Todo")).Return(nil).Once()
		s := services.NewTodoService(mockTodoRepo)
		err := s.UpdateTodo(ctx, mockTodoId, mockTodoRes)
		assert.NoError(t, err)
		mockTodoRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockTodoRepo.On("UpdateTodo", context.TODO(), mock.AnythingOfType("string"), mock.AnythingOfType("*domain.Todo")).Return(unexpectedError).Once()
		s := services.NewTodoService(mockTodoRepo)
		err := s.UpdateTodo(ctx, mockTodoId, mockTodoRes)
		assert.Error(t, err)
		mockTodoRepo.AssertExpectations(t)
	})
}

func TestUpdateStatus(t *testing.T) {
	mockTodoId := "todo_uuid_id"
	mockTodoRes := &domain.Todo{
		ID:          primitive.NewObjectID(),
		Title:       "title todo",
		Description: "description todo",
		Status:      "in progress",
		Priority:    1,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	t.Run("success", func(t *testing.T) {
		mockTodoRepo.On("FindById", context.TODO(), mock.AnythingOfType("string")).Return(mockTodoRes, nil).Once()
		mockTodoRepo.On("UpdateStatus", context.TODO(), mock.AnythingOfType("string")).Return(nil).Once()
		s := services.NewTodoService(mockTodoRepo)
		err := s.UpdateStatus(ctx, mockTodoId)
		assert.NoError(t, err)
		mockTodoRepo.AssertExpectations(t)
	})

	t.Run("success", func(t *testing.T) {
		mockTodoRes.Status = "success"
		mockTodoRepo.On("FindById", context.TODO(), mock.AnythingOfType("string")).Return(mockTodoRes, nil).Once()
		s := services.NewTodoService(mockTodoRepo)
		err := s.UpdateStatus(ctx, mockTodoId)
		assert.NoError(t, err)
		mockTodoRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockTodoRepo.On("FindById", context.TODO(), mock.AnythingOfType("string")).Return(mockTodoRes, unexpectedError).Once()
		s := services.NewTodoService(mockTodoRepo)
		err := s.UpdateStatus(ctx, mockTodoId)
		assert.Error(t, err)
		mockTodoRepo.AssertExpectations(t)
	})
	t.Run("error", func(t *testing.T) {
		mockTodoRes.Status = "in progress"
		mockTodoRepo.On("FindById", context.TODO(), mock.AnythingOfType("string")).Return(mockTodoRes, nil).Once()
		mockTodoRepo.On("UpdateStatus", context.TODO(), mock.AnythingOfType("string")).Return(unexpectedError).Once()
		s := services.NewTodoService(mockTodoRepo)
		err := s.UpdateStatus(ctx, mockTodoId)
		assert.Error(t, err)
		mockTodoRepo.AssertExpectations(t)
	})
}
