package models

import (
	"database/sql/driver"
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
	Salary         BigInt `gorm:"column:salary" json:"salary"`
	EmploymentType string  `gorm:"column:employment_type" json:"employmentType"`
}

//Implement a Valuer and Scanner interface for SQL to handle BigInt 
type BigInt struct {
	big.Int
}

// Value converts BigInt to a database value
func (bi BigInt)Value()(driver.Value, error){
	return bi.String(), nil
}

// Scan converts a database value to BigInt
func (bi *BigInt) Scan(value interface{}) error {
    var s string
    if value != nil {
        s = string(value.([]byte))
    }
    bi.SetString(s, 10)
    return nil
}