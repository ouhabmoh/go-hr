package models

import "time"

// Application represents a job application.
type Application struct {
	BaseModel
	JobID       int    `gorm:"uniqueIndex:idx_job_candidate"`
	CandidateID uint   `gorm:"uniqueIndex:idx_job_candidate"`
	ResumeID    uint   `gorm:"not null"`
	Status      string `gorm:"not null"`
	Evaluation  *int   `gorm:"check:evaluation >= 0 AND evaluation <= 10"`
	Job         Job    `gorm:"foreignKey:JobID;references:id"`
	Candidate   User   `gorm:"foreignKey:CandidateID;references:id"`
	Resume      Resume `gorm:"foreignKey:ResumeID;references:id"`
}

// CreateApplicationRequest represents the request body for creating a new job application.
type CreateApplicationRequest struct {
	JobID int `uri:"jobID" binding:"required"`
}

// UpdateApplicationRequest represents the request body for updating an existing job application.
type UpdateApplicationRequest struct {
	Status     *string `json:"status,omitempty"`     // Allow updating application status
	Evaluation *int    `json:"evaluation,omitempty"` // Allow updating evaluation (if applicable)
}

type ApplicationResponse struct {
	ID          uint      `json:"id,omitempty"`
	JobID       int       `json:"job_id,omitempty"`
	CandidateID uint      `json:"candidate_id,omitempty"`
	Status      string    `json:"status,omitempty"`
	ResumeID    uint      `json:"resume_id,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}
