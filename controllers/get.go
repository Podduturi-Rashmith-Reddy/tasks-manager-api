package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"tasks-manager-api/bizlogic"
)

func GetTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := bizlogic.GetTasksLogic()
	if err != nil {
		// Log the error for better debugging
		log.Printf("Error fetching tasks: %v", err)
		http.Error(w, "Error fetching tasks", http.StatusInternalServerError)
		return
	}

	if len(tasks) == 0 {
		// If no tasks found, return 204 No Content
		w.WriteHeader(http.StatusNoContent)
		return
	}

	// Set Content-Type and encode the tasks into JSON
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(tasks); err != nil {
		// Handle error during JSON encoding
		log.Printf("Error encoding tasks to JSON: %v", err)
		http.Error(w, "Error encoding tasks", http.StatusInternalServerError)
		return
	}
}
