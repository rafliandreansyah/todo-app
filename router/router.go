package router

import (
	"todo-app/controllers"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()

	router.GET("/", controllers.GetTodos)
	router.GET("/:id", controllers.GetTodoByID)
	router.POST("/create-todo", controllers.CreateTodo)
	router.PUT("/edit-todo/:id", controllers.EditTodo)
	router.DELETE("/delete/:id", controllers.DeleteTodoById)

	return router
}
