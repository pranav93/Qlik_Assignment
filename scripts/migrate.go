package scripts

import (
	"log"

	"github.com/pranav93/Qlik_Assignment/models"
)

// Migrate Migrate
func Migrate() {
	schema := `
	CREATE TABLE IF NOT EXISTS messages (
		id serial PRIMARY KEY,
		message VARCHAR(255)
	);
	`
	db := models.GetDBConnection()
	_, err := db.Exec(schema)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Ran migration")

	// tx, err := db.Begin()
	// if err != nil {
	// 	log.Fatalln(err)
	// 	tx.Rollback()
	// }
	// tx.Exec(schema)
	// tx.Commit()
}
