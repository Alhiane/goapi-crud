package config

import (
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitDB initializes the database
func Connect() {

	errEnv := godotenv.Load("../../.env")
	if errEnv != nil {
		panic(errEnv)
	}

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
		panic(err)
	}

	DB = connectDB

	// hwo to stop terminal
	// ctrl + c dosn't work
	// ctrl + z dosn't work
	// ctrl + d dosn't work

}
