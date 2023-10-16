package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

const (
	NomineeRelationFather   = "Father"
	NomineeRelationMother   = "Mother"
	NomineeRelationSpouse   = "Husband/Wife"
	NomineeRelationDaughter = "Daughter"
	NomineeRelationSon      = "Son"
	NomineeRelationSister   = "Sister"
	NomineeRelationBrother  = "Brother"
)

type NomineeRelation string

type Nominee struct {
	gorm.Model
	FirstName   string          `json:"first_name"`
	LastName    string          `json:"last_name"`
	DateOfBirth time.Time       `json:"date_of_birth"`
	Relation    NomineeRelation `json:"relation"`
}
