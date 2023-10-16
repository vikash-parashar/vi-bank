package models

import (
	"time"

	"github.com/google/uuid"
)

type Customer struct {
	ID           uuid.UUID     `json:"customer_id"`
	Account      Account       `json:"account"`
	Addresses    []Address     `json:"addresses"`
	KYCDetails   KYCDetails    `json:"kyc_details"`
	Transactions []Transaction `json:"transactions"`
	Contact      Contact       `json:"contact"`
	CreatedAt    time.Time     `json:"created_at"`
	UpdatedAt    time.Time     `json:"updated_at"`
}
