package cart

import (
	"github.com/valyala/fasthttp"
	"github.com/kenny26/maynooth-api/models"
	"github.com/kenny26/maynooth-api/utils"
)

type CartItemResponse struct {
	ID	uint `json:"id"`
	Sku	string `json:"sku"`
	Quantity	uint `json:"quantity"`
	ProductVariantName	string `json:"name"`
	ProductName	string `json:"productName"`
	ProductPrice	float64 `json:"productPrice"`
	ProductCategory	string `json:"productCategory"`
	ProductImageUrl	string `json:"productImageUrl"`
}

type CartResponse struct {
	ID	uint `json:"id"`
	Address	string `json:"address"`
	City	string `json:"city"`
	Items []CartItemResponse `json:"items"`
}

func ShowCart(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json")

	userId := ctx.UserValue("ID").(uint)

	shoppingCarts := models.GetShoppingCart(userId)

	var carts []CartResponse

	for _, cart := range *shoppingCarts {
		if cart.ID != 0 {
			parsedCart := CartResponse{}
			parsedCart.ID = cart.ID
			parsedCart.Address = cart.Address
			parsedCart.City = cart.City.Name

			var cartItems []CartItemResponse

			for _, cartItem := range cart.ShoppingCartItems {
				if cartItem.ID != 0 {
					parsedCartItem := CartItemResponse{}
					parsedCartItem.ID = cartItem.ID
					parsedCartItem.Quantity = cartItem.Quantity
					parsedCartItem.Sku = cartItem.ProductVariant.Sku
					parsedCartItem.ProductVariantName = cartItem.ProductVariant.Name
					parsedCartItem.ProductName = cartItem.ProductVariant.Product.Name
					parsedCartItem.ProductPrice = cartItem.ProductVariant.Product.Price
					parsedCartItem.ProductCategory = cartItem.ProductVariant.Product.Category.Name

					for _, image := range cartItem.ProductVariant.Product.ProductImages {
						if image.ID != 0 && image.IsThumbnail {
							parsedCartItem.ProductImageUrl = image.Url
						}
					}

					cartItems = append(cartItems, parsedCartItem)
				}
			}

			parsedCart.Items = cartItems

			carts = append(carts, parsedCart)
		}
	}

	resp := utils.Message(true, "success")
	resp["data"] = carts
	utils.Respond(ctx, resp)
	return
}
