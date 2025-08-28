package projectstorage

import "github.com/goferwplynie/cutie/project"

type ProjectStorage interface {
	Setup() error
	SaveProject(project *project.Project) error
	GetProject(name string) project.Project
	GetProjects() []project.Project
	GetTemplateFolder() string
}
