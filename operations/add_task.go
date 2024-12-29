package operations

import (
	"time"

	"github.com/ivivanov18/todo-go-cli/types"
	"github.com/ivivanov18/todo-go-cli/utils"
)

func AddTask(name string) {
	task := types.Task{
		Name:    name,
		Created: time.Now().String(),
		Done:    false,
		Id:      getNewId(),
	}
	tasks := make([]types.Task, 1)
	tasks[0] = task
	utils.WriteDataToFile(types.TASKS_FILENAME, tasks)
}

func getNewId() int {
	tasks := GetAllTasks()
	if len(tasks) == 0 {
		return 1
	}

	lastTask := tasks[len(tasks)-1]
	return lastTask.Id + 1
}
