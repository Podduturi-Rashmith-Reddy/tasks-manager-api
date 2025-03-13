package models

// Task represents a to-do task
type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	DueDate     string `json:"due_date"`
	Status      string `json:"status"` // e.g., "pending", "completed", "deleted"
}
