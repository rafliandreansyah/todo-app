package router

import (
	"todo-app/controllers"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()

	router.GET("/", controllers.GetTodos)

	return router
}
