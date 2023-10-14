package models

import "time"

const (
	INR = "INR"
	USD = "USD"
	EUR = "EUR"
)

// Define your data models here
type BankAccount struct {
	ID        int       `json:"id"`
	Holder    string    `json:"holder"`
	Balance   float64   `json:"balance"`
	Currency  string    `json:"currency"`
	CreatedAt time.Time `json:"created_at"`
}
