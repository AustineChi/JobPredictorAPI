package models

// JobRecommendation represents a job recommendation for a user

type JobRecommendation struct {
    RecommendationID int     `gorm:"column:recommendation_id;primary_key" json:"recommendationId"`
    UserID           int     `gorm:"column:user_id" json:"userId"`
    JobID            int     `gorm:"column:job_id" json:"jobId"`
    Score            float64 `gorm:"column:score" json:"score"`
}
