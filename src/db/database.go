package db

import (
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
	"coremind/models"
)

var DB *gorm.DB

func InitDB() {
    var err error
    DB, err = gorm.Open(sqlite.Open("users.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    DB.AutoMigrate(&models.User{})
    DB.AutoMigrate(&models.TrainingResult{})
}