package types

type Task struct {
	Name    string
	DueDate string
	Done    bool
}

const (
	TASKS_FILENAME = "tasks.csv"
)
