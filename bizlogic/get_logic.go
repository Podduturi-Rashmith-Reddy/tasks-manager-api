package bizlogic

import (
	"log"
	"task-manager/database"
	"task-manager/models"
)

// GetTasks retrieves all tasks from the database
func GetTasks() ([]models.Task, error) {
	rows, err := database.DB.Query("SELECT id, title, description, due_date, status FROM tasks")
	if err != nil {
		log.Println("Error fetching tasks:", err)
		return nil, err
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		rows.Scan(&task.ID, &task.Title, &task.Description, &task.DueDate, &task.Status)
		tasks = append(tasks, task)
	}

	return tasks, nil
}
