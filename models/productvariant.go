package models

type ProductVariant struct {
	Model
	Sku string `json:"sku" gorm:"type:varchar(255)"`
	Name string `json:"name" gorm:"type:varchar(255)"`
	Stock int `json:"stock" gorm:"type:int"`
	IsActive bool `json:"isActive" gorm:"type:boolean"`
	ProductID uint `json:"productId" gorm:"type:int"`
	Product Product `json:"product"`
}

func GetActiveProductVariants(variantIds []uint) (*[]ProductVariant) {

	variants := &[]ProductVariant{}
	tx := GetDB().Table("product_variants").
		Preload("Product").
		Where("is_active = TRUE")

	if len(variantIds) > 0 {
		tx = tx.Where("id IN (?)", variantIds)
	}

	err := tx.Find(variants).Error

	if err != nil {
		return nil
	}
	return variants
}
