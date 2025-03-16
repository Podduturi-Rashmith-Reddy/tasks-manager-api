package bizlogic

import (
	"log"
	"tasks-manager-api/database"
	"tasks-manager-api/models"
)

func GetTasksLogic() ([]models.Task, error) {
	// Execute query to fetch tasks
	rows, err := database.DB.Query("SELECT id, title, description, due_date FROM tasks")
	if err != nil {
		log.Printf("Error executing query: %v", err)
		return nil, err
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var task models.Task

		// Scan the values into the task struct
		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.DueDate); err != nil {
			log.Printf("Error scanning row: %v", err)
			return nil, err
		}

		// Append the task to the tasks slice
		tasks = append(tasks, task)
	}

	// Check if any errors occurred during row iteration
	if err := rows.Err(); err != nil {
		log.Printf("Error during row iteration: %v", err)
		return nil, err
	}

	return tasks, nil
}
