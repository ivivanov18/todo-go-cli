package operations

import (
	"fmt"
	"slices"

	"github.com/ivivanov18/todo-go-cli/types"
	"github.com/ivivanov18/todo-go-cli/utils"
)

func filterTask(id int) []types.Task {
	tasks := GetAllTasks()

	remainingTasks := slices.DeleteFunc(tasks, func(task types.Task) bool {
		return task.Id == id
	})

	if len(remainingTasks) == len(tasks) {
		return nil
	} else {
		return remainingTasks
	}
}

func DeleteTask(id int) {
	tasks := filterTask(id)

	if tasks == nil {
		fmt.Printf("Task with id %d not found.\n", id)
		return
	} else {
		utils.WriteDataToFile(types.TASKS_FILENAME, tasks)
	}
}
