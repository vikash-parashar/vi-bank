// config/db.go
package config

import (
	"fmt"
	"os"
	"vibank/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Database struct {
	DB *gorm.DB
}

func NewDatabase() (*Database, error) {
	// Read individual database-related environment variables
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// Create the connection string using individual variables
	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbHost, dbPort, dbUser, dbName, dbPassword)

	// Initialize and configure your database connection
	db, err := gorm.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&models.User{}, &models.Address{}, &models.Account{}, &models.Document{}, &models.KYCDetails{}, &models.Transaction{}, &models.Nominee{}, &models.Contact{}, &models.Customer{})
	// Set database connection options if needed
	db.LogMode(true) // Enable SQL query logging, for example

	return &Database{DB: db}, nil
}

func (db *Database) Close() {
	if db.DB != nil {
		defer db.DB.Close()
	}
}
