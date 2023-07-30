package entities

type TodoGateway interface {
	Create(todo *Todo) error
	FindAll() ([]*Todo, error)
}
