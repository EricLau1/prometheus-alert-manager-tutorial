package services

import (
	"context"
	"prometheus-alert-manager-tutorial/api/store"
	"prometheus-alert-manager-tutorial/api/types"
	"time"
)

type TodosService interface {
	CreateTodo(ctx context.Context, todo *types.Todo) (*types.Todo, error)
	UpdateTodo(ctx context.Context, todo *types.Todo) (*types.Todo, error)
	GetTodo(ctx context.Context, id int64) (*types.Todo, error)
	GetTodos(ctx context.Context) ([]*types.Todo, error)
	DeleteTodo(ctx context.Context, id int64) error
}

type todosService struct {
	todosStore store.TodosStore
}

func NewTodosService(todosStore store.TodosStore) TodosService {
	return &todosService{
		todosStore: todosStore,
	}
}

func (s *todosService) CreateTodo(ctx context.Context, todo *types.Todo) (*types.Todo, error) {
	todo.CreatedAt = time.Now()
	todo.UpdatedAt = todo.CreatedAt
	return s.todosStore.Create(ctx, todo)
}

func (s *todosService) UpdateTodo(ctx context.Context, todo *types.Todo) (*types.Todo, error) {
	_, err := s.todosStore.Get(ctx, todo.ID)
	if err != nil {
		return nil, err
	}
	todo.UpdatedAt = time.Now()
	return s.todosStore.Update(ctx, todo)
}

func (s *todosService) GetTodo(ctx context.Context, id int64) (*types.Todo, error) {
	return s.todosStore.Get(ctx, id)
}

func (s *todosService) GetTodos(ctx context.Context) ([]*types.Todo, error) {
	return s.todosStore.GetAll(ctx)
}

func (s *todosService) DeleteTodo(ctx context.Context, id int64) error {
	return s.todosStore.Delete(ctx, id)
}
