package models

import (
	"time"

	"github.com/gofrs/uuid"
)

const (
	Admin    = "Admin"
	Customer = "Customer"
)

type User struct {
	UserID      uuid.UUID `json:"user_id"`
	UserName    string    `json:"username"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	DateOfBirth time.Time `json:"date_of_birth"`
	Type        string    `json:"user_type"`
	Password    string    `json:"password"`
}
