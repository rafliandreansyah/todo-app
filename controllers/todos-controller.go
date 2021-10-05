package controllers

import (
	"net/http"
	"todo-app/models"

	"github.com/gin-gonic/gin"
)

var listTodo = []models.Todo{}

func GetTodos(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"todos": listTodo,
	})
}
