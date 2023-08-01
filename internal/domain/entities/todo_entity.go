package entities

import (
	"github.com/go-playground/validator/v10"
)

type Todo struct {
	ID        string `json:"id" validate:"required"`
	Title     string `json:"title" validate:"required"`
	Completed bool   `json:"completed"`
}

func NewTodo(id string, title string) (*Todo, error) {
	todo := &Todo{
		ID:        id,
		Title:     title,
		Completed: false,
	}

	err := todo.Validate()
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (t *Todo) Validate() error {
	validate := validator.New()
	err := validate.Struct(t)

	return err
}
