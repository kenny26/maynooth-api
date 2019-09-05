package models

import (
	"github.com/jinzhu/gorm"
	"log"
)

type Product struct {
	Model
	Name string `json:"name" gorm:"type:varchar(255)"`
	Description string `json:"description" gorm:"type:text"`
	Size string `json:"size" gorm:"type:varchar(255)"`
	Price float64 `json:"price" gorm:"type:float"`
	OriginalPrice float64 `json:"originalPrice" gorm:"type:float"`
	CategoryId uint `json:"categoryId" gorm:"type:int"`
	IsActive bool `json:"isActive" gorm:"type:boolean"`
	Category Category `json:"category,omitempty"`
	ProductImages []ProductImage `json:"images,omitempty"`
	ProductVariants []ProductVariant `json:"variants,omitempty"`
}

func ListActiveProducts(categoryId int, name string, order string) (*[]Product) {
	products := &[]Product{}

	tx := GetDB().Table("products").
			Preload("ProductImages", func(db *gorm.DB) *gorm.DB {
			return db.Where("is_active = TRUE")
		}).
			Preload("ProductVariants", func(db *gorm.DB) *gorm.DB {
			return db.Where("is_active = TRUE")
		}).
		Where("products.is_active = TRUE")

	if name != "" {
		tx = tx.Where("products.name ILIKE ?", "%"+name+"%")
	}
	if categoryId != 0 {
		tx = tx.Where("products.category_id = ?", categoryId)
	}

	if order == "createdAt" {
		tx = tx.Order("products.created_at DESC")
	}
	if order == "price" {
		tx = tx.Order("products.price")
	}

	err := tx.Find(products).Error

	if err != nil {
		log.Println("error query", err)
		return nil
	}
	return products
}

func GetDetailProduct(productId int) (*Product) {
	product := &Product{}
	err := GetDB().Table("products").Where("is_active = TRUE").
		Preload("ProductImages", func(db *gorm.DB) *gorm.DB {
			return db.Where("is_active = TRUE")
		}).
		Preload("ProductVariants", func(db *gorm.DB) *gorm.DB {
			return db.Where("is_active = TRUE")
		}).
		First(product, productId).Error
	if err != nil {
		return nil
	}

	return product
}
