package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/goferwplynie/cutie/logger"
	"github.com/goferwplynie/cutie/project"
	projectstorage "github.com/goferwplynie/cutie/projectStorage"
	"github.com/goferwplynie/cutie/utils"
)

var (
	dl       string
	reminder int
)

var initCmd = &cobra.Command{
	Use:   "init [path] [name]",
	Short: "create a new project",
	Args:  cobra.ExactArgs(2),
	Run:   startProject,
}

func init() {
	rootCmd.AddCommand(initCmd)

	initCmd.Flags().StringVar(&dl, "dl", "", "project deadline (YYYY-MM-DD)")
	initCmd.Flags().IntVar(&reminder, "reminder", 0, "reminder in days")
}

func startProject(cmd *cobra.Command, args []string) {
	var deadline time.Time
	var err error
	path, err := utils.Resolvepath(args[0])
	if err != nil {
		logger.Error(fmt.Sprintf("failed accessing path :c :%v", err))
	}
	name := args[1]

	if deadline, err = time.Parse("2006-01-02", dl); err != nil {
		logger.Error(fmt.Sprintf("error while parsing deadline date ;c : %v", err))
		return
	}
	reminderDuration := time.Duration(reminder) * 24 * time.Hour

	prj := project.New(deadline, name, path, reminderDuration)

	storage := projectstorage.New("")
	if err = storage.Setup(); err != nil {
		logger.Error(fmt.Sprintf("error at setup :c : %v", err))
		return
	}

	storage.SaveProject(prj)

	logger.Cute(fmt.Sprintf("your cute '%v' project has been created successfully and saved to storage. hope you don't abandon it cutie ;3", name))
}
