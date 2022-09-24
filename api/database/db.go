package database

import (
	"fmt"
	"my-joke-book/database/models"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("DB error(Init): ", err)
	}
	db.AutoMigrate(&models.User{})
	var users = []models.User{
		{Name: "jinzhu1", Email: "test@com", Password: "password"},
		{Name: "jinzhu2", Email: "test2@com", Password: "password"},
		{Name: "jinzhu3", Email: "test3@com", Password: "password"}}
	db.Create(&users)
	DB = db
}
