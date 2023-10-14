package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type Customer struct {
	ID         uuid.UUID  `json:"customer_id"`
	KYCDetails KYCDetails `json:"kyc_details"`
	Address    Address    `json:"address"`
	Contact    Contact    `json:"contact"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
}
