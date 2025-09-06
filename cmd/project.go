package cmd

import projectstorage "github.com/goferwplynie/cutie/projectStorage"

func SetFreeze(projectName string, freeze bool) error {
	storage := projectstorage.New("")

	if err := storage.Setup(); err != nil {
		return err
	}

	projects, err := storage.GetProjects()
	if err != nil {
		return err
	}

	for i, v := range projects {
		if v.Name == projectName {
			projects[i].Archived = freeze
			continue
		}
	}

	if err := storage.SaveProjects(projects); err != nil {
		return err
	}

	if err := storage.SyncReminders(true); err != nil {
		return err
	}

	return nil
}
