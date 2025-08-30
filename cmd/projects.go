package cmd

import (
	"fmt"

	"github.com/goferwplynie/cutie/logger"
	projectstorage "github.com/goferwplynie/cutie/projectStorage"
	"github.com/jedib0t/go-pretty/v6/table"
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
		projects, err := storage.GetProjects()
		if err != nil {
			logger.Error(fmt.Sprintf("error loading projects TwT: %v", err))
		}
		tw := table.NewWriter()
		tw.AppendHeader(table.Row{"Name", "Start", "Deadline", "Reminder", "Path", "Frozen"})

		for _, prj := range projects {
			tw.AppendRow(table.Row{prj.Name, prj.Start.Format("2006-01-02"), prj.Deadline.Format("2006-01-02"), prj.Reminder.Hours() / 24, prj.Path, prj.Archived})
		}

		tw.SetAutoIndex(true)
		tw.SetStyle(table.StyleRounded)
		fmt.Print(tw.Render())

	},
}
