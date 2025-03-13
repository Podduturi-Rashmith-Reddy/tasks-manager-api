package bizlogic

import "tasks-manager-api/model"

// tasks is an in-memory map to store tasks.
var tasks = make(map[int]model.Task)

// DeleteTask deletes a task by its ID.
// Returns true if the task was deleted, false if the task was not found.
func DeleteTask(id int) bool {
	if _, exists := tasks[id]; !exists {
		return false // Task not found
	}
	delete(tasks, id)
	return true // Task deleted successfully
}