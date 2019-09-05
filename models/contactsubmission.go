package models

type ContactSubmission struct {
	Model
	Name string `json:"name" gorm:"type:varchar(255)"`
	Email string `json:"email" gorm:"type:varchar(255)"`
	Description string `json:"description" gorm:"type:text"`
}
