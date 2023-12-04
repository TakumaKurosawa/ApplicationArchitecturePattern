package domain

type TodoRepo interface {
	GetTodoByID(id string) (*Todo, error)
	CreateTodo(dto *Todo) (*Todo, error)
	UpdateTodoStatus(dto *Todo) (*Todo, error)
}
