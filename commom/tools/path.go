package tools

import (
	"os"
	"path/filepath"
)

func GetCurrentPath() (currentPath string) {
	var err error
	if currentPath, err = filepath.Abs(filepath.Dir(os.Args[0])); err != nil {
		panic(err)
	}
	return
}

func GetLogsPath() (path string) {
	path = filepath.Join(GetCurrentPath(), "logs")
	return
}
