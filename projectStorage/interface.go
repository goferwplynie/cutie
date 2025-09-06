package projectstorage

import "github.com/goferwplynie/cutie/project"

type ProjectStorage interface {
	Setup() error
	GetProjects() ([]project.Project, error)
	GetTemplateFolder() string
	SyncReminders(forced bool) error
	GetReminders() (RemindersCache, error)
	SaveProjects([]project.Project) error
}
