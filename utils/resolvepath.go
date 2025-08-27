package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/goferwplynie/cutie/logger"
)

func Resolvepath(path string) (string, error) {
	if strings.HasPrefix(path, "~") {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}

		return strings.Replace(path, "~", home, 1), nil
	}
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	path = fmt.Sprintf("%v%c%v", wd, filepath.Separator, path)
	logger.Cute(path)

	newPath, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}
	return newPath, nil
}
