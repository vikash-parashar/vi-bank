package models

import (
	"time"

	"github.com/google/uuid"
)

const (
	INR            = "INR"
	USD            = "USD"
	EUR            = "EUR"
	SavingAccount  = "Saving Account"
	CurrentAccount = "Current Account"
	SalaryAccount  = "Salary Account"
)

type AccountType string

type Account struct {
	ID           uuid.UUID     `json:"account_id"`
	CustomerID   uuid.UUID     `json:"customer_id"`
	Balance      float64       `json:"balance"`
	Currency     string        `json:"currency"`
	AccountType  AccountType   `json:"account_type"`
	Nominee      Nominee       `json:"nominee"`
	CreatedAt    time.Time     `json:"created_at"`
	UpdatedAt    time.Time     `json:"updated_at"`
	Transactions []Transaction `json:"transactions"`
}
