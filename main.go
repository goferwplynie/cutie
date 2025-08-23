package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/goferwplynie/cutie/logger"
	projectstorage "github.com/goferwplynie/cutie/projectStorage"
)

func main() {
	setVerbose()
	args := os.Args

	if len(args) > 1 {
		switch args[1] {
		case "init":
			createProj(args[1:])
		case "projects":
			getProjects(args[1:])
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
	storage := projectstorage.New("")
	err := storage.Setup()
	if err != nil {
		logger.Error(fmt.Sprintf("error at setup :c : %v", err))
	}
}
func getProjects(args []string) {}
