package task

import (
	"fmt"
	"time"
)

// Task represents a to-do item
type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Done        bool      `json:"done"`
	CreatedAt   time.Time `json:"created_at"`
}

// MarkAsDone marks the task as completed
func (t *Task) MarkAsDone() {
	t.Done = true
}

// Summary returns a formatted task summary
func (t Task) Summary() string {
	status := "[ ]"
	if t.Done {
		status = "[✓]"
	}
	return fmt.Sprintf("%s %d: %s", status, t.ID, t.Title)
}
