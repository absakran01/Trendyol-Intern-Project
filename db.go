package main

import (

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/joho/godotenv"
	"os"
)

func ConnectToDatabase() (*gorm.DB, error) {
	// This function would contain logic to connect to a database
	// For example, using a database driver to establish a connection
	// and handle any errors that may occur during the connection process.

	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		return nil, err
	}
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	sslmode := os.Getenv("DB_SSLMODE")
	timezone := os.Getenv("DB_TIMEZONE")

	// Create a new database connection using GORM
	dsn := "host=" + host + 
	" user=" + user + 
	" password=" + password + 
	" dbname=" + dbname + 
	" port=" + port + 
	" sslmode=" + sslmode + 
	" TimeZone=" + timezone


	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{}) 
	if err != nil {
		return nil, err
	}
	return db, nil
}