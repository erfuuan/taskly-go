package tasks

import (
	"fmt"
	"time"

	"github.com/erfuuan/taskly-go/internal/db"
)

type Task struct {
	ID       int
	Title    string
	DueDate  time.Time
	Notified bool
}

// InitTasksTable creates the tasks table if it does not exist
func InitTasksTable() error {
	query := `
	CREATE TABLE IF NOT EXISTS tasks (
		id SERIAL PRIMARY KEY,
		title TEXT NOT NULL,
		duedate TIMESTAMP NOT NULL,
		notified BOOLEAN DEFAULT FALSE
	);`
	_, err := db.DB.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to create tasks table: %w", err)
	}

	fmt.Println("✅ Table 'tasks' is ready")
	return nil
}
