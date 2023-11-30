package models

import (
    "golang.org/x/crypto/bcrypt"
)


// User represents a user of the job search app
type User struct {
    UserID         int    `json:"userId" gorm:"column:user_id;primary_key"`
    Name           string `json:"name"`
    Email          string `json:"email"`
    PasswordHash   string `json:"-" gorm:"column:password"`
    JobPreferences string `json:"jobPreferences" gorm:"column:job_preferences"`
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
