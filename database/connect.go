package database

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("database/connect.go:Connect(): Error loading .env file")
	}

	url := os.Getenv("CONNECTION_URL")
	db, err := sql.Open("mysql", url)

	return db, err
}
