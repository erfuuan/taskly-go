package tasks

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/erfuuan/taskly-go/internal/db"
)

func AddTask(title string, dueDate time.Time) error {
	// Check for duplicates
	var exists bool
	err := db.DB.QueryRow(context.Background(), "SELECT EXISTS(SELECT 1 FROM tasks WHERE title=$1)", title).Scan(&exists)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("task title is duplicated")
	}

	_, err = db.DB.Exec(context.Background(), "INSERT INTO tasks(title, duedate, notified) VALUES($1, $2, $3)", title, dueDate, false)
	return err
}

func RunTasks() ([]string, error) {
	now := time.Now()
	thirtyMinutesLater := now.Add(30 * time.Minute)

	rows, err := db.DB.Query(context.Background(), "SELECT id, title FROM tasks WHERE duedate < $1 AND notified = false", thirtyMinutesLater)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []string
	var ids []int

	for rows.Next() {
		var id int
		var title string
		if err := rows.Scan(&id, &title); err != nil {
			return nil, err
		}
		messages = append(messages, fmt.Sprintf("task (%s) due date is about to be reached in 30 minutes.", title))
		ids = append(ids, id)
	}

	// Mark notified
	for _, id := range ids {
		_, _ = db.DB.Exec(context.Background(), "UPDATE tasks SET notified = true WHERE id=$1", id)
	}

	return messages, nil
}
