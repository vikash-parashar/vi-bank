package models

import (
	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
)

const (
	TransactionStatusSuccess  = "Success"
	TransactionStatusFailed   = "Failed"
	TransactionStatusCanceled = "Canceled"
)

type TransactionStatus string

type Transaction struct {
	gorm.Model
	CustomerID           uuid.UUID         `json:"customer_id"`
	SourceAccountID      uuid.UUID         `json:"source_account_id"`
	DestinationAccountID uuid.UUID         `json:"destination_account_id"`
	Amount               float64           `json:"amount"`
	Status               TransactionStatus `json:"status"`
}
