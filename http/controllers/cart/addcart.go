package cart

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
	"github.com/kenny26/maynooth-api/models"
	"github.com/kenny26/maynooth-api/utils"
)

type AddCartItemBody struct {
	VariantID uint `json:"variantId"`
	Quantity uint `json:"quantity"`
}

type AddCartBody struct {
	Address string `json:"address"`
	CityID uint `json:"cityId"`
	ProductID uint `json:"productId"`
	Items []AddCartItemBody `json:"items"`
}

func AddCart(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json")

	var param AddCartBody

	err := json.Unmarshal(ctx.PostBody(), &param)
	if err != nil {
		// If the structure of the body is wrong, return an HTTP error
		ctx.Error("BadRequest", fasthttp.StatusBadRequest)
		return
	}

	shoppingCart := &models.ShoppingCart{
		Address: param.Address,
		CityID: param.CityID,
		UserID: ctx.UserValue("ID").(uint),
	}
	models.GetDB().Create(&shoppingCart)

	for _, item := range param.Items {
		shoppingCartItem := &models.ShoppingCartItem{
			Quantity: item.Quantity,
			ProductVariantID: item.VariantID,
			ShoppingCartID: shoppingCart.ID,
		}
		models.GetDB().Create(&shoppingCartItem)
	}

	resp := utils.Message(true, "success")
	resp["data"] = shoppingCart
	utils.Respond(ctx, resp)
	return
}
