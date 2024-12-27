package handler

import (
	"encoding/json"
	"net/http"
	"todolist/models"
	"todolist/repository"

	"github.com/gorilla/mux"
)

// Mark the task as completed or not completed
func CheckBox(w http.ResponseWriter, r *http.Request) {
	var task models.TodoList
	params := mux.Vars(r)

	if err := repository.Db.First(&task, "task_id = ?", params["task_id"]).Error; err != nil {
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	if task.Status == "completed" {
		task.Status = "active"
	} else {
		task.Status = "completed"
	}

	if err := repository.Db.Save(&task).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Task updated successfully!",
		"status":  task.Status,
	})
}
