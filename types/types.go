package types

type Task struct {
	Id      int
	Name    string
	Created string
	Done    bool
}

const (
	TASKS_FILENAME = "tasks.csv"
)
