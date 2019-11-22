package cmd

import (
	"fmt"
	"strings"
	"exercises/task/db"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a todo task",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		fmt.Printf("Added \"%s\" to you list\n", task)
		db.CreateTask(task)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
