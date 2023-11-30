package services

import (
    "context"
    "database/sql"
    "errors"
    "JobPredictorAPI/models"
    "JobPredictorAPI/utils"
)

type SavedJobsService struct {
    db *sql.DB
}

// NewSavedJobsService creates a new instance of SavedJobsService
func NewSavedJobsService(db *sql.DB) *SavedJobsService {
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
    _, err = s.db.ExecContext(ctx, "INSERT INTO saved_jobs (user_id, job_id) VALUES (?, ?)", userID, jobID)
    return err
}

// GetSavedJobs retrieves all saved jobs for a user
func (s *SavedJobsService) GetSavedJobs(ctx context.Context, userID int) ([]models.Job, error) {
    // Logic to retrieve all saved jobs for the user
    rows, err := s.db.QueryContext(ctx, "SELECT job_id FROM saved_jobs WHERE user_id = ?", userID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var jobs []models.Job
    for rows.Next() {
        var jobID int
        if err := rows.Scan(&jobID); err != nil {
            return nil, err
        }

        job, err := utils.GetJobByID(s.db, jobID)
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
    _, err := s.db.ExecContext(ctx, "UPDATE saved_jobs SET job_id = ? WHERE user_id = ? AND job_id = ?", newJobID, userID, jobID)
    return err
}

// DeleteSavedJob removes a saved job from a user's list
func (s *SavedJobsService) DeleteSavedJob(ctx context.Context, userID int, jobID int) error {
    // Logic to delete a saved job
    _, err := s.db.ExecContext(ctx, "DELETE FROM saved_jobs WHERE user_id = ? AND job_id = ?", userID, jobID)
    return err
}

// isJobSaved checks if a job is already saved by a user
func (s *SavedJobsService) isJobSaved(ctx context.Context, userID int, jobID int) (bool, error) {
    var exists bool
    err := s.db.QueryRowContext(ctx, "SELECT EXISTS(SELECT 1 FROM saved_jobs WHERE user_id = ? AND job_id = ?)", userID, jobID).Scan(&exists)
    return exists, err
}

