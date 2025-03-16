package controllers

import (
	"encoding/json"
	"net/http"
	"tasks-manager-api/bizlogic"
)

func GetTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := bizlogic.GetTasksLogic()
	if err != nil {
		http.Error(w, "Error fetching tasks", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}
