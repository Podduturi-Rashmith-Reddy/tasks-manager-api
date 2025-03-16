package bizlogic

import (
	"log"
	"tasks-manager-api/database"
	"tasks-manager-api/models"
)

// CreateTask inserts a new task into the database
func CreateTaskLogic(task models.Task) (int, error) {
	query := "INSERT INTO tasks (title, description, due_date, status) VALUES (?, ?, ?, ?)"

	stmt, err := database.DB.Prepare(query)
	if err != nil {
		log.Printf("Error preparing insert statement: %v", err)
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(task.Title, task.Description, task.DueDate, task.Status)
	if err != nil {
		log.Printf("Error executing insert statement: %v", err)
		return 0, err
	}

	// Get last inserted ID
	taskID, err := result.LastInsertId()
	if err != nil {
		log.Printf("Error fetching last insert ID: %v", err)
		return 0, err
	}

	return int(taskID), nil
}
