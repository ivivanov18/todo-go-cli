package operations

import (
	"fmt"
	"strconv"

	"github.com/ivivanov18/todo-go-cli/types"
	"github.com/ivivanov18/todo-go-cli/utils"
)

func DisplayTasks() {
	tasks := GetAllTasks()

	for _, task := range tasks {
		fmt.Printf("%d - %s | %s | %v\n", task.Id, task.Name, task.Created, task.Done)
	}
}

func GetAllTasks() []types.Task {
	records := utils.ReadFile(types.TASKS_FILENAME)
	tasks := make([]types.Task, len(records))

	for i, record := range records {
		done, _ := strconv.ParseBool(record[3])
		id, _ := strconv.Atoi(record[0])
		tasks[i] = types.Task{
			Id:      id,
			Name:    record[1],
			Created: record[2],
			Done:    done,
		}
	}

	return tasks
}
