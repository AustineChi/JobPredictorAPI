package services

import (
	"JobPredictorAPI/models"
	"context"
	"errors"

	"gorm.io/gorm"
)

type SavedJobsService struct {
	db *gorm.DB
	jb JobService
}

// NewSavedJobsService creates a new instance of SavedJobsService
func NewSavedJobsService(db *gorm.DB) *SavedJobsService {
	return &SavedJobsService{db: db}
}

// SaveJob saves a job to a user's saved jobs
func (s *SavedJobsService) SaveJob(ctx context.Context, userID int, jobID int) error {
	// Logic to save a job to the user's saved jobs
	// Check if the job already exists in saved jobs
	exists, err := s.isJobSaved(ctx, userID, jobID)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("job already saved")
	}

	// SQL query to insert the job into the saved jobs table
	result := s.db.Exec("INSERT INTO saved_jobs (user_id, job_id) VALUES (?, ?)", userID, jobID).WithContext(ctx)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// GetSavedJobs retrieves all saved jobs for a user
func (s *SavedJobsService) GetSavedJobs(ctx context.Context, userID int) ([]models.Job, error) {
	//create a slice to hold ID's
	var jobIDs []int

	result := s.db.Raw("SELECT job_id FROM saved_jobs WHERE user_id = ?", userID).WithContext(ctx).Scan(&jobIDs)
	if result.Error != nil {
		return nil, result.Error
	}
	//slice to hold the jobs
	var jobs []models.Job
	// Iterate through the retrieved job IDs and fetch the corresponding jobs
	for _, jobID := range jobIDs {
		job, err := s.jb.GetJobByID(ctx, jobID)
		if err != nil {
			return nil, err
		}
		jobs = append(jobs, *job)
	}
	return jobs, nil
}

// UpdateSavedJob updates a saved job entry
func (s *SavedJobsService) UpdateSavedJob(ctx context.Context, userID int, jobID int, newJobID int) error {
	// Logic to update a saved job entry
	result := s.db.Exec("UPDATE saved_jobs SET job_id = ? WHERE user_id = ? AND job_id = ?", newJobID, userID, jobID).WithContext(ctx)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("no rows affetced by query")
	}
	return nil
}

// DeleteSavedJob removes a saved job from a user's list
func (s *SavedJobsService) DeleteSavedJob(ctx context.Context, userID int, jobID int) error {
	// Logic to delete a saved job
	result := s.db.Raw("DELETE FROM saved_jobs WHERE user_id = ? AND job_id = ?", userID, jobID).WithContext(ctx)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("no rows affetced by query")
	}
	return nil
}

// isJobSaved checks if a job is already saved by a user
func (s *SavedJobsService) isJobSaved(ctx context.Context, userID int, jobID int) (bool, error) {
	var exists bool
	result := s.db.Exec("SELECT EXISTS(SELECT 1 FROM saved_jobs WHERE user_id = ? AND job_id = ?)", userID, jobID).WithContext(ctx).Scan(&exists)
	if result.Error != nil {
		return exists, result.Error
	}
	if result.RowsAffected == 0 {
		return exists, errors.New("no rows affetced by query")
	}
	return exists, nil
}
