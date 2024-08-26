/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"

	"github.com/ivivanov18/todo-go-cli/operations"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		operations.DisplayTasks()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
