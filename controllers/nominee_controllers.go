package controllers

import (
	"database/sql"
	"time"
	"your-app/models"

	"github.com/google/uuid"
)

// CreateNominee creates a new nominee record.
func CreateNominee(db *sql.DB, firstName, lastName string, dateOfBirth time.Time, relation string) (uuid.UUID, error) {
	query := "INSERT INTO nominees (first_name, last_name, date_of_birth, relation, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id"
	var nomineeID uuid.UUID
	now := time.Now()
	err := db.QueryRow(query, firstName, lastName, dateOfBirth, relation, now, now).Scan(&nomineeID)
	if err != nil {
		return uuid.Nil, err
	}
	return nomineeID, nil
}

// GetNominee retrieves a nominee record by ID.
func GetNominee(db *sql.DB, nomineeID uuid.UUID) (*models.Nominee, error) {
	query := "SELECT id, first_name, last_name, date_of_birth, relation, created_at, updated_at FROM nominees WHERE id = $1"
	nominee := &models.Nominee{}
	err := db.QueryRow(query, nomineeID).Scan(&nominee.ID, &nominee.FirstName, &nominee.LastName, &nominee.DateOfBirth, &nominee.Relation, &nominee.CreatedAt, &nominee.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return nominee, nil
}

// UpdateNominee updates a nominee record.
func UpdateNominee(db *sql.DB, nomineeID uuid.UUID, firstName, lastName string, dateOfBirth time.Time, relation string) error {
	query := "UPDATE nominees SET first_name = $1, last_name = $2, date_of_birth = $3, relation = $4, updated_at = $5 WHERE id = $6"
	now := time.Now()
	_, err := db.Exec(query, firstName, lastName, dateOfBirth, relation, now, nomineeID)
	if err != nil {
		return err
	}
	return nil
}

// DeleteNominee deletes a nominee record by ID.
func DeleteNominee(db *sql.DB, nomineeID uuid.UUID) error {
	query := "DELETE FROM nominees WHERE id = $1"
	_, err := db.Exec(query, nomineeID)
	if err != nil {
		return err
	}
	return nil
}
