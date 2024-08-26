package utils

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/ivivanov18/todo-go-cli/types"
)

func ReadFile(fileName string) [][]string {
	if !fileExists(fileName) {
		return nil
	}

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
	var file *os.File
	var err error

	// Add line
	if len(tasks) == 1 {
		file, err = os.OpenFile(fileName, os.O_APPEND|os.O_CREATE, 0644)
	} else {
		file, err = os.OpenFile(fileName, os.O_CREATE|os.O_TRUNC, 0644)
	}

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

func fileExists(fileName string) bool {
	_, err := os.Stat(fileName)
	if err == nil {
		// File exists
		return true
	}
	if os.IsNotExist(err) {
		// File does not exist
		return false
	}
	// Some other error occurred
	fmt.Println("Error checking file:", err)
	return false
}
