package operations

import (
	"slices"

	"github.com/ivivanov18/todo-go-cli/types"
	"github.com/ivivanov18/todo-go-cli/utils"
)

func filterTask(id int) []types.Task {
	tasks := GetAllTasks()

	remainingTasks := slices.DeleteFunc(tasks, func(task types.Task) bool {
		return task.Id == id
	})

	return remainingTasks
}

func DeleteTask(id int) {
	tasks := filterTask(id)

	utils.WriteDataToFile(types.TASKS_FILENAME, tasks)
}
