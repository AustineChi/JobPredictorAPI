package models

import (
    "time"
)

// Interaction represents an interaction of a user with a job
type Interaction struct {
    InteractionID int       `gorm:"column:interaction_id;primary_key" json:"interactionId"`
    UserID        int       `gorm:"column:user_id" json:"userId"`
    JobID         int       `gorm:"column:job_id" json:"jobId"`
    Type          string    `gorm:"column:type" json:"type"`
    Timestamp     time.Time `gorm:"column:timestamp" json:"timestamp"`
}

