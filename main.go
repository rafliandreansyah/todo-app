package main

import (
	"todo-app/router"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware

	_ "todo-app/docs"
)

// @title Todo API With Swagger Documentation
// @version 1.0
// @description This is API for CRUD todo app

// @host localhost:8080

func main() {
	r := router.Router()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")
}
