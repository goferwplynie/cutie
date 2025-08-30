package projectstorage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/goferwplynie/cutie/logger"
	"github.com/goferwplynie/cutie/project"
)

type RemindersCache struct {
	LastUpdated time.Time  `json:"last_updated"`
	Reminders   []string   `json:"reminders"`
	Deadlines   []Deadline `json:"deadlines"`
}

type Deadline struct {
	Name     string `json:"name"`
	DaysLeft int    `json:"days_left"`
}

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
	_, err := os.Stat(f.appDir)
	if err != nil {
		notExists := os.IsNotExist(err)
		if notExists {
			if err := os.Mkdir(f.appDir, 0755); err != nil {
				return err
			}
			if _, err := os.Create(f.appDir + "/reminders.json"); err != nil {
				return err
			}
			logger.Cute("created app directory :3")
		} else {
			return err
		}
	}
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

func (f *FileStorage) SyncReminders(forced bool) error {
	now := time.Now()
	file, err := os.OpenFile(filepath.Join(f.appDir, "reminders.json"), os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("cant open reminders file :c : %v", err)
	}
	defer file.Close()
	var reminders RemindersCache

	info, err := file.Stat()
	if err != nil {
		return err
	}
	if info.Size() > 0 {
		err = json.NewDecoder(file).Decode(&reminders)
		if err != nil {
			return err
		}
	} else {
		reminders = RemindersCache{
			LastUpdated: now,
		}
		forced = true
	}

	if !forced && sameDay(reminders.LastUpdated, now) {
		return nil
	}

	projects, err := f.GetProjects()
	if err != nil {
		return err
	}

	reminders.Reminders = []string{}
	reminders.Deadlines = []Deadline{}

	for _, p := range projects {
		if p.Reminder != 0 {
			info, err := os.Stat(p.Path)
			if err != nil {
				return err
			}
			lastMod := info.ModTime()
			if time.Since(lastMod) > p.Reminder {
				reminders.Reminders = append(reminders.Reminders, p.Name)
			}

		}
		if !p.Deadline.IsZero() {
			if time.Until(p.Deadline) < time.Hour*24*7 {
				reminders.Deadlines = append(reminders.Deadlines, Deadline{
					Name:     p.Name,
					DaysLeft: int(p.Deadline.Sub(now).Hours() / 24),
				})
			}

		}
	}
	file.Truncate(0)
	file.Seek(0, 0)
	if err := json.NewEncoder(file).Encode(&reminders); err != nil {
		return err
	}

	return nil
}

func (f *FileStorage) GetProject(name string) (_ project.Project) {
	panic("not implemented") // TODO: Implement
}
func (f *FileStorage) GetProjects() ([]project.Project, error) {
	var projects []project.Project
	file, err := os.OpenFile(f.appDir+"/projects.json", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return projects, err
	}
	defer file.Close()

	var projectsDB projectDB
	info, err := file.Stat()
	if err != nil {
		return projects, err
	}
	if info.Size() > 0 {
		err = json.NewDecoder(file).Decode(&projectsDB)
		if err != nil {
			return projects, err
		}
	} else {
		projectsDB = projectDB{
			Projects: make([]project.Project, 0),
		}
	}

	return projectsDB.Projects, nil
}

func (f *FileStorage) GetReminders() (RemindersCache, error) {
	var reminders RemindersCache

	file, err := os.OpenFile(filepath.Join(f.appDir, "reminders.json"), os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return reminders, fmt.Errorf("cant open reminders file :c : %v", err)
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&reminders)
	if err != nil {
		return reminders, err
	}

	return reminders, nil

}

func sameDay(a, b time.Time) bool {
	ay, am, ad := a.Date()
	by, bm, bd := b.Date()
	return ay == by && am == bm && ad == bd
}
