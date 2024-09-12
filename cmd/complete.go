/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"strconv"

	"github.com/ivivanov18/todo-go-cli/operations"
	"github.com/spf13/cobra"
)

// completeCmd represents the complete command
var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			taskNb, _ := strconv.Atoi(args[0])
			operations.CompleteTask(taskNb)
		}
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)
}
