package bizlogic

import (
	"tasks-manager-api/database"
	"tasks-manager-api/models"
)

func GetTasksLogic() ([]models.Task, error) {
	rows, err := database.DB.Query("SELECT id, title, description, due_date, status FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.DueDate, &task.Status)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}
