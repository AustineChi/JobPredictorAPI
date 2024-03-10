package models

import (
    "math/big"
)

// Job represents a job listing in the job search app
type Job struct {
	JobID          int     `gorm:"column:job_id;primary_key" json:"jobId"`
	Title          string  `gorm:"column:title" json:"title"`
	Description    string  `gorm:"column:description" json:"description"`
	Company        string  `gorm:"column:company" json:"company"`
	Location       string  `gorm:"column:location" json:"location"`
	SkillsRequired string  `gorm:"column:skills_required" json:"skillsRequired"`
	Salary         big.Int `gorm:"column:salary" json:"salary"`
	EmploymentType string  `gorm:"column:employment_type" json:"employmentType"`
}

