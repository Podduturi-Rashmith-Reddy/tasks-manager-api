package controllers

import (
	"encoding/json"
	"net/http"
	"tasks-manager-api/bizlogic"
	"tasks-manager-api/models"
)

// CreateTask handles the POST request to add a new task
func CreateTask(w http.ResponseWriter, r *http.Request) {
	// Ensure method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var task models.Task

	// Decode JSON request body
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}

	// Call business logic to insert task
	taskID, err := bizlogic.CreateTask(task)
	if err != nil {
		http.Error(w, "Error creating task", http.StatusInternalServerError)
		return
	}

	// Return success response
	task.ID = taskID
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}
