package models

import (
	"time"

	"github.com/google/uuid"
)

const (
	PanCard        = "Pan Card"
	AadharCard     = "Aadhar Card"
	DrivingLicence = "Driving Licence"
	VoterCard      = "Voter Card"
	RationCard     = "Ration Card"
)

type Document struct {
	DocumentName   string    `json:"document_name"`
	DocumentNumber string    `json:"document_number"`
	ExpireOn       time.Time `json:"expire_on"`
}

type KYCDetails struct {
	ID          uuid.UUID  `json:"kyc_id"`
	CustomerID  uuid.UUID  `json:"customer_id"`
	FirstName   string     `json:"first_name"`
	LastName    string     `json:"last_name"`
	DateOfBirth time.Time  `json:"date_of_birth"`
	Gender      string     `json:"gender"`
	Documents   []Document `json:"documents"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}
