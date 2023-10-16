package models

import (
	"time"

	"github.com/google/uuid"
)

const (
	Success = "Success"
	Fail    = "Failed"
	Cancel  = "Canceled"
)

type TransactionStatus string

type Transaction struct {
	ID            uuid.UUID         `json:"transaction_id"`
	CustomerID    uuid.UUID         `json:"customer_id"`
	SourceID      uuid.UUID         `json:"source_account_id"`
	DestinationID uuid.UUID         `json:"destination_account_id"`
	TxnAmount     float64           `json:"transaction_amount"`
	Timestamp     time.Time         `json:"timestamp"`
	Status        TransactionStatus `json:"transaction_status"`
}
