package models

import (
	"time"

	"github.com/google/uuid"
)

type KYCDetails struct {
	ID            uuid.UUID `json:"kyc_id"`
	CustomerID    uuid.UUID `json:"customer_id"`
	FirstName     string    `json:"first_name"`
	LastName      string    `json:"last_name"`
	DateOfBirth   time.Time `json:"date_of_birth"`
	Gender        string    `json:"gender"`
	AadhaarNumber string    `json:"aadhaar_number"`
	PANCardNumber string    `json:"pan_card_number"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
