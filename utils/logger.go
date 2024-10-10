// Path: project_root/utils/logger.go

package utils

import (
	"log"
	"os"
)

var Logger *log.Logger

func init() {
	Logger = log.New(os.Stdout, "API: ", log.LstdFlags|log.Lshortfile)
}

func LogInfo(message string) {
	Logger.Printf("INFO: %s", message)
}

func LogError(message string, err error) {
	Logger.Printf("ERROR: %s: %v", message, err)
}
