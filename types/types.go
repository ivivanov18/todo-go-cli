package types

import "time"

type Task struct {
	Id      int
	Name    string
	Created time.Time
	Done    bool
}

const (
	TASKS_FILENAME = "tasks.csv"
)
