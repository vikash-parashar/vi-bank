package models

import (
	"time"

	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
)

const (
	DocumentTypePanCard        = "Pan Card"
	DocumentTypeAadharCard     = "Aadhar Card"
	DocumentTypeDrivingLicence = "Driving Licence"
	DocumentTypeVoterCard      = "Voter Card"
	DocumentTypeRationCard     = "Ration Card"
)

type DocumentType string

type Document struct {
	gorm.Model
	DocumentName   string    `json:"document_name"`
	DocumentNumber string    `json:"document_number"`
	ExpireOn       time.Time `json:"expire_on"`
}

type KYCDetails struct {
	gorm.Model
	CustomerID  uuid.UUID  `json:"customer_id"`
	FirstName   string     `json:"first_name"`
	LastName    string     `json:"last_name"`
	DateOfBirth time.Time  `json:"date_of_birth"`
	Gender      string     `json:"gender"`
	Documents   []Document `json:"documents"`
}
