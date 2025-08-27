package cmd

import (
	"fmt"

	"github.com/goferwplynie/cutie/logger"
	projectstorage "github.com/goferwplynie/cutie/projectStorage"
	"github.com/spf13/cobra"
)

var projectsCmd = &cobra.Command{
	Use:   "projects",
	Short: "List or manage projects",
	Run: func(cmd *cobra.Command, args []string) {
		storage := projectstorage.New("")
		if err := storage.Setup(); err != nil {
			logger.Error(fmt.Sprintf("error setting up storage: %v", err))
			return
		}
		// TODO: Implement project listing logic
	},
}
