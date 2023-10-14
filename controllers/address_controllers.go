package controllers

import (
	"database/sql"
	"go-bank/models"
	"time"

	"github.com/google/uuid"
)

// CreateAddress creates a new customer address.
func CreateAddress(db *sql.DB, customerID uuid.UUID, street, city, state, pinCode, country string) (uuid.UUID, error) {
	query := "INSERT INTO addresses (customer_id, street, city, state, pin_code, country, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id"
	var addressID uuid.UUID
	now := time.Now()
	err := db.QueryRow(query, customerID, street, city, state, pinCode, country, now, now).Scan(&addressID)
	if err != nil {
		return uuid.Nil, err
	}
	return addressID, nil
}

// GetAddress retrieves a customer address by ID.
func GetAddress(db *sql.DB, addressID uuid.UUID) (*models.Address, error) {
	query := "SELECT id, customer_id, street, city, state, pin_code, country, created_at, updated_at FROM addresses WHERE id = $1"
	address := &models.Address{}
	err := db.QueryRow(query, addressID).Scan(&address.ID, &address.CustomerID, &address.Street, &address.City, &address.State, &address.PinCode, &address.Country, &address.CreatedAt, &address.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return address, nil
}

// UpdateAddress updates a customer address.
func UpdateAddress(db *sql.DB, addressID uuid.UUID, street, city, state, pinCode, country string) error {
	query := "UPDATE addresses SET street = $1, city = $2, state = $3, pin_code = $4, country = $5, updated_at = $6 WHERE id = $7"
	now := time.Now()
	_, err := db.Exec(query, street, city, state, pinCode, country, now, addressID)
	if err != nil {
		return err
	}
	return nil
}

// DeleteAddress deletes a customer address by ID.
func DeleteAddress(db *sql.DB, addressID uuid.UUID) error {
	query := "DELETE FROM addresses WHERE id = $1"
	_, err := db.Exec(query, addressID)
	if err != nil {
		return err
	}
	return nil
}
