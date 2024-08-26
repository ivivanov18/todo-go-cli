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

func WriteDataToFile(fileName string, tasks []types.Task) {
	for _, task := range tasks {
		fmt.Println(task)
	}

	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	records := make([][]string, len(tasks))
	for i, task := range tasks {
		records[i] = []string{
			fmt.Sprintf("%d", task.Id),
			task.Name,
			fmt.Sprintf("%s", task.Created),
			fmt.Sprintf("%t", task.Done),
		}
		if err := writer.Write(records[i]); err != nil {
			fmt.Println("Error: ", err)
		}
	}

	if err := writer.Error(); err != nil {
		fmt.Println("Error: ", err)
	}
}
