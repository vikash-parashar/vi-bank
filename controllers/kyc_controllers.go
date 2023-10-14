package controllers

import (
	"database/sql"
	"go-bank/models"
	"time"

	"github.com/google/uuid"
)

// CreateKYCDetails creates KYC details for a customer.
func CreateKYCDetails(db *sql.DB, customerID uuid.UUID, firstName, lastName string, dateOfBirth time.Time, gender, aadhaarNumber, panCardNumber string) (uuid.UUID, error) {
	query := "INSERT INTO kyc_details (customer_id, first_name, last_name, date_of_birth, gender, aadhaar_number, pan_card_number, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id"
	var kycID uuid.UUID
	now := time.Now()
	err := db.QueryRow(query, customerID, firstName, lastName, dateOfBirth, gender, aadhaarNumber, panCardNumber, now, now).Scan(&kycID)
	if err != nil {
		return uuid.Nil, err
	}
	return kycID, nil
}

// GetKYCDetails retrieves KYC details by ID.
func GetKYCDetails(db *sql.DB, kycID uuid.UUID) (*models.KYCDetails, error) {
	query := "SELECT id, customer_id, first_name, last_name, date_of_birth, gender, aadhaar_number, pan_card_number, created_at, updated_at FROM kyc_details WHERE id = $1"
	kycDetails := &models.KYCDetails{}
	err := db.QueryRow(query, kycID).Scan(&kycDetails.ID, &kycDetails.CustomerID, &kycDetails.FirstName, &kycDetails.LastName, &kycDetails.DateOfBirth, &kycDetails.Gender, &kycDetails.AadhaarNumber, &kycDetails.PANCardNumber, &kycDetails.CreatedAt, &kycDetails.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return kycDetails, nil
}

// UpdateKYCDetails updates KYC details for a customer.
func UpdateKYCDetails(db *sql.DB, kycID uuid.UUID, firstName, lastName string, dateOfBirth time.Time, gender, aadhaarNumber, panCardNumber string) error {
	query := "UPDATE kyc_details SET first_name = $1, last_name = $2, date_of_birth = $3, gender = $4, aadhaar_number = $5, pan_card_number = $6, updated_at = $7 WHERE id = $8"
	now := time.Now()
	_, err := db.Exec(query, firstName, lastName, dateOfBirth, gender, aadhaarNumber, panCardNumber, now, kycID)
	if err != nil {
		return err
	}
	return nil
}

// DeleteKYCDetails deletes KYC details by ID.
func DeleteKYCDetails(db *sql.DB, kycID uuid.UUID) error {
	query := "DELETE FROM kyc_details WHERE id = $1"
	_, err := db.Exec(query, kycID)
	if err != nil {
		return err
	}
	return nil
}
