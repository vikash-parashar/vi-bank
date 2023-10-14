package controllers

import (
	"database/sql"
	"time"
	"your-app/models"

	"github.com/google/uuid"
)

// CreateTransaction creates a new transaction record.
func CreateTransaction(db *sql.DB, sourceID, destinationID uuid.UUID, amount float64, status models.TransactionStatus) (uuid.UUID, error) {
	query := "INSERT INTO transactions (source_id, destination_id, amount, timestamp, status) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	var transactionID uuid.UUID
	now := time.Now()
	err := db.QueryRow(query, sourceID, destinationID, amount, now, status).Scan(&transactionID)
	if err != nil {
		return uuid.Nil, err
	}
	return transactionID, nil
}

// GetTransaction retrieves a transaction record by ID.
func GetTransaction(db *sql.DB, transactionID uuid.UUID) (*models.Transaction, error) {
	query := "SELECT id, source_id, destination_id, amount, timestamp, status FROM transactions WHERE id = $1"
	transaction := &models.Transaction{}
	err := db.QueryRow(query, transactionID).Scan(&transaction.ID, &transaction.SourceID, &transaction.DestinationID, &transaction.Amount, &transaction.Timestamp, &transaction.Status)
	if err != nil {
		return nil, err
	}
	return transaction, nil
}

// UpdateTransactionStatus updates the status of a transaction.
func UpdateTransactionStatus(db *sql.DB, transactionID uuid.UUID, status models.TransactionStatus) error {
	query := "UPDATE transactions SET status = $1 WHERE id = $2"
	_, err := db.Exec(query, status, transactionID)
	if err != nil {
		return err
	}
	return nil
}

// DeleteTransaction deletes a transaction record by ID.
func DeleteTransaction(db *sql.DB, transactionID uuid.UUID) error {
	query := "DELETE FROM transactions WHERE id = $1"
	_, err := db.Exec(query, transactionID)
	if err != nil {
		return err
	}
	return nil
}
