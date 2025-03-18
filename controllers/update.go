package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"task-manager/bizlogic"
	"task-manager/models"

	"github.com/gorilla/mux"
)

// UpdateTask updates an existing task.
func UpdateTask(w http.ResponseWriter, r *http.Request) {
	// Extract task ID from URL
	vars := mux.Vars(r)
	idStr := vars["id"]

	// Convert ID from string to integer
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	var updatedTask models.Task

	// Decode JSON request body
	err = json.NewDecoder(r.Body).Decode(&updatedTask)
	if err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}

	// Set task ID from URL
	updatedTask.ID = id

	// Call Business Logic to update task
	err = bizlogic.UpdateTask(updatedTask)
	if err != nil {
		log.Printf("Error updating task: %v", err) // Log the error details
		http.Error(w, "Error updating task: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Return success response with updated task data
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedTask)
}
