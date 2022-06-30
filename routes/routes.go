package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/suzuka4316/todo-backend/controllers"
)

func Routes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/todos", controllers.GetTodos)
	incomingRoutes.GET("/todos/:id", controllers.GetTodo)
	incomingRoutes.PATCH("/todos/:id", controllers.ToggleTodoStatus)
	incomingRoutes.POST("/todos", controllers.AddTodo)
}