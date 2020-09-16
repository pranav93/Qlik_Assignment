package models

import "log"

// Message Message
type Message struct {
	ID      int64  `json:"id"`
	Message string `json:"title"`
}

// CreateMessage CreateMessage
func CreateMessage(message string) (int64, error) {
	log.Println("before GetDBConnection")
	db := GetDBConnection()
	log.Println("after GetDBConnection")
	tx := db.MustBegin()
	q := `
	INSERT INTO messages(message)
	VALUES ($1) 
	RETURNING id`

	var ID int64
	err := tx.Get(&ID, q, message)
	if err != nil {
		log.Println("Error is", err)
		tx.Rollback()
		return 0, err
	}
	tx.Commit()
	return ID, nil
}
