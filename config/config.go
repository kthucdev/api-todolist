package config

import (
	"chilley.com.todolist/handlers"
	"chilley.com.todolist/logger"
	"chilley.com.todolist/repository"
	"chilley.com.todolist/routes"
	"github.com/gin-gonic/gin"
)

func Init() {
	// Initialize logger
	logger.InitLogger()

	// Initialize repository
	taskRepo := repository.NewTaskRepository()

	// Initialize handler
	taskHandler := handlers.NewTaskHandler(taskRepo)

	// Setup router
	r := gin.Default()
	routes.SetupRoutes(r, taskHandler)

	// Start server
	r.Run(":8080")

	// Start server
	logger.InfoLogger.Println("Starting server on :8080")
}
