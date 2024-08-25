package operations

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ivivanov18/todo-go-cli/types"
	"github.com/ivivanov18/todo-go-cli/utils"
)

func ListTasks() {
	records := utils.ReadFile(types.TASKS_FILENAME)

	for i, record := range records {
		task := strings.Split(record[0], ";")
		fmt.Printf("%d - %s | %s | %s\n", i+1, task[0], task[1], task[2])
	}
}

func DisplayTasks() {
	tasks := GetAllTasks()

	for i, task := range tasks {
		fmt.Printf("%d - %s | %s | %v\n", i+1, task.Name, task.DueDate, task.Done)
	}
}

func GetAllTasks() []types.Task {
	records := utils.ReadFile(types.TASKS_FILENAME)
	tasks := make([]types.Task, len(records))

	for i, record := range records {
		task := strings.Split(record[0], ";")
		done, _ := strconv.ParseBool(task[2])
		tasks[i] = types.Task{
			Name:    task[0],
			DueDate: task[1],
			Done:    done,
		}
	}

	return tasks
}
