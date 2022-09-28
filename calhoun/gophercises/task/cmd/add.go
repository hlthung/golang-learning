package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to your task list.",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO
		fmt.Printf("added call")
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
