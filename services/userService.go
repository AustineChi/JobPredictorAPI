package services

import (
	"JobPredictorAPI/models"
	"context"
	"errors"

	"gorm.io/gorm"
)

type UserService struct {
	Db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{Db: db}
}

// CreateUser creates a new user in the database
func (s *UserService) CreateUser(ctx context.Context, user *models.User) error {
	query := "INSERT INTO public.user (name, email, password, job_preferences) VALUES ($1, $2, $3, $4)"
	result := s.Db.Exec(query, user.Name, user.Email, user.PasswordHash, user.JobPreferences).WithContext(ctx)
	if result.Error != nil {
		return result.Error
	}
	// Check affected rows if needed
	if result.RowsAffected == 0 {
		return errors.New("no rows affected")
	}
	return nil
}

// GetUserByID retrieves a user by their ID
func (s *UserService) GetUserByID(ctx context.Context, userID int) (*models.User, error) {
	query := "SELECT user_id, name, email, job_preferences FROM public.user WHERE user_id = $1"
	row := s.Db.Exec(query, userID).WithContext(ctx)

	var user models.User
	result := row.Scan(&user).WithContext(ctx)
	if result.Error != nil {
		return nil, result.Error
	}
	// Check affected rows if needed
	if result.RowsAffected == 0 {
		return nil, errors.New("no rows affected")
	}
	return &user, nil
}

// UpdateUser updates a user's details in the database
func (s *UserService) UpdateUser(ctx context.Context, user *models.User) error {
	query := "UPDATE public.user SET name = $1, email = $2, job_preferences = $3 WHERE user_id = $4"
	result := s.Db.Exec(query, user.Name, user.Email, user.JobPreferences, user.UserID).WithContext(ctx)
	if result.Error != nil {
		return result.Error
	}
	// Check affected rows if needed
	if result.RowsAffected == 0 {
		return errors.New("no rows affected")
	}
	return nil
}

// DeleteUser removes a user from the database
func (s *UserService) DeleteUser(ctx context.Context, userID int) error {
	query := "DELETE FROM public.user WHERE user_id = $1"
	result := s.Db.Exec(query, userID).WithContext(ctx)
	if result.Error != nil {
		return result.Error
	}
	// Check affected rows if needed
	if result.RowsAffected == 0 {
		return errors.New("no rows affected")
	}
	return nil
}
