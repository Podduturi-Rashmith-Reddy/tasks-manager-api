package controllers

import (
	"net/http"
	"strconv"

	"tasks-manager-api/bizlogic"
)
// DeleteTask deletes a task.
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	// Extract the task ID from the URL path
	idStr := r.URL.Path[len("/tasks/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	// Call the bizlogic to delete the task
	deleted := bizlogic.DeleteTask(id)
	if !deleted {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	// Return a 404 No Content response
	w.WriteHeader(http.StatusNoContent)
}