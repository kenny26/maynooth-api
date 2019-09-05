package cart

import (
	"github.com/valyala/fasthttp"
	"github.com/kenny26/maynooth-api/utils"
	"encoding/json"
	"fmt"
	"github.com/kenny26/maynooth-api/repository/redis"
	"strconv"
	"github.com/kenny26/maynooth-api/models"
)

type CartItemResponse struct {
	Sku	string `json:"sku"`
	Quantity	uint `json:"quantity"`
	ProductID	uint `json:"productId"`
	ProductVariantID	uint `json:"productVariantId"`
	ProductVariantName	string `json:"name"`
	ProductName	string `json:"productName"`
	ProductPrice	float64 `json:"productPrice"`
	ProductCategory	string `json:"productCategory"`
	ProductImageUrl	string `json:"productImageUrl"`
}

type CartResponse struct {
	CityID	int `json:"cityId"`
	Items []CartItemResponse `json:"items"`
}

type CartItemBody struct {
	ProductID uint `json:"productId"`
	VariantID uint `json:"variantId"`
	Quantity uint `json:"quantity"`
}

func ShowCart(rm *redis.RedisManager) func(*fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.SetContentType("application/json")

		userId := ctx.UserValue("ID").(uint)

		var userCartKey= "cart-user-" + fmt.Sprint(userId)
		res, err := rm.Client.HGetAll(userCartKey).Result()

		var cartRedis []CartItemBody

		err = json.Unmarshal([]byte(res["items"]), &cartRedis)
		if err != nil {
			// If the structure of the body is wrong, return an HTTP error
			ctx.Error("BadRequest", fasthttp.StatusBadRequest)
			return
		}

		var cart CartResponse

		cityId, _ := strconv.Atoi(res["cityId"])
		cart.CityID = cityId

		var variantIds []uint

		for _, cart := range cartRedis {
			variantIds = append(variantIds, cart.VariantID)
		}

		variants := models.GetActiveProductVariants(variantIds)

		var itemResponse []CartItemResponse

		for _, cart := range cartRedis {
			var item CartItemResponse

			item.Quantity = cart.Quantity

			for _, variant := range *variants {
				if variant.ID == cart.VariantID {
					item.Sku = variant.Product.Name
					item.ProductID = variant.Product.ID
					item.ProductVariantID = variant.ID
					item.ProductVariantName = variant.Name
					item.ProductName = variant.Product.Name
					item.ProductPrice = variant.Product.Price
					item.ProductCategory = variant.Product.Category.Name

					for _, image := range variant.Product.ProductImages {
						if image.ID != 0 && image.IsThumbnail {
							item.ProductImageUrl = image.Url
						}
					}

					itemResponse = append(itemResponse, item)
					break
				}
			}
		}

		cart.Items = itemResponse

		resp := utils.Message(true, "success")
		resp["data"] = cart
		utils.Respond(ctx, resp)
		return
	}
}
