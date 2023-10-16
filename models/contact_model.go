package models

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Contact struct {
	gorm.Model
	CustomerID uuid.UUID `json:"customer_id"`
	Phone      string    `json:"phone"`
	Email      string    `json:"email"`
}
