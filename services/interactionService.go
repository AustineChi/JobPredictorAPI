package services

import (
	"JobPredictorAPI/models"
	"context"

	"gorm.io/gorm"
)

type InteractionService struct {
	Db *gorm.DB
}

func NewInteractionService(db *gorm.DB) *InteractionService {
	return &InteractionService{Db: db}
}

func (s *InteractionService) CreateInteraction(ctx context.Context, interaction *models.Interaction) error {
	if err := s.Db.Create(interaction).WithContext(ctx).Error; err != nil {
		return err
	}
	return nil
}

func (s *InteractionService) UpdateInteraction(ctx context.Context, interaction *models.Interaction) error {
	query := "UPDATE public.interactions SET user_id = $1, job_id = $2, type= $3, timestamp=$4 WHERE interaction_id = $5"

	result := s.Db.Exec(query, interaction.UserID, interaction.JobID, interaction.Type, interaction.Timestamp, interaction.InteractionID).WithContext(ctx)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// LogInteraction creates a new interaction record in the database
func (s *InteractionService) LogInteraction(ctx context.Context, interaction *models.Interaction) error {
	query := "INSERT INTO public.interactions (user_id, job_id, type, timestamp) VALUES ($1, $2, $3, $4)"
	if err := s.Db.Exec(query, interaction.UserID, interaction.JobID, interaction.Type, interaction.Timestamp).WithContext(ctx).Error; err != nil {
		return err
	}
	return nil
}

// GetInteractionsByUserID retrieves all interactions for a specific user
func (s *InteractionService) GetInteractionsByUserID(ctx context.Context, userID int) ([]models.Interaction, error) {
	var interactions []models.Interaction
	if err := s.Db.Raw("SELECT interaction_id, user_id, job_id, type, timestamp FROM public.interactions WHERE user_id = ?", userID).WithContext(ctx).Scan(&interactions).Error; err != nil {
		return nil, err
	}
	return interactions, nil
}

// DeleteInteraction removes an interaction record from the database
func (s *InteractionService) DeleteInteraction(ctx context.Context, interactionID int) error {
	if err := s.Db.Exec("DELETE FROM public.interactions WHERE interaction_id = ?", interactionID).WithContext(ctx).Error; err != nil {
		return err
	}
	return nil
}
