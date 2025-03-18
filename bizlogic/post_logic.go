package bizlogic

import (
	"fmt"
	"log"
	"task-manager/database"
	"task-manager/models"
)

// CreateTask inserts a new task into the database
func CreateTask(task models.Task) (int, error) {
	query := "INSERT INTO tasks (title, description, due_date, status) VALUES (?, ?, ?, ?)"

	stmt, err := database.DB.Prepare(query)
	if err != nil {
		log.Println("Error preparing insert statement:", err)
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(task.Title, task.Description, task.DueDate, task.Status)
	if err != nil {
		log.Println("Error executing insert statement:", err)
		return 0, err
	}

	// Get last inserted ID
	taskID, err := result.LastInsertId()
	if err != nil {
		log.Println("Error fetching last insert ID:", err)
		return 0, err
	}

	fmt.Println("New task inserted with ID:", taskID)
	return int(taskID), nil
}
