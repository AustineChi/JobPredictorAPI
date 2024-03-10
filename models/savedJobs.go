package models

import (
    "time"
)

// SavedJob represents a job saved by a user
type SavedJob struct {
    SavedJobID int       `gorm:"column:saved_job_id;primary_key" json:"savedJobId"`
    UserID     int       `gorm:"column:user_id" json:"userId"`
    JobID      int       `gorm:"column:job_id" json:"jobId"`
    DateSaved  time.Time `gorm:"column:date_saved" json:"dateSaved"`
}
