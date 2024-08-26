/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"strconv"

	"github.com/ivivanov18/todo-go-cli/operations"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes a task from the list",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			taskNb, _ := strconv.Atoi(args[0])
			operations.DeleteTask(taskNb)
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
