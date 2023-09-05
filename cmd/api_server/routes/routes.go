package routes

import (
	"database/sql"

	"github.com/FelipeRomao/todo/cmd/api_server/handlers"
	"github.com/FelipeRomao/todo/internal/infra/database"
	"github.com/FelipeRomao/todo/internal/usecases"
	"github.com/go-chi/chi"
	_ "github.com/lib/pq"
)

func SetRoutes(r chi.Router) {
	connStr := "user=root password=postgres dbname=todo_app sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	todoRepository := database.NewTodoRepository(db)

	listTodo := usecases.NewGetAllTodo(todoRepository)
	getOneTodo := usecases.NewGetOneTodo(todoRepository)
	createTodo := usecases.NewCreateTodo(todoRepository)
	removeTodo := usecases.NewRemoveTodo(todoRepository)
	updateTodo := usecases.NewUpdateTodo(todoRepository)

	r.Get("/api/todo", handlers.GetTodosHandler(listTodo))
	r.Get("/api/todo/{id}", handlers.GetOneTodosHandler(getOneTodo))
	r.Post("/api/todo", handlers.CreateTodoHandler(createTodo))
	r.Delete("/api/todo/{id}", handlers.RemoveTodoHandler(removeTodo))
	r.Put("/api/todo/{id}", handlers.UpdateTodoHandler(updateTodo))
}
