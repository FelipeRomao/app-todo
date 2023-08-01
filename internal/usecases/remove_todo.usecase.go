package usecases

import "github.com/FelipeRomao/todo/internal/domain/entities"

type RemoveTodo struct {
	TodoGateway entities.TodoGateway
}

func NewRemoveTodo(todoGateway entities.TodoGateway) *RemoveTodo {
	return &RemoveTodo{TodoGateway: todoGateway}
}

func (r *RemoveTodo) Execute(id string) error {
	err := r.TodoGateway.Remove(id)

	if err != nil {
		return err
	}

	return nil
}
