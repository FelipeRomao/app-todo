package usecases

import "github.com/FelipeRomao/todo/internal/domain/entities"

type GetOneTodo struct {
	TodoGateway entities.TodoGateway
}

func NewGetOneTodo(todoGateway entities.TodoGateway) *GetOneTodo {
	return &GetOneTodo{TodoGateway: todoGateway}
}

func (g *GetOneTodo) Execute(id string) (*entities.Todo, error) {
	todo, err := g.TodoGateway.FindOne(id)

	if err != nil {
		return nil, err
	}

	return todo, nil
}
