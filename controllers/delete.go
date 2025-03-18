package controllers

import (
	"net/http"
	"strconv"

	"task-manager/bizlogic"

	"github.com/gorilla/mux"
)

// DeleteTask deletes a task.
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	// Extract the task ID from the URL path using mux
	vars := mux.Vars(r)
	idStr := vars["id"]

	// Convert ID from string to integer
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	// Call the bizlogic to delete the task
	err = bizlogic.DeleteTask(id)
	if err != nil {
		http.Error(w, "Task not found or already deleted", http.StatusNotFound)
		return
	}

	// Return a 204 No Content response (successful deletion)
	w.WriteHeader(http.StatusNoContent)
}
