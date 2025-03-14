package bizlogic

import (
	"tasks-manager-api/database"
	models "tasks-manager-api/models"
)

func UpdateTaskLogic(task models.Task) error {
	query := "UPDATE tasks SET title = ?, description = ?, due_date = ?, status = ? WHERE id = ?"
	_, err := database.DB.Exec(query, task.Title, task.Description, task.DueDate, task.Status, task.ID)
	return err
}
