package tools

import (
	"os"
	"path/filepath"
	"time"
)

func todayFilename() string {
	today := time.Now().Format("2006-01-02")
	return today + ".txt"
}

func NewLogFile() *os.File {
	filename := todayFilename()
	filePath := filepath.Join(GetLogsPath(), filename)
	f, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	return f
}
