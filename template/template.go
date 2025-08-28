package template

import (
	"encoding/json"
	"fmt"
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

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return template, fmt.Errorf("template does not exist: %w", err)
	}

	file, err := os.Open(path)
	if err != nil {
		return template, fmt.Errorf("failed opening template: %w", err)
	}
	defer file.Close()

	if err = json.NewDecoder(file).Decode(&template); err != nil {
		return template, fmt.Errorf("failed reading template: %w", err)
	}

	return template, nil
}

func (t Template) Use(basePath string, projectName string) error {
	for _, v := range t.Files {
		idx := strings.LastIndex(v, string(filepath.Separator))
		if idx == -1 {
			os.Create(basePath + string(filepath.Separator) + v)
			continue
		}
		paths := v[:idx]
		file := v[idx+1:]
		if err := os.MkdirAll(basePath+string(filepath.Separator)+paths, 0755); err != nil {
			return err
		}
		if file != "" {
			os.Create(basePath + string(filepath.Separator) + v)
		}
	}

	for _, cmd := range t.Commands {
		cmd = strings.ReplaceAll(cmd, "$NAME", projectName)

		logger.Cute(fmt.Sprintf("Executing: %s", cmd))

		command := exec.Command("sh", "-c", cmd)
		command.Dir = basePath

		output, err := command.CombinedOutput()
		if err != nil {
			return fmt.Errorf("command failed: %s, error: %w, output: %s", cmd, err, string(output))
		}
	}

	return nil
}
