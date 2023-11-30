package models

// JobRecommendation represents a job recommendation for a user
type JobRecommendation struct {
    RecommendationID int     `json:"recommendationId" gorm:"column:recommendation_id;primary_key"`
    UserID           int     `json:"userId" gorm:"column:user_id"`
    JobID            int     `json:"jobId" gorm:"column:job_id"`
    Score            float64 `json:"score"`
}
