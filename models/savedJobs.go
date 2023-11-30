package models

import (
    "time"
)

// SavedJob represents a job saved by a user
type SavedJob struct {
    SavedJobID int       `json:"savedJobId" gorm:"column:saved_job_id;primary_key"`
    UserID     int       `json:"userId" gorm:"column:user_id"`
    JobID      int       `json:"jobId" gorm:"column:job_id"`
    DateSaved  time.Time `json:"dateSaved" gorm:"column:date_saved"`
}
