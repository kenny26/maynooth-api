package cart

import (
	"github.com/valyala/fasthttp"
	"github.com/kenny26/maynooth-api/models"
	"github.com/kenny26/maynooth-api/utils"
	"encoding/json"
)

type RemoveCartBody struct {
	ShoppingCartItemIDs []uint `json:"shoppingCartItemIds"`
}

func RemoveCartItem (ctx *fasthttp.RequestCtx) {
	var param RemoveCartBody

	err := json.Unmarshal(ctx.PostBody(), &param)
	if err != nil {
		// If the structure of the body is wrong, return an HTTP error
		ctx.Error("BadRequest", fasthttp.StatusBadRequest)
		return
	}

	json.Unmarshal(ctx.PostBody(), &param)

	userId := ctx.UserValue("ID").(uint)

	models.RemoveCartItem(userId, param.ShoppingCartItemIDs)

	resp := utils.Message(true, "success")
	utils.Respond(ctx, resp)
	return
}
