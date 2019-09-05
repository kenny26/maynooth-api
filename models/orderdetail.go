package models

type OrderDetail struct {
	Model
	ShippingPrice float64 `json:"shippingPrice" gorm:"type:float"`
	Status string `json:"status" gorm:"type:varchar(255)"`
	CityID uint `json:"cityId" gorm:"type:int"`
	OrderID uint `json:"orderId" gorm:"type:int"`
}
