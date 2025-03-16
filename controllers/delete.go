package controllers

import (
	"net/http"
	"tasks-manager-api/bizlogic"

	"github.com/gorilla/mux" // For URL parameter handling
)

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	// Retrieve the task ID from the URL parameter
	vars := mux.Vars(r)
	taskID := vars["id"]

	// Make sure taskID is not empty
	if taskID == "" {
		http.Error(w, "Task ID is required", http.StatusBadRequest)
		return
	}

	// Call the business logic function with the task ID
	err := bizlogic.DeleteTaskLogic(w, r, taskID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
