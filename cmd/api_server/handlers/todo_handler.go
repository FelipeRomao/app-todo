package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/FelipeRomao/todo/internal/domain/entities"
	"github.com/FelipeRomao/todo/internal/usecases"
)

func GetTodos(listTodosUseCase *usecases.ListTodo) http.HandlerFunc {
	todos, err := listTodosUseCase.Execute()
	if err != nil {
		panic(err)
	}

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		err := json.NewEncoder(w).Encode(todos)

		if err != nil {
			http.Error(w, fmt.Sprintf("Erro ao codificar os todos: %v", err), http.StatusInternalServerError)
			return
		}
	}
}

func CreateTodo(createTodoUseCase *usecases.CreateTodo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Erro ao ler o corpo da requisição", http.StatusInternalServerError)
			return
		}

		var todo entities.Todo

		err = json.Unmarshal(body, &todo)
		if err != nil {
			http.Error(w, "Erro na decodificação do JSON", http.StatusBadRequest)
			return
		}

		input := &usecases.TodoInput{
			ID:    todo.ID,
			Title: todo.Title,
		}

		output, err := createTodoUseCase.Execute(input)
		if err != nil {
			http.Error(w, fmt.Sprintf("Erro ao criar o todo: %v", err), http.StatusBadRequest)
			return
		}

		response := map[string]string{
			"id":        output.ID,
			"title":     output.Title,
			"createdAt": output.CreatedAt.String(),
		}

		responseJSON, err := json.Marshal(response)
		if err != nil {
			http.Error(w, "Erro ao criar a resposta JSON", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write(responseJSON)
	}
}
