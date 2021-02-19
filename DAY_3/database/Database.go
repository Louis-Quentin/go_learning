package database

import (
	"SoftwareGoDay_3/models"
	"gorm.io/gorm"
)

type Database struct {
	dsn := getenv(DB_URL)
	gorm.DB
	models.Developer
}