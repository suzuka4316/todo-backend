package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/suzuka4316/todo-backend/routes"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())

	routes.Routes(router)

	log.Fatal(router.Run(":" + port))
}