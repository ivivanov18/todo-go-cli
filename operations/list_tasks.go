package operations

import (
	"fmt"
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
