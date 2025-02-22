package routes

import (
	"chilley.com.todolist/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, taskHandler *handlers.TaskHandler) {
	r.GET("/tasks", taskHandler.GetTasks)
	r.POST("/tasks", taskHandler.CreateTask)
	r.PUT("/tasks/:id", taskHandler.UpdateTask)
	r.DELETE("/tasks/:id", taskHandler.DeleteTask)
}
