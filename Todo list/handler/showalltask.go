package handler

import (
	"encoding/json"
	"net/http"
	"todolist/models"
	"todolist/repository"
)

// Display all the tasks
func ShowAllTask(w http.ResponseWriter, r *http.Request) {
	var tasks []models.TodoList
	err := repository.Db.Find(&tasks).Error
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}
