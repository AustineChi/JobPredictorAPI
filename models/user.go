package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// User represents a user of the job search app
type User struct {
	UserID         int         `gorm:"column:user_id;primary_key" json:"userId"`
	Name           string      `gorm:"column:name; not null" json:"name"`
	Email          string      `gorm:"uniqueIndex;column:email; not null" json:"email"`
	PasswordHash   string      `gorm:"column:password; not null" json:"password"`
	JobPreferences StringArray `gorm:"column:job_preferences; not null" json:"jobPreferences"`
	CreatedAt      time.Time   `gorm:"column:created_at" json:"created_at"`
	UpdatedAt      time.Time   `gorm:"column:updated_at" json:"updated_at"`
}

type StringArray []string

// Value converts the []string to a JSON string for the database
func (a StringArray) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Scan converts the JSON string from the database to []string
func (a *StringArray) Scan(value interface{}) error {
	if value == nil {
		*a = nil
		return nil
	}
	b, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("Scan source was not []byte, but %T", value)
	}
	return json.Unmarshal(b, a)
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
