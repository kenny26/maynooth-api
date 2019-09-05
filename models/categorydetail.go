package models

type CategoryDetail struct {
	Model
	Name string `json:"name" gorm:"type:varchar(255)"`
	Company string `json:"company" gorm:"type:varchar(255)"`
	Description string `json:"description" gorm:"type:varchar(255)"`
	ImageUrl string `json:"imageUrl" gorm:"type:varchar(255)"`
	IsActive bool `json:"isActive" gorm:"type:boolean"`
	CategoryId uint `json:"categoryId" gorm:"type:int"`
}
