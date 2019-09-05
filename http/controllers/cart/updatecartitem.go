package cart

import (
	"github.com/valyala/fasthttp"
	"github.com/kenny26/maynooth-api/models"
	"github.com/kenny26/maynooth-api/utils"
	"encoding/json"
)

type UpdateCartBody struct {
	ShoppingCartItemID uint `json:"shoppingCartItemId"`
	Quantity int `json:"quantity"`
}

func UpdateCartItem (ctx *fasthttp.RequestCtx) {
	var param UpdateCartBody

	err := json.Unmarshal(ctx.PostBody(), &param)
	if err != nil {
		// If the structure of the body is wrong, return an HTTP error
		ctx.Error("BadRequest", fasthttp.StatusBadRequest)
		return
	}

	json.Unmarshal(ctx.PostBody(), &param)

	userId := ctx.UserValue("ID").(uint)

	models.UpdateCartItem(userId, param.ShoppingCartItemID, param.Quantity)

	resp := utils.Message(true, "success")
	utils.Respond(ctx, resp)
	return
}
