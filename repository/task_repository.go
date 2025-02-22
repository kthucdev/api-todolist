package repository

import (
	"sync"

	"chilley.com.todolist/models"
)

type TaskRepository struct {
	sync.RWMutex
	tasks  map[int]models.Task
	nextID int
}

func NewTaskRepository() *TaskRepository {
	return &TaskRepository{
		tasks:  make(map[int]models.Task),
		nextID: 1,
	}
}

func (r *TaskRepository) Create(task models.Task) int {
	r.Lock()
	defer r.Unlock()

	task.ID = r.nextID
	r.nextID++
	task.Completed = false
	r.tasks[task.ID] = task

	return task.ID
}

func (r *TaskRepository) GetAll() []models.Task {
	r.RLock()
	defer r.RUnlock()

	tasks := make([]models.Task, 0, len(r.tasks))
	for _, task := range r.tasks {
		tasks = append(tasks, task)
	}
	return tasks
}

func (r *TaskRepository) Update(id int, completed bool) (models.Task, bool) {
	r.Lock()
	defer r.Unlock()

	if task, exists := r.tasks[id]; exists {
		task.Completed = completed
		r.tasks[id] = task
		return task, true
	}
	return models.Task{}, false
}

func (r *TaskRepository) Delete(id int) bool {
	r.Lock()
	defer r.Unlock()

	if _, exists := r.tasks[id]; exists {
		delete(r.tasks, id)
		return true
	}
	return false
}
