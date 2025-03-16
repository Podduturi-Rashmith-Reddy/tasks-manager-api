package bizlogic

import (
	"database/sql"
	"log"
	"tasks-manager-api/database"
)

// DeleteTask removes a task from the database
func DeleteTask(id int) error {
	query := "DELETE FROM tasks WHERE id = ?"

	stmt, err := database.DB.Prepare(query)
	if err != nil {
		log.Printf("Error preparing delete statement: %v", err)
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(id)
	if err != nil {
		log.Printf("Error executing delete statement: %v", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Error getting rows affected: %v", err)
		return err
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows // Return an error if no task was deleted
	}

	return nil
}
