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

func GetAllTodoRepository(db *sql.DB) *TodoRepository {
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
	rows, err := t.Db.Query("SELECT * FROM todo")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []*entities.Todo
	for rows.Next() {
		todo := &entities.Todo{}
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Completed); err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return todos, nil
}
