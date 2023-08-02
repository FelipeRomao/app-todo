package usecases

import "github.com/FelipeRomao/todo/internal/domain/entities"

type UpdateTodo struct {
	TodoGateway entities.TodoGateway
}

func NewUpdateTodo(todoGateway entities.TodoGateway) *UpdateTodo {
	return &UpdateTodo{TodoGateway: todoGateway}
}

func (u *UpdateTodo) Execute(id string, todo *entities.Todo) error {
	err := u.TodoGateway.Update(id, todo)

	if err != nil {
		return err
	}

	return nil
}
