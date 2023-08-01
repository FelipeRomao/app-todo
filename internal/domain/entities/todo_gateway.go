package entities

type TodoGateway interface {
	Create(todo *Todo) error
	FindAll() ([]*Todo, error)
	Remove(id string) error
	FindOne(id string) (*Todo, error)
}
