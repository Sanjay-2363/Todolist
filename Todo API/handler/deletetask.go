package handler

import (
	"encoding/json"
	"net/http"
	"todolist/models"
	"todolist/repository"

	"github.com/gorilla/mux"
)

// Delete the task
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	var task models.TodoList
	params := mux.Vars(r)
	if err := repository.Db.First(&task, "task_id=?", params["task_id"]).Error; err != nil {
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	repository.Db.Delete(&task, "task_id=?", params["task_id"])
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "User deleted successfully"})
}
