package handler

import (
	"encoding/json"
	"net/http"
	"todolist/models"
	"todolist/repository"
)

// Display all the completed tasks
func CompletedTask(w http.ResponseWriter, r *http.Request) {
	var completedTasks []models.TodoList
	if err := repository.Db.Where("status", "completed").Find(&completedTasks).Error; err != nil {
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(completedTasks)
}
