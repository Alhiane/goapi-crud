package config

import (
	"os"

	"github.com/Alhiane/goapi-crud/pkg/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitDB initializes the database
func Connect() {
	// getting the cerdentials from the .env file
	dbhost := os.Getenv("DB_HOST")
	dbport := os.Getenv("DB_PORT")
	dbuser := os.Getenv("DB_USER")
	dbpassword := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	dsn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"

	// open a connection to the database
	connectDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// migrate the schema to the database (book table)
	err = connectDB.AutoMigrate(&models.Book{})
	if err != nil {
		panic("failed to migrate database")
	}

	DB = connectDB

}
