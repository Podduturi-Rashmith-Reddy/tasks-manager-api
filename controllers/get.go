package controllers

import (
	"encoding/json"
	"net/http"
	"task-manager/bizlogic"
)

// GetTasks retrieves all tasks
func GetTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := bizlogic.GetTasks()
	if err != nil {
		http.Error(w, "Error fetching tasks", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}
