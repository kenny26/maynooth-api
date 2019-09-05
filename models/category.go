package models

import "github.com/jinzhu/gorm"

type Category struct {
	Model
	Name string `json:"name" gorm:"type:varchar(255)"`
	IsActive bool `json:"isActive" gorm:"type:boolean"`
	ImageUrl string `json:"imageUrl" gorm:"type:varchar(255)"`
	CategoryDetails []CategoryDetail `json:"details"`
}

func GetActiveCategories() (*[]Category) {

	categories := &[]Category{}
	err := GetDB().Table("categories").Where("is_active = TRUE").Find(categories).Error
	if err != nil {
		return nil
	}
	return categories
}

func GetDetailCategory(categoryId int) (*Category) {
	category := &Category{}
	err := GetDB().Table("categories").Where("is_active = TRUE").
		Preload("CategoryDetails", func(db *gorm.DB) *gorm.DB {
			return db.Where("is_active = TRUE")
		}).
		First(category, categoryId).Error
	if err != nil {
		return nil
	}
	return category
}
