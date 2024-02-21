package cmd

import (
	"github.com/STommydx/FolioForge/app"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the API server",
	Long:  `Start the API server.`,
	Run: func(cmd *cobra.Command, args []string) {
		mainApp := app.New()
		mainApp.Run()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
