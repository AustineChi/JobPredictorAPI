package models

import (
	"errors"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// User represents a user of the job search app
type User struct {
	UserID         int       `gorm:"column:user_id;primary_key" json:"userId"`
	Name           string    `gorm:"column:name; not null" json:"name"`
	Email          string    `gorm:"uniqueIndex;column:email; not null" json:"email"`
	PasswordHash   string    `gorm:"column:password; not null" json:"password"`
	JobPreferences string    `gorm:"column:job_preferences; not null" json:"jobPreferences"`
	CreatedAt      time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at" json:"updated_at"`
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
