package models

import (
    "time"
)

// Interaction represents an interaction of a user with a job
type Interaction struct {
    InteractionID int       `json:"interactionId" gorm:"column:interaction_id;primary_key"`
    UserID        int       `json:"userId" gorm:"column:user_id"`
    JobID         int       `json:"jobId" gorm:"column:job_id"`
    Type          string    `json:"type"`
    Timestamp     time.Time `json:"timestamp"`
}
