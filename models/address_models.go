package models

import (
	"time"

	"github.com/google/uuid"
)

type Address struct {
	ID         uuid.UUID `json:"address_id"`
	CustomerID uuid.UUID `json:"customer_id"`
	Street     string    `json:"street"`
	City       string    `json:"city"`
	State      string    `json:"state"`
	PinCode    string    `json:"pin_code"`
	Country    string    `json:"country"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
