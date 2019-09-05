package product

import (
	"github.com/valyala/fasthttp"
	"github.com/kenny26/maynooth-api/models"
	"github.com/kenny26/maynooth-api/utils"
	"fmt"
	"strconv"
)

func DetailProduct(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json")

	productId, _ := strconv.Atoi(fmt.Sprintf("%v", ctx.UserValue("productId")))
	product := models.GetDetailProduct(productId)

	ctx.SetStatusCode(fasthttp.StatusOK)

	resp := utils.Message(true, "success")
	resp["data"] = product
	utils.Respond(ctx, resp)
}
