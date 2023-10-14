package models

import (
	"time"

	"github.com/google/uuid"
)

type UserType string

const (
	Cst UserType = "Customer"
	Adm UserType = "Admin"
)

type User struct {
	ID         uuid.UUID `json:"user_id"`
	UserType   UserType  `json:"user_type"`
	CustomerID uuid.UUID `json:"customer_id,omitempty"`
	Email      string    `json:"user_email"`
	Password   string    `json:"password"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
