/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/ivivanov18/todo-go-cli/operations"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a new task to the list",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			operations.AddTask(args[0])
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
