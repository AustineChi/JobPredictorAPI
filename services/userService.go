package services

import (
	"JobPredictorAPI/models"
	"context"
	"errors"
	"log"
	"time"

	"gorm.io/gorm"
)

type UserService struct {
	Db *gorm.DB
	M  models.User
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{Db: db}
}

// CreateUser creates a new user in the database
func (s *UserService) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	hashpassword, err := s.M.SetPassword(user.PasswordHash)
	if err != nil {
		return nil, err
	}
	newUser := &models.User{
		Name:           user.Name,
		Email:          user.Email,
		PasswordHash:   string(hashpassword),
		JobPreferences: user.JobPreferences,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
	err = s.Db.Create(newUser).WithContext(ctx).Error
	if err != nil {
		log.Println("Unable to create user", err)
		return nil, err
	}
	return newUser, nil
}

func (s *UserService) GetEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	err := s.Db.Where("email=?", email).WithContext(ctx).First(&user).Error
	if err != nil {
		log.Println("unable to fetch email:", err)
		return nil, err
	}
	return &user, nil
}

// GetUserByID retrieves a user by their ID
func (s *UserService) GetUserByID(ctx context.Context, userID int) (*models.User, error) {

	var user models.User
	result := s.Db.Model(&models.User{}).Select(" user_id, name, email, job_preferences, created_at, updated_at").
		Where("user_id = $1", userID).WithContext(ctx).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, result.Error
		}
		return nil, result.Error
	}
	return &user, nil
}

// UpdateUser updates a user's details in the database
func (s *UserService) UpdateUser(ctx context.Context, user *models.User) error {
	query := "UPDATE public.users SET name = $1, email = $2, job_preferences = $3 WHERE user_id = $4"
	result := s.Db.Exec(query, user.Name, user.Email, user.JobPreferences, user.UserID).WithContext(ctx)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// DeleteUser removes a user from the database
func (s *UserService) DeleteUser(ctx context.Context, userID int) error {
	query := "DELETE FROM public.users WHERE user_id = $1"
	result := s.Db.Exec(query, userID).WithContext(ctx)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
