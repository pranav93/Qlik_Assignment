package models

import (
	"log"

	"github.com/google/uuid"
)

// LogLatency LogLatency
func LogLatency(latency int64, path string) (uuid.UUID, error) {
	db := GetDBConnection()
	tx := db.MustBegin()
	q := `
	INSERT INTO latency_information(latency_ns, path)
	VALUES ($1, $2)
	RETURNING id`

	var ID uuid.UUID
	err := tx.Get(&ID, q, latency, path)
	if err != nil {
		log.Println("Error is", err)
		tx.Rollback()
		return [16]byte{}, err
	}
	tx.Commit()
	return ID, nil
}
