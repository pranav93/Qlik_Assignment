package models

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // driver
)

// GetDBConnection GetDBConnection
func GetDBConnection() *sqlx.DB {
	db, err := sqlx.Connect("postgres", "postgres://awesome_user:secret_password@localhost:5432/qlik_database?sslmode=disable")
	// remove hardcoding
	if err != nil {
		log.Fatalln(err)
	}
	return db
}
