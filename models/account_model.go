package models

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

const (
	CurrencyINR    = "INR"
	CurrencyUSD    = "USD"
	CurrencyEUR    = "EUR"
	AccountSaving  = "Saving Account"
	AccountCurrent = "Current Account"
	AccountSalary  = "Salary Account"
)

type AccountType string

type Account struct {
	gorm.Model
	CustomerID   uuid.UUID     `json:"customer_id"`
	Balance      float64       `json:"balance"`
	Currency     string        `json:"currency"`
	Type         AccountType   `json:"account_type"`
	Nominee      Nominee       `json:"nominee"`
	NomineeID    uuid.UUID     `json:"nominee_id"`
	Transactions []Transaction `json:"transactions"`
}
