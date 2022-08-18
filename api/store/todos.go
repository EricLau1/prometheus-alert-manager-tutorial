package store

import (
	"context"
	"database/sql"
	"prometheus-alert-manager-tutorial/api/types"
	"prometheus-alert-manager-tutorial/api/utils"
	"strconv"
)

type TodosStore interface {
	Create(ctx context.Context, todo *types.Todo) (*types.Todo, error)
	Update(ctx context.Context, todo *types.Todo) (*types.Todo, error)
	Get(ctx context.Context, id int64) (*types.Todo, error)
	GetAll(ctx context.Context) ([]*types.Todo, error)
	Delete(ctx context.Context, id int64) error
}

type todosStore struct {
	conn *sql.DB
}

func NewTodosStore(conn *sql.DB) TodosStore {
	return &todosStore{conn: conn}
}

func (s *todosStore) Create(ctx context.Context, todo *types.Todo) (*types.Todo, error) {
	query := `insert into todos (title, description, created_at, updated_at) values (?,?,?,?)`
	stmt, err := s.conn.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer utils.HandleClose(stmt)
	result, err := stmt.ExecContext(ctx, todo.Title, todo.Description, todo.CreatedAt, todo.UpdatedAt)
	if err != nil {
		return nil, err
	}
	todo.ID, err = result.LastInsertId()
	return todo, err
}

func (s *todosStore) Update(ctx context.Context, todo *types.Todo) (*types.Todo, error) {
	query := `update todos set title=?,description=?,done=?,updated_at=? where id=?`
	stmt, err := s.conn.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer utils.HandleClose(stmt)
	_, err = stmt.ExecContext(ctx, todo.Title, todo.Description, strconv.FormatBool(todo.Done), todo.UpdatedAt, todo.ID)
	if err != nil {
		return nil, err
	}
	return todo, err
}

func (s *todosStore) Get(ctx context.Context, id int64) (*types.Todo, error) {
	query := `select * from todos where id=?`
	row := s.conn.QueryRowContext(ctx, query, id)
	todo := new(types.Todo)
	var done string
	err := row.Scan(&todo.ID, &todo.Title, &todo.Description, &done, &todo.CreatedAt, &todo.UpdatedAt)
	if err != nil {
		return nil, err
	}
	todo.Done = done == "true"
	return todo, err
}

func (s *todosStore) GetAll(ctx context.Context) ([]*types.Todo, error) {
	query := `select * from todos`
	rows, err := s.conn.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer utils.HandleClose(rows)
	var todos []*types.Todo
	for rows.Next() {
		todo := new(types.Todo)
		var done string
		err = rows.Scan(&todo.ID, &todo.Title, &todo.Description, &done, &todo.CreatedAt, &todo.UpdatedAt)
		if err != nil {
			return nil, err
		}
		todo.Done = done == "true"
		todos = append(todos, todo)
	}
	return todos, nil
}

func (s *todosStore) Delete(ctx context.Context, id int64) error {
	query := `delete from todos where id=?`
	row := s.conn.QueryRowContext(ctx, query, id)
	return row.Err()
}
