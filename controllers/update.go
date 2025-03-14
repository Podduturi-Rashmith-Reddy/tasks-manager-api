package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings" // Add this
	"tasks-manager-api/bizlogic"
	models "tasks-manager-api/models"

	"github.com/gorilla/mux"
)

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	// Trim any whitespace or newlines
	idStr = strings.TrimSpace(idStr)

	fmt.Println("ID from URL:", idStr) // Debug log

	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Error converting ID:", err) // Debug log
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	var task models.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	task.ID = id // Ensure the ID from the URL is set

	if err := bizlogic.UpdateTaskLogic(task); err != nil {
		http.Error(w, "Failed to update task", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(task)
}
