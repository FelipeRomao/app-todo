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
	TodoRepository entities.TodoGateway
}

func NewCreateTodo(todoGateway entities.TodoGateway) *CreateTodo {
	return &CreateTodo{TodoRepository: todoGateway}
}

func (c *CreateTodo) Execute(input *TodoInput) (*TodoOutput, error) {
	todo, err := entities.NewTodo(
		input.ID,
		input.Title,
		input.Completed,
	)
	if err != nil {
		return nil, err
	}

	err = c.TodoRepository.Create(todo)
	if err != nil {
		return nil, err
	}

	return &TodoOutput{
		ID:        todo.ID,
		Title:     todo.Title,
		Completed: todo.Completed,
		CreatedAt: time.Now(),
	}, nil

}
