package cart

import (
	"github.com/valyala/fasthttp"
	"github.com/kenny26/maynooth-api/models"
	"github.com/kenny26/maynooth-api/utils"
)

func ClearCart(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json")

	userId := ctx.UserValue("ID").(uint)

	models.ClearShoppingCart(userId)

	resp := utils.Message(true, "success")
	utils.Respond(ctx, resp)
	return
}
