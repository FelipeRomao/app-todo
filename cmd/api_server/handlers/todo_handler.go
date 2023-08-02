package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/FelipeRomao/todo/internal/domain/entities"
	"github.com/FelipeRomao/todo/internal/usecases"
	"github.com/go-chi/chi"
)

func GetTodosHandler(listTodosUseCase *usecases.ListTodo) http.HandlerFunc {
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

func CreateTodoHandler(createTodoUseCase *usecases.CreateTodo) http.HandlerFunc {
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

func RemoveTodoHandler(removeTodoUseCase *usecases.RemoveTodo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		id := chi.URLParam(r, "id")

		err := removeTodoUseCase.Execute(id)
		if err != nil {
			http.Error(w, fmt.Sprintf("Erro ao remover o todo: %v", err), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

func GetOneTodosHandler(getOneTodoUseCase *usecases.GetOneTodo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		id := chi.URLParam(r, "id")

		todo, err := getOneTodoUseCase.Execute(id)
		if err != nil {
			http.Error(w, fmt.Sprintf("Erro ao buscar o todo: %v", err), http.StatusBadRequest)
			return
		}

		response := map[string]string{
			"id":    todo.ID,
			"title": todo.Title,
		}

		responseJSON, err := json.Marshal(response)
		if err != nil {
			http.Error(w, "Erro ao criar a resposta JSON", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(responseJSON)
	}
}

func UpdateTodoHandler(updateTodoUseCase *usecases.UpdateTodo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		id := chi.URLParam(r, "id")

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

		input := &entities.Todo{
			ID:        todo.ID,
			Title:     todo.Title,
			Completed: todo.Completed,
		}

		err = updateTodoUseCase.Execute(id, input)
		if err != nil {
			http.Error(w, fmt.Sprintf("Erro ao atualizar o todo: %v", err), http.StatusBadRequest)
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
