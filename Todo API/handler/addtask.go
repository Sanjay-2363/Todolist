package handler

import (
	"encoding/json"
	"net/http"
	"todolist/models"
	"todolist/repository"
)

// Add the task
func AddTask(w http.ResponseWriter, r *http.Request) {
	var task models.TodoList
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if task.Task_name == "" {
		json.NewEncoder(w).Encode(map[string]string{"error": "Task Name field cannot be empty!"})
		return
	}

	err := repository.Db.Create(&task).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Task added successfully!"})
}
