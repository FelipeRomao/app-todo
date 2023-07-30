package main

import (
	"database/sql"
	"fmt"

	"github.com/FelipeRomao/todo/internal/infra/database"
	"github.com/FelipeRomao/todo/internal/usecases"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "db.sqlite3")
	if err != nil {
		panic(err)
	}

	todoRepository := database.NewTodoRepository(db)
	createTodo := usecases.NewCreateTodo(todoRepository)

	input := &usecases.TodoInput{
		ID:        "1",
		Title:     "My first todo",
		Completed: false,
	}

	output, err := createTodo.Execute(input)
	if err != nil {
		panic(err)
	}

	fmt.Println(output)
}
