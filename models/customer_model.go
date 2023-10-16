package models

import (
	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
)

type Customer struct {
	gorm.Model
	AccountID    uuid.UUID     `gorm:"unique_index"`
	Account      Account       `json:"account" gorm:"foreignkey:AccountID"`
	Addresses    []Address     `json:"addresses" gorm:"foreignkey:CustomerID"`
	KYCDetails   KYCDetails    `json:"kyc_details" gorm:"foreignkey:CustomerID"`
	Transactions []Transaction `json:"transactions" gorm:"foreignkey:CustomerID"`
	Contact      Contact       `json:"contact" gorm:"foreignkey:CustomerID"`
}
