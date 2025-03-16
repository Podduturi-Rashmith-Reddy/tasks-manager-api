package bizlogic

import (
	"fmt"
	"net/http"
	"tasks-manager-api/database"
)

func DeleteTaskLogic(w http.ResponseWriter, r *http.Request, taskID string) error {
	// Prepare the SQL query to delete a task by its ID
	query := "DELETE FROM tasks WHERE id=?"

	// Ensure taskID is converted to an appropriate type (e.g., int)
	_, err := database.DB.Exec(query, taskID)
	if err != nil {
		return fmt.Errorf("could not delete task: %v", err)
	}

	// Send back a success response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Task deleted successfully"))
	return nil
}
