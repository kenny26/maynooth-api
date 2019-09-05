package models

type Order struct {
	Model
	Number string `json:"number" gorm:"type:varchar(255)"`
	Amount float64 `json:"amount" gorm:"type:float"`
	Status string `json:"status" gorm:"type:varchar(255)"`
	UserID uint `json:"userId" gorm:"type:int"`
}

