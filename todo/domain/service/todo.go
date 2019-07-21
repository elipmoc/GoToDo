package service

import (
	"GoToDo/todo/domain/entity"
	"GoToDo/todo/domain/repository"
)

//Todo serviceオブジェクト
type Todo struct{ todoRepository repository.Todo }

//Todo serviceを生成
func CreateTodo(todoRepository repository.Todo) Todo {
	return Todo{todoRepository}
}

//全てのTodo値オブジェクトを返す
func (t *Todo) GetAll() []entity.Todo {
	return t.todoRepository.GetAll()
}

//新規にTodo値オブジェクトを生成する
func (t *Todo) New(text string) bool {
	return t.todoRepository.Add(text)
}

//指定したTodo値オブジェクトを削除する
func (t *Todo) Delete(id uint) bool {
	return t.todoRepository.Delete(id)
}
