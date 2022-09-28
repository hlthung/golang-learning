package cmd

import (
	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Marks a task as complete",
}

func init() {
	RootCmd.AddCommand(doCmd)
}
