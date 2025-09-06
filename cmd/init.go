package cmd

import (
	"fmt"
	"os"
	"os/exec"
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
	noGit    bool
	branch   string
	commit   string
	remote   string
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

	//git setup flags

	initCmd.Flags().BoolVarP(&noGit, "nogit", "G", false, "dont use git in this project")
	initCmd.Flags().StringVar(&branch, "branch", "", "create main git branch")
	initCmd.Flags().StringVar(&commit, "commit", "", "stage and commit changes with some message")
	initCmd.Flags().StringVar(&remote, "remote", "", "add remote for git")
}

func startProject(cmd *cobra.Command, args []string) {
	var deadline time.Time
	var err error
	path, err := utils.Resolvepath(args[0])
	if err != nil {
		logger.Error(fmt.Sprintf("failed accessing path :c :%v", err))
	}
	name := args[1]

	if dl != "" {
		if deadline, err = time.Parse("2006-01-02", dl); err != nil {
			logger.Error(fmt.Sprintf("error while parsing deadline date ;c : %v", err))
			return
		}
	}
	reminderDuration := time.Duration(reminder) * 24 * time.Hour

	prj := project.New(deadline, name, filepath.Join(path, name), reminderDuration)

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
	projects, err := storage.GetProjects()
	if err != nil {
		logger.Error(err)
		return
	}
	projects = append(projects, *prj)
	if err = storage.SaveProjects(projects); err != nil {
		logger.Error(err)
		return
	}

	logger.Cute(fmt.Sprintf("your cute '%v' project has been created successfully and saved to storage. hope you don't abandon it cutie ;3", name))

	if err = storage.SyncReminders(true); err != nil {
		logger.Error(fmt.Sprintf("failed syncing reminders :c : %v", err))
	}

	if !noGit {
		if err := setupGit(prj.Path); err != nil {
			logger.Error(fmt.Sprintf("failed seting up git TwT : %v", err))
		}
	}
}

func runGit(path string, args ...string) error {
	cmd := exec.Command("git", args...)
	cmd.Dir = path
	return cmd.Run()
}

func setupGit(path string) error {
	if err := runGit(path, "init"); err != nil {
		return fmt.Errorf("can't setup git repository: %v", err)
	}

	if branch != "" {
		if err := runGit(path, "branch", "-M", branch); err != nil {
			return fmt.Errorf("can't setup main branch: %v", err)
		}

	}
	if commit != "" {
		if err := runGit(path, "add", "."); err != nil {
			return fmt.Errorf("can't stage changes: %v", err)
		}
		if err := runGit(path, "commit", "-m", commit); err != nil {
			return fmt.Errorf("can't commit changes: %v", err)
		}

	}
	if remote != "" {
		if err := runGit(path, "remote", "add", "origin", remote); err != nil {
			return fmt.Errorf("can't add remote: %v", err)
		}
	}

	return nil
}
