package models

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

const (
	AddressTypeCurrent   = "Current"
	AddressTypePermanent = "Permanent"
)

type Address struct {
	gorm.Model
	CustomerID uuid.UUID `json:"customer_id"`
	Type       string    `json:"type"`
	Street     string    `json:"street"`
	City       string    `json:"city"`
	State      string    `json:"state"`
	PinCode    string    `json:"pin_code"`
	Country    string    `json:"country"`
}
