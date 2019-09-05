package models

type ShoppingCartItem struct {
	Model
	Quantity uint `json:"quantity" gorm:"type:int"`
	ProductVariantID uint `json:"productVariantId" gorm:"type:int"`
	ShoppingCartID uint `json:"shoppingCartId" gorm:"type:int"`
	ProductVariant ProductVariant `json:"productVariant"`
}
