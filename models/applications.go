package models

// Application represents a job application.
type Application struct {
	BaseModel
	JobID       int `gorm:"not null"`
	CandidateID int `gorm:"not null"`

	Status     string `gorm:"not null"`
	Evaluation int    `gorm:"check:evaluation >= 1 AND evaluation <= 10"`
	Job        Job    `gorm:"foreignKey:JobID;references:id"`
	Candidate  User   `gorm:"foreignKey:CandidateID;references:id"`
}

// CreateApplicationRequest represents the request body for creating a new job application.
type CreateApplicationRequest struct {
	JobID       int `json:"job_id" binding:"required"`
	CandidateID int `json:"candidate_id" binding:"required"`
	// Optional fields (depending on your workflow)
	Evaluation *int `json:"evaluation,omitempty"` // Can be added during application submission or later
}

// UpdateApplicationRequest represents the request body for updating an existing job application.
type UpdateApplicationRequest struct {
	Status     *string `json:"status,omitempty"`     // Allow updating application status
	Evaluation *int    `json:"evaluation,omitempty"` // Allow updating evaluation (if applicable)
}
