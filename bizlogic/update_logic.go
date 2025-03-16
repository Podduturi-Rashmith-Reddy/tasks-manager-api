package bizlogic

import (
	"fmt"
	"log"
	"tasks-manager-api/database"
	"tasks-manager-api/models"
)

// UpdateTask updates an existing task in the database
func UpdateTaskLogic(task models.Task) error {
	query := "UPDATE tasks SET title = ?, description = ?, due_date = ?, status = ? WHERE id = ?"

	stmt, err := database.DB.Prepare(query)
	if err != nil {
		log.Printf("Error preparing update statement: %v", err)
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(task.Title, task.Description, task.DueDate, task.Status, task.ID)
	if err != nil {
		log.Printf("Error executing update statement: %v", err)
		return err
	}

	// Check if a row was actually updated
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Error getting rows affected: %v", err)
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no task found with ID %d", task.ID)
	}

	return nil
}
