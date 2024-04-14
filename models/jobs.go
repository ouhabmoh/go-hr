package models

// Job represents a job posting.
type Job struct {
	BaseModel
	Title          string `gorm:"not null"`
	Description    string `gorm:"not null"`
	Location       string `gorm:"not null"`
	EmploymentType string `gorm:"not null"`
	Deadline       string `gorm:"not null"`

	RecruiterID  int
	Recruiter    User `gorm:"foreignKey:RecruiterID;references:id"`
	Applications []Application
}

// CreateJobRequest represents the request body for creating a new job.
type CreateJobRequest struct {
	Title          string `json:"title" binding:"required"`
	Description    string `json:"description" binding:"required"`
	Location       string `json:"location" binding:"required"`
	EmploymentType string `json:"employment_type" binding:"required"`
	Deadline       string `json:"deadline" binding:"required"`
}

// UpdateJobRequest represents the request body for updating an existing job.
type UpdateJobRequest struct {
	Title          string `json:"title,omitempty"`
	Description    string `json:"description,omitempty"`
	Location       string `json:"location,omitempty"`
	EmploymentType string `json:"employment_type,omitempty"`
	Deadline       string `json:"deadline,omitempty"`
}