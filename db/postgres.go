// Path: project_root/db/postgres.go

package db

import (
  "Go_lang_Microservice/config"
	"Go_lang_Microservice/api/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitPostgres() error {
	var err error
	dsn := config.Get("DATABASE_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	// Auto migrate the Task model
	return DB.AutoMigrate(&models.Task{})
}
