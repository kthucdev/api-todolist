package handlers

import (
	"net/http"
	"strconv"

	"chilley.com.todolist/logger"
	"chilley.com.todolist/models"
	"chilley.com.todolist/repository"
	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	repo *repository.TaskRepository
}

func NewTaskHandler(repo *repository.TaskRepository) *TaskHandler {
	logger.InfoLogger.Println("Initializing task handler")
	return &TaskHandler{repo: repo}
}
func (h *TaskHandler) CreateTask(c *gin.Context) {

	var task struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}

	if err := c.ShouldBindJSON(&task); err != nil {
		logger.ErrorLogger.Printf("Invalid request body: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	if task.Title == "" {
		logger.ErrorLogger.Println("Task title is required")
		c.JSON(http.StatusBadRequest, gin.H{"error": "title is required"})
		return
	}

	id := h.repo.Create(models.Task{
		Title:       task.Title,
		Description: task.Description,
		Completed:   false,
	})
	c.JSON(http.StatusCreated, id)
}

func (h *TaskHandler) GetTasks(c *gin.Context) {

	tasks := h.repo.GetAll()

	c.JSON(http.StatusOK, tasks)

}
func (h *TaskHandler) UpdateTask(c *gin.Context) {

	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)

	if err != nil {
		logger.ErrorLogger.Printf("Invalid task ID: %s", idStr)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid task ID"})
		return
	}

	var update struct {
		Completed bool `json:"completed"`
	}

	if err := c.ShouldBindJSON(&update); err != nil {
		logger.ErrorLogger.Printf("Invalid request body: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if task, exists := h.repo.Update(id, update.Completed); exists {
		c.JSON(http.StatusOK, task)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
	}

}

func (h *TaskHandler) DeleteTask(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logger.ErrorLogger.Printf("Invalid task ID: %s", idStr)

		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid task ID"})
		return
	}

	if h.repo.Delete(id) {
		c.Status(http.StatusNoContent)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
	}
}
