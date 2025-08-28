package projectstorage

import (
	"encoding/json"
	"os"

	"github.com/goferwplynie/cutie/logger"
	"github.com/goferwplynie/cutie/project"
)

type FileStorage struct {
	appDir    string
	configDir string
}

type projectDB struct {
	Projects []project.Project `json:"projects"`
}

func New(storageType string) ProjectStorage {
	return newFs()
}

func newFs() *FileStorage {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	conf, err := os.UserConfigDir()
	if err != nil {
		panic(err)
	}
	return &FileStorage{
		appDir:    home + "/.cutie",
		configDir: conf + "/cutie",
	}
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

func (f *FileStorage) SaveProject(prj *project.Project) error {
	file, err := os.OpenFile(f.appDir+"/projects.json", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	var projects projectDB
	info, err := file.Stat()
	if err != nil {
		return err
	}
	if info.Size() > 0 {
		err = json.NewDecoder(file).Decode(&projects)
		if err != nil {
			return err
		}
	} else {
		projects = projectDB{
			Projects: make([]project.Project, 0),
		}
	}
	projects.Projects = append(projects.Projects, *prj)
	file.Truncate(0)
	file.Seek(0, 0)
	err = json.NewEncoder(file).Encode(projects)
	if err != nil {
		return err
	}
	return nil
}

func (f *FileStorage) GetTemplateFolder() string {
	return f.appDir
}

func (f *FileStorage) GetProject(name string) (_ project.Project) {
	panic("not implemented") // TODO: Implement
}
func (f *FileStorage) GetProjects() (_ []project.Project) {
	panic("not implemented") // TODO: Implement
}
