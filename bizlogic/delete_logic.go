package bizlogic

import (
	"database/sql"
	"log"
	"task-manager/database"
)

// DeleteTask removes a task from the database
func DeleteTask(id int) error {
	query := "DELETE FROM tasks WHERE id = ?"

	stmt, err := database.DB.Prepare(query)
	if err != nil {
		log.Println("Error preparing delete statement:", err)
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(id)
	if err != nil {
		log.Println("Error executing delete statement:", err)
		return err
	}

	// Check if a row was actually deleted
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error getting rows affected:", err)
		return err
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows // Return an error if no task was deleted
	}

	return nil
}
