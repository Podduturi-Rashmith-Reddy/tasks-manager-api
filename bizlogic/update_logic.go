package bizlogic

import (
	"fmt"
	"log"
	"task-manager/database"
	"task-manager/models"
)

// UpdateTask modifies an existing task in the database
func UpdateTask(task models.Task) error {
	// Validate that task fields are not empty (example validation, modify as needed)
	if task.Title == "" || task.Description == "" || task.DueDate == "" || task.Status == "" {
		return fmt.Errorf("all fields must be provided: Title, Description, DueDate, and Status")
	}

	// Prepare SQL query
	query := "UPDATE tasks SET title = ?, description = ?, due_date = ?, status = ? WHERE id = ?"

	stmt, err := database.DB.Prepare(query)
	if err != nil {
		log.Printf("Error preparing update statement: %v, Query: %v", err, query)
		return fmt.Errorf("failed to prepare the update statement: %v", err)
	}
	defer stmt.Close()

	// Execute SQL query
	result, err := stmt.Exec(task.Title, task.Description, task.DueDate, task.Status, task.ID)
	if err != nil {
		log.Printf("Error executing update statement: %v, Query: %v", err, query)
		return fmt.Errorf("failed to execute the update statement: %v", err)
	}

	// Check if any row was affected (task with given ID)
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Error getting rows affected: %v", err)
		return fmt.Errorf("failed to retrieve rows affected: %v", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no task found with ID %d", task.ID)
	}

	return nil
}
