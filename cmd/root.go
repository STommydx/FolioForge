package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "FolioForge",
	Short: "API toolkit designed to streamline the management and presentation of your professional portfolio",
	Long:  `FolioForge is an API toolkit designed to streamline the management and presentation of your professional portfolio. With FolioForge, you can easily update your experience, projects, and skills through its intuitive CRUD endpoints. Plus, it seamlessly generates a polished resume based on your data and a provided LaTeX template, empowering you to showcase your qualifications with confidence.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
