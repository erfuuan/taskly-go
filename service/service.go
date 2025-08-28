package service

import (
	"fmt"
	"time"

	"github.com/erfuuan/taskly-go/internal/db"
)

func RunTasks() ([]string, error) {
	query := `SELECT id, title FROM tasks WHERE duedate < $1 AND notified = false`
	rows, err := db.DB.Query(query, time.Now())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var msgs []string
	for rows.Next() {
		var id int
		var title string
		if err := rows.Scan(&id, &title); err != nil {
			return nil, err
		}
		msgs = append(msgs, fmt.Sprintf("Task %d: %s", id, title))
	}

	// Optional: mark tasks as notified
	_, _ = db.DB.Exec(`UPDATE tasks SET notified = true WHERE duedate < $1 AND notified = false`, time.Now())

	return msgs, nil
}

func AddTask(title string, dueDate time.Time) error {
	query := `INSERT INTO tasks (title, duedate, notified) VALUES ($1, $2, false)`
	_, err := db.DB.Exec(query, title, dueDate)
	if err != nil {
		return fmt.Errorf("failed to add task: %w", err)
	}
	fmt.Println("✅ Task added:", title)
	return nil
}
