package models

import (
	"time"
)

// Job represents a job posting.
type Job struct {
	BaseModel
	Title          string    `gorm:"not null"`
	Description    string    `gorm:"not null"`
	Location       string    `gorm:"not null"`
	EmploymentType string    `gorm:"not null"`
	Deadline       time.Time `gorm:"not null"`

	RecruiterID  uint
	Recruiter    User `gorm:"foreignKey:RecruiterID;references:id"`
	Applications []Application
}

// CreateJobRequest represents the request body for creating a new job.
type CreateJobRequest struct {
	Title          string `json:"title" binding:"required,min=3"`        // Minimum 3 characters for title
	Description    string `json:"description" binding:"required,min=10"` // Minimum 10 characters for description
	Location       string `json:"location" binding:"required,min=3"`     // Minimum 3 characters for location
	EmploymentType string `json:"employment_type" binding:"required,min=5"`
	Deadline       string `json:"deadline" binding:"required"`
}

// UpdateJobRequest represents the request body for updating an existing job.
type UpdateJobRequest struct {
	Title          string    `json:"title,omitempty"`
	Description    string    `json:"description,omitempty"`
	Location       string    `json:"location,omitempty"`
	EmploymentType string    `json:"employment_type,omitempty"`
	Deadline       time.Time `json:"deadline,omitempty"`
}

type JobResponse struct {
	ID             uint      `json:"id,omitempty"`
	Title          string    `json:"title,omitempty"`
	Description    string    `json:"description,omitempty"`
	Location       string    `json:"location,omitempty"`
	EmploymentType string    `json:"employment_type,omitempty"`
	Deadline       time.Time `json:"deadline,omitempty"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
	UpdatedAt      time.Time `json:"updated_at,omitempty"`
}
