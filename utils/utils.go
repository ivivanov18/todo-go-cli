package utils

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/ivivanov18/todo-go-cli/types"
)

func ReadFile(fileName string) [][]string {
	file, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	return records
}

func WriteDataToFile(fileName string, tasks []types.Task, new bool) {
	var access int = os.O_WRONLY
	if new {
		access = os.O_APPEND
	}
	file, err := os.OpenFile(fileName, access, 0666)
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	records := make([][]string, len(tasks))
	for i, task := range tasks {
		records[i] = []string{task.Name, task.DueDate, fmt.Sprintf("%t", task.Done)}
		if err := writer.Write(records[i]); err != nil {
			fmt.Println("Error: ", err)
		}
	}

	if err := writer.Error(); err != nil {
		fmt.Println("Error: ", err)
	}
}
