package bizlogic

import (
	"log"
	"tasks-manager-api/database"
	"tasks-manager-api/models"
)

// GetTasks retrieves all tasks from the database
func GetTasksLogic() ([]models.Task, error) {
	query := "SELECT id, title, description, due_date, status FROM tasks"

	rows, err := database.DB.Query(query)
	if err != nil {
		log.Printf("Error fetching tasks: %v", err)
		return nil, err
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.DueDate, &task.Status); err != nil {
			log.Printf("Error scanning task row: %v", err)
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}
