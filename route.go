package main

import (
	"GoToDo/todo/controller"

	"github.com/gin-gonic/gin"
)

func CreateRouter() *gin.Engine {
	r := gin.Default()
	todoController := controller.CreateTodoController()
	r.LoadHTMLGlob("templates/*")
	r.GET("/", todoController.Index)
	r.POST("/", todoController.New)
	r.POST("/:id", todoController.Delete)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	return r
}
