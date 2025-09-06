package cmd

import (
	"fmt"
	"slices"

	"github.com/goferwplynie/cutie/logger"
	projectstorage "github.com/goferwplynie/cutie/projectStorage"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove [project name]",
	Short: "remove project from list",
	Args:  cobra.ExactArgs(1),
	Long:  `removes project from lis tand stops tracking reminders and deadlines`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("remove called")

		storage := projectstorage.New("")

		if err := storage.Setup(); err != nil {
			logger.Error(err)
			return
		}

		projects, err := storage.GetProjects()
		if err != nil {
			logger.Error(err)
			return
		}
		for i, v := range projects {
			logger.Cute(v)
			if v.Name == args[0] {
				projects = slices.Delete(projects, i, i+1)
				logger.Cute(projects)
			}
		}

		if err := storage.SaveProjects(projects); err != nil {
			logger.Error(err)
			return
		}

		if err := storage.SyncReminders(true); err != nil {
			logger.Error(err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
