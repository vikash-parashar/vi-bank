package controllers

import (
	"database/sql"
	"time"
	"your-app/models"

	"github.com/google/uuid"
)

// CreateAccount creates a new bank account.
func CreateAccount(db *sql.DB, customerID uuid.UUID, initialBalance float64, currency string, accountType models.AccountType) (uuid.UUID, error) {
	query := "INSERT INTO accounts (customer_id, balance, currency, account_type, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id"
	var accountID uuid.UUID
	now := time.Now()
	err := db.QueryRow(query, customerID, initialBalance, currency, accountType, now, now).Scan(&accountID)
	if err != nil {
		return uuid.Nil, err
	}
	return accountID, nil
}

// GetAccount retrieves a bank account by ID.
func GetAccount(db *sql.DB, accountID uuid.UUID) (*models.Account, error) {
	query := "SELECT id, customer_id, balance, currency, account_type, created_at, updated_at FROM accounts WHERE id = $1"
	account := &models.Account{}
	err := db.QueryRow(query, accountID).Scan(&account.ID, &account.CustomerID, &account.Balance, &account.Currency, &account.Type, &account.CreatedAt, &account.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return account, nil
}

// UpdateAccount updates the balance of a bank account.
func UpdateAccount(db *sql.DB, accountID uuid.UUID, newBalance float64) error {
	query := "UPDATE accounts SET balance = $1, updated_at = $2 WHERE id = $3"
	now := time.Now()
	_, err := db.Exec(query, newBalance, now, accountID)
	if err != nil {
		return err
	}
	return nil
}

// DeleteAccount deletes a bank account by ID.
func DeleteAccount(db *sql.DB, accountID uuid.UUID) error {
	query := "DELETE FROM accounts WHERE id = $1"
	_, err := db.Exec(query, accountID)
	if err != nil {
		return err
	}
	return nil
}

// GetAllAccounts retrieves all bank accounts.
func GetAllAccounts(db *sql.DB) ([]*models.Account, error) {
	query := "SELECT id, customer_id, balance, currency, account_type, created_at, updated_at FROM accounts"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var accounts []*models.Account
	for rows.Next() {
		account := &models.Account{}
		err := rows.Scan(&account.ID, &account.CustomerID, &account.Balance, &account.Currency, &account.Type, &account.CreatedAt, &account.UpdatedAt)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}
	return accounts, nil
}
