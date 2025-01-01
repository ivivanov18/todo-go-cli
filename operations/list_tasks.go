package operations

import (
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"
	"time"

	"github.com/ivivanov18/todo-go-cli/types"
	"github.com/ivivanov18/todo-go-cli/utils"
	"github.com/mergestat/timediff"
)

func DisplayTasks() {
	tasks := GetAllTasks()

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 0, ' ', tabwriter.Debug)

	fmt.Fprintln(w, "ID\tName\tCreated\tDone\t")
	for _, task := range tasks {
		fmt.Fprintln(w, task.Id, "\t", task.Name, "\t", dateToHumanReadableFormat(task.Created), "\t", task.Done, "\t")
	}
	w.Flush()
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

func dateToHumanReadableFormat(dateInRFC3339Format string) string {
	timeParsed, err := time.Parse(time.RFC3339, dateInRFC3339Format)

	if err != nil {
		fmt.Println("Error parsing date: ", err)
		return dateInRFC3339Format
	}

	return timediff.TimeDiff(timeParsed.Add(-23))
}
