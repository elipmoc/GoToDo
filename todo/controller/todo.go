package controller

import (
	"GoToDo/todo/domain/service"
	"GoToDo/todo/factory"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//ルーターからの要求を捌くTodoコントローラ
type Todo struct{ todoService service.Todo }

//コントローラの生成
func CreateTodoController() Todo {
	return Todo{todoService: factory.CreateTodoService()}
}

func (t *Todo) Index(c *gin.Context) {
	todoList := t.todoService.GetAll()
	c.HTML(http.StatusOK, "index.html", gin.H{"todoList": todoList})
}

func (t *Todo) New(c *gin.Context) {
	text, isOk := c.GetPostForm("text")
	if isOk {
		t.todoService.New(text)
	}
	c.Redirect(http.StatusMovedPermanently, "")
}

func (t *Todo) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err == nil {
		t.todoService.Delete(uint(id))
	}
	c.Redirect(http.StatusMovedPermanently, "")
}
