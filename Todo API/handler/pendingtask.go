package handler

import (
	"encoding/json"
	"net/http"
	"todolist/models"
	"todolist/repository"
)

// Display all the pending tasks
func Pendingtask(w http.ResponseWriter, r *http.Request) {
	var pendingTasks []models.TodoList
	if err := repository.Db.Where("status", "active").Find(&pendingTasks).Error; err != nil {
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(pendingTasks)
}
