// config/user.go
package config

import (
	"vibank/models"

	"github.com/jinzhu/gorm"
)

func (db *Database) CreateUser(user *models.User) error {
	return db.DB.Create(user).Error
}

func (db *Database) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	err := db.DB.First(&user, id).Error
	return &user, err
}

func (db *Database) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := db.DB.Where("email = ?", email).First(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil // User not found
		}
		return nil, err // Other errors
	}
	return &user, nil
}

// AdminExists checks if an admin user already exists in the database.
func (db *Database) AdminExists() (bool, error) {
	var count int
	if err := db.DB.Model(&models.User{}).Where("user_type = ?", "admin").Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}
func (db *Database) GetAllUsers() ([]models.User, error) {
	var users []models.User
	if err := db.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (db *Database) UpdateUser(user *models.User) error {
	return db.DB.Save(user).Error
}

func (db *Database) DeleteUser(user *models.User) error {
	return db.DB.Delete(user).Error
}
