package projectstorage

import (
	"os"

	"github.com/goferwplynie/cutie/logger"
	"github.com/goferwplynie/cutie/project"
)

type FileStorage struct {
}

func New(storageType string) ProjectStorage {
	return newFs()
}

func newFs() *FileStorage {
	return &FileStorage{}
}

func (f *FileStorage) Setup() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	_, err = os.Stat(home + "/.cutie")
	if err != nil {
		notExists := os.IsNotExist(err)
		if notExists {
			err := os.Mkdir(home+"/.cutie", 0755)
			if err != nil {
				return err
			}
			logger.Cute("created app directory :3")
		} else {
			return err
		}
	}
	logger.Cute("app directory already exists =^w^=")
	return nil

}

func (f *FileStorage) SaveProject(project project.Project) {
	panic("not implemented") // TODO: Implement
}
func (f *FileStorage) GetProject(name string) (_ project.Project) {
	panic("not implemented") // TODO: Implement
}
func (f *FileStorage) GetProjects() (_ []project.Project) {
	panic("not implemented") // TODO: Implement
}
