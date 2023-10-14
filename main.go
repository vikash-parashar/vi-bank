package main

import (
	"fmt"
	"go-bank/routes"
	"go-bank/storage"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	dsn string
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	// Create the PostgreSQL DSN
	dsn = fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)
}

func main() {
	db, err := storage.ConnectStorage(dsn)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()
	router := routes.SetupRoutes(db)

	fmt.Println("Server is running on :8080")
	http.ListenAndServe(":8080", router)

}
