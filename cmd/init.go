package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"time"

	"github.com/spf13/cobra"

	"github.com/goferwplynie/cutie/logger"
	"github.com/goferwplynie/cutie/project"
	projectstorage "github.com/goferwplynie/cutie/projectStorage"
	"github.com/goferwplynie/cutie/template"
	"github.com/goferwplynie/cutie/utils"
)

var (
	dl       string
	reminder int
	tmpl     string
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
	initCmd.Flags().StringVar(&tmpl, "template", "", "template to use while creating project")
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

	if err = os.MkdirAll(path+string(filepath.Separator)+name, 0755); err != nil {
		logger.Error("cant create project directory :c")
		return
	}

	if tmpl != "" {
		templ, err := template.LoadTemplate(storage.GetTemplateFolder() + string(filepath.Separator) + tmpl)
		if err != nil {
			logger.Error(fmt.Sprintf("error loading template :c : %v", err))
			return
		}
		if err = templ.Use(path+string(filepath.Separator)+name, name); err != nil {
			logger.Error(fmt.Sprintf("error executing template commands :c : %v", err))
			return
		}

	}

	storage.SaveProject(prj)

	logger.Cute(fmt.Sprintf("your cute '%v' project has been created successfully and saved to storage. hope you don't abandon it cutie ;3", name))
}
