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

// CreateTodo godoc
// @Tags todo
// @Summary Create todo
// @Description Post object todo with name and status complete
// @Accept json
// @Accept x-www-form-urlencoded
// @Produce json
// @Param todo body models.Todo true "Add Todo"
// @Success 200 {object} models.Todo
// @Failure 400 {object} utils.HTTPError
// @Failure 500 {object} utils.HTTPError
// @Failure default {object} utils.HTTPError
// @Router /create-todo [post]
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

	c.JSON(http.StatusCreated, todo)

}

func EditTodo(c *gin.Context){

	id := c.Param("id")
	todo := models.Todo{}

	headerContentType := c.Request.Header.Get("Content-Type")

	if headerContentType == "application/json" {
		c.ShouldBindJSON(&todo)
	} else{
		c.ShouldBind(&todo)
	}

	for i, v := range listTodo{
		if v.ID == id {
			listTodo[i].Name = todo.Name
			listTodo[i].Complete = todo.Complete
			c.JSON(http.StatusOK, listTodo)
			return
		}
	}

	utils.NewError(c, http.StatusNotFound, errors.New(fmt.Sprintf("todo with id %v not found", id)))

}

func DeleteTodoById(c *gin.Context){

	id := c.Param("id")

	for i, v := range listTodo{
		if v.ID == id {
			listTodo = append(listTodo[:i], listTodo[i+1:]...)
			c.JSON(http.StatusOK, listTodo)
			return
		}
	}

	utils.NewError(c, http.StatusNotFound, errors.New(fmt.Sprintf("todo with id %v not found", id)))

}
