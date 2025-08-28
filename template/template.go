package template

import (
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/goferwplynie/cutie/logger"
)

type Template struct {
	Files    []string `json:"files"`
	Commands []string `json:"commands"`
}

func LoadTemplate(path string) (Template, error) {
	var template Template
	_, err := os.Stat(path)
	if err != nil {
		notExists := os.IsNotExist(err)
		if notExists {
			logger.Error("template does not exist :c")
			return template, err
		}
	}
	file, err := os.Open(path)
	if err != nil {
		logger.Error("failed opening template :c")
		return template, err
	}
	defer file.Close()

	if err = json.NewDecoder(file).Decode(&template); err != nil {
		logger.Error("failed reading template :c")
		return template, err
	}
	return template, nil
}

func (t Template) Use(path string, project_name string) error {
	for _, v := range t.Files {
		logger.Cute(v)
		idx := strings.LastIndex(v, string(filepath.Separator))
		if idx == -1 {
			os.Create(path + string(filepath.Separator) + v)
			continue
		}
		paths := v[:idx]
		file := v[idx+1:]
		if err := os.MkdirAll(path+string(filepath.Separator)+paths, 0755); err != nil {
			return err
		}
		if file != "" {
			os.Create(path + string(filepath.Separator) + v)
		}
	}
	for _, v := range t.Commands {
		v = strings.ReplaceAll(v, "$NAME", project_name)
		command := exec.Command("sh", "-c", v)
		command.Dir = path
		if err := command.Run(); err != nil {
			return err
		}
	}
	return nil
}
