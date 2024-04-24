package models

type Resume struct {
	BaseModel

	CandidateID uint   `gorm:"not null"`
	Filename    string `gorm:"not null"`
	Candidate   User   `gorm:"foreignKey:CandidateID;references:id"`
}

// type CreateResumeRequest struct {
// 	CandidateID uint `json`
// }
