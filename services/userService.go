package services

import (
    "context"
    "database/sql"
    "JobPredictorAPI/models"
)

type UserService struct {
    Db *sql.DB
}

func NewUserService(db *sql.DB) *UserService {
    return &UserService{Db: db}
}

// CreateUser creates a new user in the database
func (s *UserService) CreateUser(ctx context.Context, user *models.User) error {
    query := "INSERT INTO public.user (name, email, password, job_preferences) VALUES ($1, $2, $3, $4)"
    _, err := s.Db.ExecContext(ctx, query, user.Name, user.Email, user.Password, user.JobPreferences)
    return err
}

// GetUserByID retrieves a user by their ID
func (s *UserService) GetUserByID(ctx context.Context, userID int) (*models.User, error) {
    query := "SELECT user_id, name, email, job_preferences FROM public.user WHERE user_id = $1"
    row := s.Db.QueryRowContext(ctx, query, userID)

    var user models.User
    err := row.Scan(&user.UserID, &user.Name, &user.Email, &user.JobPreferences)
    if err != nil {
        return nil, err
    }

    return &user, nil
}

// UpdateUser updates a user's details in the database
func (s *UserService) UpdateUser(ctx context.Context, user *models.User) error {
    query := "UPDATE public.user SET name = $1, email = $2, job_preferences = $3 WHERE user_id = $4"
    _, err := s.Db.ExecContext(ctx, query, user.Name, user.Email, user.JobPreferences, user.UserID)
    return err
}

// DeleteUser removes a user from the database
func (s *UserService) DeleteUser(ctx context.Context, userID int) error {
    query := "DELETE FROM public.user WHERE user_id = $1"
    _, err := s.Db.ExecContext(ctx, query, userID)
    return err
}
