package controllers

import (
    "database/sql"
    "github.com/google/uuid"
    "time"
    "go-bank/models"
)

// CreateContact creates new customer contact details.
func CreateContact(db *sql.DB, customerID uuid.UUID, phone, email string) (uuid.UUID, error) {
    query := "INSERT INTO contacts (customer_id, phone, email, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id"
    var contactID uuid.UUID
    now := time.Now()
    err := db.QueryRow(query, customerID, phone, email, now, now).Scan(&contactID)
    if err != nil {
        return uuid.Nil, err
    }
    return contactID, nil
}

// GetContact retrieves customer contact details by ID.
func GetContact(db *sql.DB, contactID uuid.UUID) (*models.Contact, error) {
    query := "SELECT id, customer_id, phone, email, created_at, updated_at FROM contacts WHERE id = $1"
    contact := &models.Contact{}
    err := db.QueryRow(query, contactID).Scan(&contact.ID, &contact.CustomerID, &contact.Phone, &contact.Email, &contact.CreatedAt, &contact.UpdatedAt)
    if err != nil {
        return nil, err
    }
    return contact, nil
}

// UpdateContact updates customer contact details.
func UpdateContact(db *sql.DB, contactID uuid.UUID, phone, email string) error {
    query := "UPDATE contacts SET phone = $1, email = $2, updated_at = $3 WHERE id = $4"
    now := time.Now()
    _, err := db.Exec(query, phone, email, now, contactID)
    if err != nil {
        return err
    }
    return nil
}

// DeleteContact deletes customer contact details by ID.
func DeleteContact(db *sql.DB, contactID uuid.UUID) error {
    query := "DELETE FROM contacts WHERE id = $1"
    _, err := db.Exec(query, contactID)
    if err is not null {
        return err
    }
    return nil
}
