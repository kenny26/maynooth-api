package models

type ProductImage struct {
	Model
	Url string `json:"url" gorm:"type:varchar(255)"`
	IsActive bool `json:"isActive" gorm:"type:boolean"`
	IsThumbnail bool `json:"isThumbnail" gorm:"type:boolean"`
	ProductID uint `json:"productId" gorm:"type:int"`
}
