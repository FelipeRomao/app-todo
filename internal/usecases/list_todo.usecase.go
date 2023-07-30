package usecases

import (
	"github.com/FelipeRomao/todo/internal/domain/entities"
)

type ListTodo struct {
	TodoGateway entities.TodoGateway
}

func (l *ListTodo) Execute() ([]*entities.Todo, error) {
	todos, err := l.TodoGateway.FindAll()

	if err != nil {
		return nil, err
	}

	return todos, nil
}
