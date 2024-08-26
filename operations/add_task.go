package operations

import (
	"time"

	"github.com/ivivanov18/todo-go-cli/types"
	"github.com/ivivanov18/todo-go-cli/utils"
)

func AddTask(name string) {
	task := types.Task{
		Name:    name,
		Created: time.Now(),
		Done:    false,
		Id:      getNewId(),
	}
	tasks := make([]types.Task, 1)
	tasks[0] = task
	utils.WriteDataToFile(types.TASKS_FILENAME, tasks, true)
}

// TODO: logic is not good - if task is deleted & removed from the list, ids for newer
// will be smaller than older tasks
func getNewId() int {
	tasks := GetAllTasks()
	return len(tasks) + 1
}
