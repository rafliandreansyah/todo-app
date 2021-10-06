package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"todo-app/models"
	"todo-app/utils"

	"github.com/gin-gonic/gin"
)

var listTodo = []models.Todo{}
//var TodoDoc = []models.Todo{
//	{
//		ID: "1",
//		Name: "Create app get weather",
//		Complete: false,
//	},
//	{
//		ID: "2",
//		Name: "Create app post weather",
//		Complete: true,
//	},
//}

// GetTodos godoc
// @Tags todo
// @Summary Show All Todo
// @Description Get json array todo
// @Accept json
// @Accept x-www-form-urlencoded
// @Produce json
// @Success 200 {array} models.Todo
// @Router / [get]
func GetTodos(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"todos": listTodo,
	})
}

// GetTodoByID godoc
// @Tags todo
// @Summary Get todo by ID
// @Description Get object todo by id
// @Accept json
// @Accept x-www-form-urlencoded
// @Produce json
// @Param id path string true "ID"
// @Success 200 {object} models.Todo
// @Failure 400,404 {object} utils.HTTPError
// @Failure 500 {object} utils.HTTPError
// @Failure default {object} utils.HTTPError
// @Router /{id} [get]
func GetTodoByID(c *gin.Context) {
	id := c.Param("id")

	if len(listTodo) < 1 {
		utils.NewError(c, http.StatusNotFound, errors.New("todo not found"))
		return
	}
	for _, value := range listTodo {
		if value.ID == id {
			c.JSON(http.StatusOK, gin.H{
				"todo": value,
			})
			return
		} else {
			utils.NewError(c, http.StatusNotFound, errors.New("todo not found"))
			return
		}
	}
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
