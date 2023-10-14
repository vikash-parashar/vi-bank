package controllers

import (
	"database/sql"
	"go-bank/models"
	"log"
	"time"

	"github.com/google/uuid"
)

// CreateUser creates a new user.
func CreateUser(db *sql.DB, u *models.User) (uuid.UUID, error) {
	query := "INSERT INTO users (user_id,user_type, customer_id, email, password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id"
	if u.UserType == "Customer" {
		u.CustomerID = u.ID
	}
	u.ID = uuid.New()
	now := time.Now()
	res, err := db.Exec(query, u.ID, u.UserType, u.CustomerID, u.Email, u.Password, now, now)

	if err != nil {
		return uuid.Nil, err
	}
	log.Printf("Rows Affected : %d .\n", res.RowsAffected)
	return u.ID, nil
}

// GetUser retrieves a user by ID.
func GetUser(db *sql.DB, userID uuid.UUID) (*models.User, error) {
	query := "SELECT id, user_type, customer_id, email, password, created_at, updated_at FROM users WHERE id = $1"
	user := &models.User{}
	err := db.QueryRow(query, userID).Scan(&user.ID, &user.UserType, &user.CustomerID, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// UpdateUser updates user information.
func UpdateUserPassword(db *sql.DB, userID uuid.UUID, email, password string) error {
	query := "UPDATE users SET email = $1, password = $2, updated_at = $3 WHERE id = $4"
	now := time.Now()
	_, err := db.Exec(query, email, password, now, userID)
	if err != nil {
		return err
	}
	return nil
}

// DeleteUser deletes a user by ID.
func DeleteUser(db *sql.DB, userID uuid.UUID) error {
	query := "DELETE FROM users WHERE id = $1"
	_, err := db.Exec(query, userID)
	if err != nil {
		return err
	}
	return nil
}
