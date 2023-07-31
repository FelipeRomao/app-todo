package entities

import "errors"

type Todo struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func NewTodo(id string, title string, completed bool) (*Todo, error) {
	todo := &Todo{
		ID:        id,
		Title:     title,
		Completed: completed,
	}

	err := todo.Validate()
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (t *Todo) Validate() error {
	if t.ID == "" {
		return errors.New("ID cannot be empty")
	}

	if t.Title == "" {
		return errors.New("Title cannot be empty")
	}

	return nil
}
