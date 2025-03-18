package controllers

import (
	"encoding/json"
	"net/http"
	"task-manager/bizlogic"
	"task-manager/models"
)

// CreateTask handles task creation
func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}

	taskID, err := bizlogic.CreateTask(task)
	if err != nil {
		http.Error(w, "Error creating task", http.StatusInternalServerError)
		return
	}

	task.ID = taskID
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}
