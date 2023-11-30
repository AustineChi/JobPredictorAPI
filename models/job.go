package models

import (
    "math/big"
)

// Job represents a job listing in the job search app
type Job struct {
    JobID          int     `json:"jobId" gorm:"column:job_id;primary_key"`
    Title          string  `json:"title"`
    Description    string  `json:"description"`
    Company        string  `json:"company"`
    Location       string  `json:"location"`
    SkillsRequired string  `json:"skillsRequired" gorm:"column:skills_required"`
    Salary         big.Int `json:"salary"`
    EmploymentType string  `json:"employmentType" gorm:"column:employment_type"`
}

