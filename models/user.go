package models

import (
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
)

// User represents a user of the job search app
type User struct {
	UserID         int    `gorm:"column:user_id;primary_key" json:"userId"`
	Name           string `gorm:"column:name" json:"name"`
	Email          string `gorm:"column:email" json:"email"`
	PasswordHash   string `gorm:"column:password" json:"password"`
	JobPreferences string `gorm:"column:job_preferences" json:"jobPreferences"`
}

// SetPassword sets a hashed password for a user
func (u *User) SetPassword(password string) ([]byte, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return nil, err
	}
	//u.PasswordHash = string(bytes)
	return bytes, nil
}

// CheckPassword compares a provided password with the user's hashed password
func (u *User) CheckPassword(password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			log.Println("Invalid Password")
			return false, nil
		default:
			return false, err
		}
	}
	log.Println("password validated")
	return true, nil
}
