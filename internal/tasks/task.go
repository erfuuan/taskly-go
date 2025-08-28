package tasks

import "time"

type Task struct {
	ID       int
	Title    string
	DueDate  time.Time
	Notified bool
}
