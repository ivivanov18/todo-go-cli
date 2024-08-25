package operations

import (
	"github.com/ivivanov18/todo-go-cli/types"
	"github.com/ivivanov18/todo-go-cli/utils"
)

func AddTask(task types.Task) {
	tasks := make([]types.Task, 1)
	tasks[0] = task
	utils.WriteDataToFile(types.TASKS_FILENAME, tasks, true)
}
