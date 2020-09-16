package models

import (
	"log"
)

// Message Message
type Message struct {
	ID      int    `json:"id"`
	Message string `json:"title"`
}

// CreateMessage CreateMessage
func CreateMessage(message string) (int64, error) {
	db := GetDBConnection()
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

// GetMessage GetMessage
func GetMessage(ID int) (*Message, error) {
	db := GetDBConnection()
	tx := db.MustBegin()
	q := `
	SELECT * FROM messages 
	WHERE id=$1 LIMIT 1`

	var m Message
	err := tx.Get(&m, q, ID)
	if err != nil {
		log.Println("Error is", err)
		tx.Rollback()
		return nil, err
	}
	log.Println("Message is", m)
	return &m, nil
}

// DeleteMessage DeleteMessage
func DeleteMessage(ID int) error {
	db := GetDBConnection()
	tx := db.MustBegin()
	q := `
	DELETE FROM messages 
	WHERE id=$1`

	_, err := tx.Exec(q, ID)
	if err != nil {
		log.Println("Error is", err)
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

// GetMessageList GetMessageList
func GetMessageList() ([]Message, error) {
	db := GetDBConnection()
	q := `
	SELECT * FROM messages`

	rows, err := db.Query(q)
	if err != nil {
		log.Println("Error is", err)
		return nil, err
	}
	defer rows.Close()

	var m []Message
	for rows.Next() {
		var ID int
		var message string
		err = rows.Scan(&ID, &message)
		if err != nil {
			// handle this error
			panic(err)
		}
		m = append(m, Message{
			ID:      ID,
			Message: message,
		})
	}
	return m, nil
}
