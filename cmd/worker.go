package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var workerCmd = &cobra.Command{
	Use:   "worker",
	Short: "Start the worker",
	Long:  `Start the worker.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("worker called")
	},
}

func init() {
	rootCmd.AddCommand(workerCmd)
}
