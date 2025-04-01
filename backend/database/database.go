package database

type TodoRepository interface {
	AddTodo(title string, priority *string, timeToComplete *int) map[string]interface{}
	GetTodos(completed *bool) []map[string]interface{}
	ToggleTodo(id uint)
	DeleteTodo(id uint)
	UpdateTodo(id uint, title string, priority *string, timeToComplete *int) (map[string]interface{}, error)
}
