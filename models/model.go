package models

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // driver
)

var db *sqlx.DB

// GetDBConnection GetDBConnection
func GetDBConnection() *sqlx.DB {

	if db != nil {
		return db
	}

	dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatalln(err)
	}

	connectionString := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		dbPort,
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"),
	)
	db, err = sqlx.Connect("postgres", connectionString)
	// remove hardcoding
	if err != nil {
		log.Fatalln(err)
	}
	return db
}
