package models

import (
    "golang.org/x/crypto/bcrypt"
)


// User represents a user of the job search app
type User struct {
    UserID         int    `gorm:"column:user_id;primary_key" json:"userId"`
    Name           string `gorm:"column:name" json:"name"`
    Email          string `gorm:"column:email" json:"email"`
    PasswordHash   string `gorm:"column:password" json:"-"`
    JobPreferences string `gorm:"column:job_preferences" json:"jobPreferences"`
}

// SetPassword sets a hashed password for a user
func (u *User) SetPassword(password string) error {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    if err != nil {
        return err
    }
    u.PasswordHash = string(bytes)
    return nil
}

// CheckPassword compares a provided password with the user's hashed password
func (u *User) CheckPassword(password string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
    return err == nil
}
