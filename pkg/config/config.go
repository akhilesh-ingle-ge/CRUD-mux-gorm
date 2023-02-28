package config

import (
	"crud/pkg/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var dsn = "postgres://postgres:12345@localhost:5432/gorm?sslmode=disable"

func SetUpDB() {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Could not connect to the database")
	}
	db.AutoMigrate(&models.Order{}, &models.Item{})
	DB = db
	fmt.Println("Connected to the Database")
}
