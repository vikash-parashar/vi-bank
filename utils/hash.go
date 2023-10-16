package utils

import "golang.org/x/crypto/bcrypt"

// HashPassword takes a plain text password and returns its bcrypt hashed equivalent.
func HashPassword(password string) ([]byte, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return []byte{}, err
	}
	return hashedPassword, nil
}

// ComparePasswordHash takes a plain text password and a hashed password, and checks if they match.
func ComparePasswordHash(plainPassword string, hashedPassword []byte) error {
	err := bcrypt.CompareHashAndPassword(hashedPassword, []byte(plainPassword))
	return err
}
