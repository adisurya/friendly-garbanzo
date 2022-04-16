package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/garbanzo?parseTime=true&loc=Asia%2FJakarta")

	return db, err
}
