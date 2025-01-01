package operations

import (
	"fmt"
	"slices"

	"github.com/ivivanov18/todo-go-cli/types"
	"github.com/ivivanov18/todo-go-cli/utils"
)

func CompleteTask(id int) {
	tasks := GetAllTasks()
	taskIdx := slices.IndexFunc(tasks, func(task types.Task) bool {
		return task.Id == id
	})
	if taskIdx != -1 {
		tasks[taskIdx].Done = true
		utils.WriteDataToFile(types.TASKS_FILENAME, tasks)
	} else {
		fmt.Printf("Task with id %d not found.\n", id)
	}
}
