package models

type City struct {
	Model
	Name string `json:"name" gorm:"type:varchar(255)"`
	ShippingPrice float64 `json:"shippingPrice" gorm:"type:float"`
	IsActive bool `json:"isActive" gorm:"type:boolean"`
}

func GetActiveCities(cityIds []uint) (*[]City) {

	cities := &[]City{}
	tx := GetDB().Table("cities").Where("is_active = TRUE")

	if cityIds != nil && len(cityIds) > 0 {
		tx = tx.Where("id IN (?)", cityIds)
	}

	err := tx.Find(cities).Error

	if err != nil {
		return nil
	}
	return cities
}
