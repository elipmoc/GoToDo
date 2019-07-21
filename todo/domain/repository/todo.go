package repository

import "GoToDo/todo/domain/entity"

//Todo　Repositoryインターフェース
type Todo interface {
	//Todo
	GetAll() []entity.Todo
	Add(text string) bool
	Delete(id uint) bool
}
