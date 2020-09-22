package models

import (
	"log"

	"github.com/google/uuid"
)

// LogLatency LogLatency
func LogLatency(latency int64) (uuid.UUID, error) {
	db := GetDBConnection()
	tx := db.MustBegin()
	q := `
	INSERT INTO latency_information(latency)
	VALUES ($1) 
	RETURNING id`

	var ID uuid.UUID
	err := tx.Get(&ID, q, latency)
	if err != nil {
		log.Println("Error is", err)
		tx.Rollback()
		return [16]byte{}, err
	}
	tx.Commit()
	return ID, nil
}
