package controllers

import (
	"database/sql"
	"time"
	"your-app/models"

	"github.com/google/uuid"
)

// CreateCustomer creates a new customer with associated KYC details, address, and contact.
func CreateCustomer(db *sql.DB, kycDetails models.KYCDetails, address models.Address, contact models.Contact) (uuid.UUID, error) {
	tx, err := db.Begin()
	if err != nil {
		return uuid.Nil, err
	}
	defer tx.Rollback()

	// Create KYC details and get the KYC ID
	kycID, err := CreateKYCDetails(tx, kycDetails)
	if err != nil {
		return uuid.Nil, err
	}

	// Create an address and get the address ID
	addressID, err := CreateAddress(tx, address)
	if err != nil {
		return uuid.Nil, err
	}

	// Create contact details and get the contact ID
	contactID, err := CreateContact(tx, contact)
	if err != nil {
		return uuid.Nil, err
	}

	// Create the customer with associated KYC, address, and contact
	query := "INSERT INTO customers (kyc_id, address_id, contact_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	var customerID uuid.UUID
	now := time.Now()
	err = tx.QueryRow(query, kycID, addressID, contactID, now, now).Scan(&customerID)
	if err != nil {
		return uuid.Nil, err
	}

	if err = tx.Commit(); err != nil {
		return uuid.Nil, err
	}

	return customerID, nil
}

// GetCustomer retrieves a customer by ID.
func GetCustomer(db *sql.DB, customerID uuid.UUID) (*models.Customer, error) {
	query := "SELECT id, kyc_id, address_id, contact_id, created_at, updated_at FROM customers WHERE id = $1"
	customer := &models.Customer{}
	err := db.QueryRow(query, customerID).Scan(&customer.ID, &customer.KYCDetails.ID, &customer.Address.ID, &customer.Contact.ID, &customer.CreatedAt, &customer.UpdatedAt)
	if err != nil {
		return nil, err
	}

	// Retrieve associated KYC details, address, and contact
	customer.KYCDetails, err = GetKYCDetails(db, customer.KYCDetails.ID)
	if err != nil {
		return nil, err
	}

	customer.Address, err = GetAddress(db, customer.Address.ID)
	if err != nil {
		return nil, err
	}

	customer.Contact, err = GetContact(db, customer.Contact.ID)
	if err != nil {
		return nil, err
	}

	return customer, nil
}

// UpdateCustomer updates customer information, KYC details, address, and contact.
func UpdateCustomer(db *sql.DB, customerID uuid.UUID, kycDetails models.KYCDetails, address models.Address, contact models.Contact) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Update KYC details
	err = UpdateKYCDetails(tx, kycDetails.ID, kycDetails.FirstName, kycDetails.LastName, kycDetails.DateOfBirth, kycDetails.Gender, kycDetails.AadhaarNumber, kycDetails.PANCardNumber)
	if err != nil {
		return err
	}

	// Update address
	err = UpdateAddress(tx, address.ID, address.Street, address.City, address.State, address.PinCode, address.Country)
	if err != nil {
		return err
	}

	// Update contact details
	err = UpdateContact(tx, contact.ID, contact.Phone, contact.Email)
	if err != nil {
		return err
	}

	// Update customer record
	query := "UPDATE customers SET kyc_id = $1, address_id = $2, contact_id = $3, updated_at = $4 WHERE id = $5"
	now := time.Now()
	_, err = tx.Exec(query, kycDetails.ID, address.ID, contact.ID, now, customerID)
	if err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

// DeleteCustomer deletes a customer by ID along with associated KYC details, address, and contact.
func DeleteCustomer(db *sql.DB, customerID uuid.UUID) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Get customer information
	customer, err := GetCustomer(tx, customerID)
	if err != nil {
		return err
	}

	// Delete associated KYC details, address, and contact
	err = DeleteKYCDetails(tx, customer.KYCDetails.ID)
	if err != nil {
		return err
	}

	err = DeleteAddress(tx, customer.Address.ID)
	if err != nil {
		return err
	}

	err = DeleteContact(tx, customer.Contact.ID)
	if err != nil {
		return err
	}

	// Delete customer record
	query := "DELETE FROM customers WHERE id = $1"
	_, err = tx.Exec(query, customerID)
	if err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}
