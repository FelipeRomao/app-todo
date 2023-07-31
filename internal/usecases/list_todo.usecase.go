package usecases

import (
	"github.com/FelipeRomao/todo/internal/domain/entities"
)

type ListTodo struct {
	TodoRepository entities.TodoGateway
}

func GetAllTodo(todoGateway entities.TodoGateway) *ListTodo {
	return &ListTodo{TodoRepository: todoGateway}
}

func (l *ListTodo) Execute() ([]*entities.Todo, error) {
	todos, err := l.TodoRepository.FindAll()

	if err != nil {
		return nil, err
	}

	return todos, nil
}
