package controllers

import "database/sql"

// CreateAccount creates a new bank account.
func CreateAccount(db *sql.DB, holder string, initialBalance float64) (int, error) {
	query := "INSERT INTO accounts (holder, balance) VALUES ($1, $2) RETURNING id"
	var accountID int
	err := db.QueryRow(query, holder, initialBalance).Scan(&accountID)
	if err != nil {
		return 0, err
	}
	return accountID, nil
}

// GetAccount retrieves a bank account by ID.
func GetAccount(db *sql.DB, accountID int) (*models.Account, error) {
	query := "SELECT id, holder, balance FROM accounts WHERE id = $1"
	account := &Account{}
	err := db.QueryRow(query, accountID).Scan(&account.ID, &account.Holder, &account.Balance)
	if err != nil {
		return nil, err
	}
	return account, nil
}

// UpdateAccount updates the balance of a bank account.
func UpdateAccount(db *sql.DB, accountID int, newBalance float64) error {
	query := "UPDATE accounts SET balance = $1 WHERE id = $2"
	_, err := db.Exec(query, newBalance, accountID)
	if err != nil {
		return err
	}
	return nil
}

// DeleteAccount deletes a bank account by ID.
func DeleteAccount(db *sql.DB, accountID int) error {
	query := "DELETE FROM accounts WHERE id = $1"
	_, err := db.Exec(query, accountID)
	if err != nil {
		return err
	}
	return nil
}

// GetAllAccounts retrieves all bank accounts.
func GetAllAccounts(db *sql.DB) ([]*models.Account, error) {
	query := "SELECT id, holder, balance FROM accounts"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var accounts []*models.Account
	for rows.Next() {
		account := &Account{}
		err := rows.Scan(&account.ID, &account.Holder, &account.Balance)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}
	return accounts, nil
}

// RecordTransaction records a transaction in the database.
func RecordTransaction(db *sql.DB, sourceID, destinationID int, amount float64) (int, error) {
	tx, err := db.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	// Deduct the amount from the source account
	sourceQuery := "UPDATE accounts SET balance = balance - $1 WHERE id = $2 RETURNING id"
	var sourceAccountID int
	err = tx.QueryRow(sourceQuery, amount, sourceID).Scan(&sourceAccountID)
	if err != nil {
		return 0, err
	}

	// Add the amount to the destination account
	destinationQuery := "UPDATE accounts SET balance = balance + $1 WHERE id = $2 RETURNING id"
	var destinationAccountID int
	err = tx.QueryRow(destinationQuery, amount, destinationID).Scan(&destinationAccountID)
	if err != nil {
		return 0, err
	}

	// Insert the transaction record
	transactionQuery := "INSERT INTO transactions (source, destination, amount) VALUES ($1, $2, $3) RETURNING id"
	var transactionID int
	err = tx.QueryRow(transactionQuery, sourceAccountID, destinationAccountID, amount).Scan(&transactionID)
	if err != nil {
		return 0, err
	}

	if err = tx.Commit(); err != nil {
		return 0, err
	}
	return transactionID, nil
}
