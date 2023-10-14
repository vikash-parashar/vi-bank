package models

import "time"

type Address struct {
	FirstName   string
	LastName    string
	DateOfBirth time.Time
	PanCard     string
	AadharCard  string
}
