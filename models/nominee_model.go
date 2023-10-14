package models

import (
	"time"

	"github.com/google/uuid"
)

const (
	Father   = "Father"
	Mother   = "Mother"
	Spouse   = "Husband/Wife"
	Daughter = "Daughter"
	Son      = "Son"
	Sister   = "Sister"
	Brother  = "Brother"
)

type Nominee struct {
	ID          uuid.UUID `json:"nominee_id"`
	FirstName   string    `json:"nominee_first_name"`
	LastName    string    `json:"nominee_last_name"`
	DateOfBirth time.Time `json:"nominee_date_of_birth"`
	Relation    string    `json:"relation_with_nominee"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
