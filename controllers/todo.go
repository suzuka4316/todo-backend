package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suzuka4316/todo-backend/models"
)

var (
	ErrCantFindTodo = errors.New("todo not found")
)

var todos = []models.Todo {
	{ID: "1", Item: "Grocery shopping", Completed: false},
	{ID: "2", Item: "Eat lunch", Completed: false},
	{ID: "3", Item: "Play PoGo", Completed: false},
}

func GetTodos(context *gin.Context) {
	// convert array to json
	context.IndentedJSON(http.StatusOK, todos)
}

func AddTodo(context *gin.Context) {
	var newTodo models.Todo

	// bind json to array
	if err := context.BindJSON(&newTodo); err != nil {
		return
	}

	todos = append(todos, newTodo)

	context.IndentedJSON(http.StatusCreated, newTodo)
}

func GetTodo(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, todo)
}

func ToggleTodoStatus(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}

	todo.Completed = !todo.Completed

	context.IndentedJSON(http.StatusOK, todo)
}

func getTodoById(id string) (*models.Todo, error) {
	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}
	}

	return nil, ErrCantFindTodo
}