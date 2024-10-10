// Path: project_root/config/config.go

package config

import (
	"os"

	"github.com/joho/godotenv"
)

var config map[string]string

func Load() error {
	var err error
	config, err = godotenv.Read()
	return err
}

func Get(key string) string {
	if value, exists := config[key]; exists {
		return value
	}
	return os.Getenv(key)
}
