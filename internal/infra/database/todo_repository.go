package database

import (
	"database/sql"

	"github.com/FelipeRomao/todo/internal/domain/entities"
)

type TodoRepository struct {
	Db *sql.DB
}

func NewTodoRepository(db *sql.DB) *TodoRepository {
	return &TodoRepository{Db: db}
}

func (t *TodoRepository) Create(todo *entities.Todo) error {
	_, err := t.Db.Exec("INSERT INTO todo (id, title, completed) VALUES (?, ?, ?)", todo.ID, todo.Title, todo.Completed)
	if err != nil {
		return err
	}
	return nil
}

func (t *TodoRepository) FindAll() ([]*entities.Todo, error) {
	var todos []*entities.Todo
	err := t.Db.QueryRow("SELECT * FROM todo").Scan(&todos)
	if err != nil {
		return nil, err
	}
	return todos, nil
}
