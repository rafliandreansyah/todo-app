package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"todo-app/models"

	"github.com/gin-gonic/gin"
)

var listTodo = []models.Todo{}

func GetTodos(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"todos": listTodo,
	})
}

func CreateTodo(c *gin.Context) {

	todo := models.Todo{}

	headerContentType := c.Request.Header.Get("Content-Type")

	fmt.Println("Content-Type:", headerContentType)

	if headerContentType == "application/json" {
		c.ShouldBindJSON(&todo)
	} else {
		c.ShouldBind(&todo)
	}

	// Generate ID
	todo.ID = strconv.Itoa(len(listTodo) + 1)

	//Append to list todo
	listTodo = append(listTodo, todo)

	c.JSON(http.StatusCreated, gin.H{
		"message": "Todo Created",
		"todo":    todo,
	})

}
