package services

import (
	"JobPredictorAPI/models"
	"context"
	"errors"
	"log"
	"time"

	"gorm.io/gorm"
)

type SavedJobsService struct {
	db *gorm.DB
	jb *JobService //initialize jobservice connetion pool for swapping contents to the savedJob service
}

// NewSavedJobsService creates a new instance of SavedJobsService
func NewSavedJobsService(db *gorm.DB, jb* JobService) *SavedJobsService {
	return &SavedJobsService{db: db, jb: jb}
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
	savedJobs := &models.SavedJob{
		UserID:    userID,
		JobID:     jobID,
		DateSaved: time.Now(),
	}
	// SQL query to insert the job into the saved jobs table
	result := s.db.Create(savedJobs).WithContext(ctx)
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
        log.Println("unable to select job_id :", result.Error)
		return nil, result.Error
	}
	//slice to hold the jobs
	var jobs []models.Job
	// Iterate through the retrieved job IDs and fetch the corresponding jobs
	for _, jobID := range jobIDs {
		job, err := s.jb.GetJobByID(ctx, jobID)
		if err != nil {
            log.Println("unable to get jobID from job's DB:", err)
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
	return nil
}

// DeleteSavedJob removes a saved job from a user's list
func (s *SavedJobsService) DeleteSavedJob(ctx context.Context, userID int) error {
	// Logic to delete all saved jobs for a user
	result := s.db.Where("user_id = ?", userID).Delete(&models.SavedJob{}).WithContext(ctx)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
// isJobSaved checks if a job is already saved by a user
func (s *SavedJobsService) isJobSaved(ctx context.Context, userID int, jobID int) (bool, error) {
	var exists bool

	// SQL query to check if the job is already saved
	query := "SELECT EXISTS(SELECT 1 FROM saved_jobs WHERE user_id = ? AND job_id = ?)"
	//log.Println("Executing SQL Query:", query, "with userID:", userID, "and jobID:", jobID)

	result := s.db.Raw(query, userID, jobID).WithContext(ctx).Row()
	if err := result.Scan(&exists); err != nil {
		log.Println("Error scanning result:", err)
		return exists, err
	}
	return exists, nil
}
