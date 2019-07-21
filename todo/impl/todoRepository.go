package impl

import (
	"GoToDo/todo/domain/entity"
)

//repository.Todoインターフェースの実装

type TodoRepository struct {
	todoList []entity.Todo
	nextId   uint
}

func CreateTodoRepository() *TodoRepository {
	return &TodoRepository{todoList: []entity.Todo{}, nextId: 0}
}

func (t *TodoRepository) GetAll() []entity.Todo {
	return t.todoList
}

func (t *TodoRepository) Add(text string) bool {
	t.todoList = append(t.todoList, entity.Todo{Text: text, Id: t.nextId})
	t.nextId++
	return true
}

func (t *TodoRepository) Delete(id uint) bool {
	newTodoList := []entity.Todo{}
	for _, todo := range t.todoList {
		if todo.Id != id {
			newTodoList = append(newTodoList, todo)
		}
	}
	t.todoList = newTodoList
	return true
}
