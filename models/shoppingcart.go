package models

import (
	"time"
	"log"
)

type ShoppingCart struct {
	Model
	Address string `json:"address" gorm:"type:varchar(255)"`
	CityID uint `json:"cityId" gorm:"type:int"`
	UserID uint `json:"userId" gorm:"type:int"`
	City City `json:"city"`
	ShoppingCartItems []ShoppingCartItem `json:"shoppingCartItems"`
}

func ClearShoppingCart(userId uint) (bool) {
	shoppingCarts := &[]ShoppingCart{}
	err := GetDB().Table("shopping_carts").
		Where("user_id = ?", userId).
		Find(shoppingCarts).Error

	var shoppingCartIds []uint
	for _, cart := range *shoppingCarts {
		shoppingCartIds = append(shoppingCartIds, cart.ID)
	}

	GetDB().Where("shopping_cart_id IN (?)", shoppingCartIds).Delete(ShoppingCartItem{})
	GetDB().Where("id IN (?)", shoppingCartIds).Delete(ShoppingCart{})

	if err != nil {
		return false
	}
	return true
}

func RemoveCartItem (userId uint, shoppingCartItemIds []uint) (bool) {

	shoppingCartItems := &[]ShoppingCartItem{}

	err := GetDB().Table("shopping_cart_items").
		Joins("JOIN shopping_carts on shopping_carts.id = shopping_cart_items.shopping_cart_id").
		Where("shopping_cart_items.id IN (?)", shoppingCartItemIds).
		Where("shopping_carts.user_id = ?", userId).Debug().
		Find(shoppingCartItems).Error

	var cartItemIds []uint

	for _, cartItem := range *shoppingCartItems {
		if cartItem.ID != 0 {
			cartItemIds = append(cartItemIds, cartItem.ID)
		}
	}

	if err == nil && cartItemIds != nil {
		GetDB().Where("id IN (?)", cartItemIds).Debug().Delete(ShoppingCartItem{})

		db.Debug().Exec(`
			UPDATE
				shopping_carts
			SET
				"deleted_at" = ?
			WHERE id IN (
				SELECT
					"shopping_carts".id
				FROM "shopping_carts"
				LEFT JOIN shopping_cart_items
					ON shopping_cart_items.shopping_cart_id = shopping_carts.id
					AND shopping_cart_items.deleted_at IS NULL
				WHERE "shopping_carts"."deleted_at" IS NULL
					AND shopping_carts.user_id = ?
					AND shopping_cart_items.id IS NULL
			)
		`, time.Now(), userId)
	}

	if err != nil {
		return false
	}
	return true
}


func UpdateCartItem (userId uint, shoppingCartItemId uint, quantity int) (bool) {
	shoppingCartItem := &ShoppingCartItem{}

	err := GetDB().Table("shopping_cart_items").
		Joins("JOIN shopping_carts on shopping_carts.id = shopping_cart_items.shopping_cart_id").
		Where("shopping_cart_items.id = ?", shoppingCartItemId).
		Where("shopping_carts.user_id = ?", userId).
		First(shoppingCartItem).Error

	if shoppingCartItem.ID != 0 {
		err = db.Model(shoppingCartItem).
			Update("quantity", quantity).Error
	}

	if err != nil {
		log.Println("error update", err)
		return false
	}
	log.Println("success update", err)
	return true
}

func GetShoppingCart(userId uint) (*[]ShoppingCart) {
	shoppingCarts := &[]ShoppingCart{}
	err := GetDB().Table("shopping_carts").
		Preload("City").
		Preload("ShoppingCartItems").
		Preload("ShoppingCartItems.ProductVariant").
		Preload("ShoppingCartItems.ProductVariant.Product").
		Preload("ShoppingCartItems.ProductVariant.Product.Category").
		Preload("ShoppingCartItems.ProductVariant.Product.ProductImages").
		Where("user_id = ?", userId).
		Find(shoppingCarts).Error

	if err != nil {
		return nil
	}
	return shoppingCarts
}
