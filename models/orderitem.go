package models

type OrderItem struct {
	Model
	ProductName string `json:"productName" gorm:"type:varchar(255)"`
	ProductVariantName string `json:"productVariantName" gorm:"type:varchar(255)"`
	ProductPrice float64 `json:"productPrice" gorm:"type:float"`
	Quantity int `json:"quantity" gorm:"type:int"`
	OrderDetailID uint `json:"orderDetailId" gorm:"type:int"`
	ProductVariantID uint `json:"productVariantId" gorm:"type:int"`
}
