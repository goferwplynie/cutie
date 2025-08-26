package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/goferwplynie/cutie/logger"
	"github.com/goferwplynie/cutie/project"
	projectstorage "github.com/goferwplynie/cutie/projectStorage"
)

func main() {
	setVerbose()
	logger.Cute("verbose mode on cutie >///<")
	args := os.Args
	fmt.Println(args[1])
	fmt.Println(logger.Verbose)

	if len(args) > 1 {
		switch args[1] {
		case "init":
			createProj(args[2:])
		case "projects":
			getProjects(args[2:])
		}
	}
}

func setVerbose() {
	verbose := flag.Bool("v", false, "verbose mode")
	flag.Parse()
	logger.Verbose = *verbose
}

func createProj(args []string) {
	path := args[0]
	name := args[1]

	flagset := flag.NewFlagSet("init", flag.ExitOnError)
	dl := flagset.String("dl", "", "deadline")
	//template := flagset.String("template", "", "template")
	days := flagset.Int("reminder", 0, "reminder in days")
	flagset.Parse(args)

	var deadline time.Time
	var err error
	if *dl != "" {
		deadline, err = time.Parse("2006-01-02", *dl)
	}
	if err != nil {
		logger.Error(fmt.Sprintf("error while parsing deadline date ;c : %v", err))
		return
	}
	reminder := time.Duration(*days) * 24 * time.Hour

	prj := project.New(deadline, name, path, reminder)

	storage := projectstorage.New("")
	err = storage.Setup()
	if err != nil {
		logger.Error(fmt.Sprintf("error at setup :c : %v", err))
	}

	storage.SaveProject(prj)

	logger.Cute(fmt.Sprintf("your cute '%v' project has been created succesfully and saved to storage. hope you don't abandon it cutie ;3", name))
}
func getProjects(args []string) {}
