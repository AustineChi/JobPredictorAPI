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
	query := "INSERT INTO public.interaction(user_id, job_id, type, timestamp) VALUES($1, $2, $3, $4)"
	_, err := s.Db.ExecContext(ctx, query, interaction.UserID, interaction.JobID, interaction.Type, interaction.Timestamp)
	if err != nil {
		return err
	}
	return nil
}

func (s *InteractionService) UpdateInteraction(ctx context.Context, interaction *models.Interaction) error {
	query := "UPDATE public.interaction SET user_id = $1, job_id = $2, type= $3, timestamp=$4 WHERE interaction_id = $5"
	_, err := s.Db.QueryContext(ctx, query, interaction.UserID, interaction.JobID, interaction.Type, interaction.Timestamp, interaction.InteractionID)
	if err != nil {
		return err
	}
	return nil
}

// LogInteraction creates a new interaction record in the database
func (s *InteractionService) LogInteraction(ctx context.Context, interaction *models.Interaction) error {
	query := "INSERT INTO public.interaction (user_id, job_id, type, timestamp) VALUES ($1, $2, $3, $4)"
	_, err := s.Db.ExecContext(ctx, query, interaction.UserID, interaction.JobID, interaction.Type, interaction.Timestamp)
	return err
}

// GetInteractionsByUserID retrieves all interactions for a specific user
func (s *InteractionService) GetInteractionsByUserID(ctx context.Context, userID int) ([]models.Interaction, error) {
	query := "SELECT interaction_id, user_id, job_id, type, timestamp FROM public.interaction WHERE user_id = $1"
	rows, err := s.Db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var interactions []models.Interaction
	for rows.Next() {
		var interaction models.Interaction
		if err := rows.Scan(&interaction.InteractionID, &interaction.UserID, &interaction.JobID, &interaction.Type, &interaction.Timestamp); err != nil {
			return nil, err
		}
		interactions = append(interactions, interaction)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return interactions, nil
}

// DeleteInteraction removes an interaction record from the database
func (s *InteractionService) DeleteInteraction(ctx context.Context, interactionID int) error {
	query := "DELETE FROM public.interaction WHERE interaction_id = $1"
	_, err := s.Db.ExecContext(ctx, query, interactionID)
	return err
}
