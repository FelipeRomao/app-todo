package routes

import (
	"database/sql"

	"github.com/FelipeRomao/todo/cmd/api_server/handlers"
	"github.com/FelipeRomao/todo/internal/infra/database"
	"github.com/FelipeRomao/todo/internal/usecases"
	"github.com/go-chi/chi"
	_ "github.com/mattn/go-sqlite3"
)

func SetRoutes(r chi.Router) {
	db, err := sql.Open("sqlite3", "db.sqlite3")
	if err != nil {
		panic(err)
	}

	todoRepository := database.NewTodoRepository(db)

	listTodos := usecases.GetAllTodo(todoRepository)
	createTodo := usecases.NewCreateTodo(todoRepository)

	r.Get("/api/todos", handlers.GetTodos(listTodos))
	r.Post("/api/todo", handlers.CreateTodo(createTodo))
}
