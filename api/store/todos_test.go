package store_test

import (
	"context"
	"database/sql"
	"flag"
	"github.com/joho/godotenv"
	"log"
	"prometheus-alert-manager-tutorial/api/db"
	"prometheus-alert-manager-tutorial/api/store"
	"prometheus-alert-manager-tutorial/api/types"
	"prometheus-alert-manager-tutorial/api/utils"
	"sort"
	"testing"
	"time"
)

func init() {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Fatalln(err)
	}
	db.LoadConfigs(flag.CommandLine)
}

func TestTodosStore_Create(t *testing.T) {

	conn := db.New()
	defer utils.HandleClose(conn)
	ts := store.NewTodosStore(conn)
	ctx := context.Background()

	todo := &types.Todo{
		Title:       "CREATE TODO",
		Description: "TESTING TODO CREATION",
		Done:        false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	created, err := ts.Create(ctx, todo)
	if err != nil {
		t.Errorf("unexpected error on create todo: %s", err.Error())
	}
	if created.ID < 1 {
		t.Errorf("unexpected id on create todo: expected > 0, got=%d", created.ID)
	}
	if created.Title != todo.Title {
		t.Errorf("unexpected title on create todo: expected=%s, got=%s", todo.Title, created.Title)
	}
	if created.Description != todo.Description {
		t.Errorf("unexpected description on create todo: expected=%s, got=%s", todo.Description, created.Description)
	}
	if created.Done != todo.Done {
		t.Errorf("unexpected boolean field on create todo: expected=%v, got=%v", todo.Done, created.Done)
	}
	if created.CreatedAt.IsZero() {
		t.Errorf("unexpected time field on create todo: expected=%v, got=%v", todo.CreatedAt.String(), created.CreatedAt.String())
	}
	if created.UpdatedAt.IsZero() {
		t.Errorf("unexpected time field on create todo: expected=%v, got=%v", todo.UpdatedAt.String(), created.UpdatedAt.String())
	}
}

func TestTodosStore_Update(t *testing.T) {

	conn := db.New()
	defer utils.HandleClose(conn)
	ts := store.NewTodosStore(conn)
	ctx := context.Background()

	todo := &types.Todo{
		Title:       "UPDATE TODO",
		Description: "CREATE TODO BEFORE UPDATE",
		Done:        false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	created, err := ts.Create(ctx, todo)
	if err != nil {
		t.Errorf("unexpected error on create todo: %s", err.Error())
	}

	created.Description = "TESTING TODO UPDATING"
	created.Done = true
	created.UpdatedAt = time.Now()

	updated, err := ts.Update(ctx, created)
	if err != nil {
		t.Errorf("unexpected error on update todo: %s", err.Error())
	}
	if updated.ID != created.ID {
		t.Errorf("unexpected id on update todo: expected=%d, got=%d", created.ID, updated.ID)
	}
	if created.Title != updated.Title {
		t.Errorf("unexpected title on update todo: expected=%s, got=%s", created.Title, updated.Title)
	}
	if created.Description != updated.Description {
		t.Errorf("unexpected description on update todo: expected=%s, got=%s", created.Description, updated.Description)
	}
	if !updated.Done {
		t.Errorf("unexpected boolean field on update todo: expected=%v, got=%v", true, updated.Done)
	}
	if updated.CreatedAt.IsZero() {
		t.Errorf("unexpected time field on update todo: expected=%v, got=%v", updated.CreatedAt.String(), created.CreatedAt.String())
	}
	if updated.UpdatedAt.IsZero() {
		t.Errorf("unexpected time field on update todo: expected=%v, got=%v", updated.UpdatedAt.String(), created.UpdatedAt.String())
	}
}

func TestTodosStore_Get(t *testing.T) {

	conn := db.New()
	defer utils.HandleClose(conn)
	ts := store.NewTodosStore(conn)
	ctx := context.Background()

	todo := &types.Todo{
		Title:       "GET TODO",
		Description: "TESTING GET TODO BY ID",
		Done:        false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	created, err := ts.Create(ctx, todo)
	if err != nil {
		t.Errorf("unexpected error on create todo: %s", err.Error())
	}

	found, err := ts.Get(ctx, created.ID)
	if err != nil {
		t.Errorf("unexpected error on get todo: %s", err.Error())
	}
	if found.ID != created.ID {
		t.Errorf("unexpected id on get todo: expected=%d, got=%d", created.ID, found.ID)
	}
	if created.Title != found.Title {
		t.Errorf("unexpected title on get todo: expected=%s, got=%s", created.Title, found.Title)
	}
	if created.Description != found.Description {
		t.Errorf("unexpected description on get todo: expected=%s, got=%s", created.Description, found.Description)
	}
	if created.Done != found.Done {
		t.Errorf("unexpected boolean field on get todo: expected=%v, got=%v", true, found.Done)
	}
	if found.CreatedAt.IsZero() {
		t.Errorf("unexpected time field on get todo: expected=%v, got=%v", found.CreatedAt.String(), created.CreatedAt.String())
	}
	if found.UpdatedAt.IsZero() {
		t.Errorf("unexpected time field on get todo: expected=%v, got=%v", found.UpdatedAt.String(), created.UpdatedAt.String())
	}
}

func TestTodosStore_GetAll(t *testing.T) {

	conn := db.New()
	defer utils.HandleClose(conn)
	ts := store.NewTodosStore(conn)
	ctx := context.Background()

	todo := &types.Todo{
		Title:       "GET TODOS",
		Description: "TESTING GET ALL TODOS",
		Done:        false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	created, err := ts.Create(ctx, todo)
	if err != nil {
		t.Errorf("unexpected error on create todo: %s", err.Error())
	}

	list, err := ts.GetAll(ctx)
	if err != nil {
		t.Errorf("unexpected error on get todo: %s", err.Error())
	}
	if len(list) == 0 {
		t.Errorf("unexpected list on get all todos: expected > 0, got=%d", len(list))
	}

	index := sort.Search(len(list), func(i int) bool {
		return list[i].ID == created.ID
	})

	if index < 0 {
		t.Errorf("unexpected item on get all todos: expected=%v, got=%v", created, list)
	}

	if list[index].ID != created.ID {
		t.Errorf("unexpected item on get all todos: expected=%v, got=%v", created, list[index])
	}
}

func TestTodosStore_Delete(t *testing.T) {

	conn := db.New()
	defer utils.HandleClose(conn)
	ts := store.NewTodosStore(conn)
	ctx := context.Background()

	todo := &types.Todo{
		Title:       "DELETE TODO",
		Description: "TESTING DELETE TODO",
		Done:        false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	created, err := ts.Create(ctx, todo)
	if err != nil {
		t.Errorf("unexpected error on create todo: %s", err.Error())
	}

	err = ts.Delete(ctx, created.ID)
	if err != nil {
		t.Errorf("unexpected error on delete todo: %s", err.Error())
	}

	_, err = ts.Get(ctx, created.ID)
	if err.Error() != sql.ErrNoRows.Error() {
		t.Errorf("unexpected error on delete todo: expected=%s, got=%s", err.Error(), sql.ErrNoRows.Error())
	}
}
