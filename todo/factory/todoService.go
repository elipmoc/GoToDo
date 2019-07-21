package factory

import (
	"GoToDo/todo/domain/service"
	"GoToDo/todo/impl"
)

//todo serviceを生成
func CreateTodoService() service.Todo {
	return service.CreateTodo(impl.CreateTodoRepository())
}
