package models

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

const (
	RoleCustomer Role = "customer"
	RoleAdmin    Role = "admin"
)

type Role string

type User struct {
	gorm.Model
	CustomerID uuid.UUID `gorm:"type:uuid;default:null" json:"customer_id"`
	Email      string    `gorm:"type:varchar(255);unique;not null" json:"email"`
	Password   string    `gorm:"type:varchar" json:"password"`
	Role       Role      `gorm:"type:varchar(50)" json:"role"`
}
