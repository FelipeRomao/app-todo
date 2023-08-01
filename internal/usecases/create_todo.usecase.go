package usecases

import (
	"time"

	"github.com/FelipeRomao/todo/internal/domain/entities"
)

type TodoInput struct {
	ID        string
	Title     string
	Completed bool
}

type TodoOutput struct {
	ID        string
	Title     string
	Completed bool
	CreatedAt time.Time
}

type CreateTodo struct {
	TodoGateway entities.TodoGateway
}

func NewCreateTodo(todoGateway entities.TodoGateway) *CreateTodo {
	return &CreateTodo{TodoGateway: todoGateway}
}

func (c *CreateTodo) Execute(input *TodoInput) (*TodoOutput, error) {
	todo, err := entities.NewTodo(
		input.ID,
		input.Title,
	)
	if err != nil {
		return nil, err
	}

	err = c.TodoGateway.Create(todo)
	if err != nil {
		return nil, err
	}

	return &TodoOutput{
		ID:        todo.ID,
		Title:     todo.Title,
		CreatedAt: time.Now(),
	}, nil

}
