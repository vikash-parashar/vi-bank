package models

import (
	"time"

	"github.com/google/uuid"
)

type Contact struct {
	ID         uuid.UUID `json:"contact_id"`
	CustomerID uuid.UUID `json:"customer_id"`
	Phone      string    `json:"phone"`
	Email      string    `json:"email"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	DeleteAt   time.Time `json:"deleted_at"`
}
