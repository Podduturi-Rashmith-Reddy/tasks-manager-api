package bizlogic

import (
	"fmt"
	"tasks-manager-api/database"
	"tasks-manager-api/models"
)

func CreateTaskLogic(task models.Task) (int, error) {
	query := "INSERT INTO task(title, description, due_date, status)VALUES(?,?,?,?)RETURNING id"
	var taskID int
	if task.Status == "" {
		task.Status = "pending"
	}
	// QueryRow is efficient for single-row operations like inserting one task and getting its ID.
	err := database.DB.QueryRow(query, task.Title, task.Description, task.DueDate, task.Status).Scan(&taskID)
	if err != nil {
		return 0, fmt.Errorf("failed to insert task: %v", err)
	}
	return taskID, nil

}
