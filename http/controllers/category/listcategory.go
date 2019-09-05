package category

import (
	"github.com/valyala/fasthttp"
	"github.com/kenny26/maynooth-api/models"
	"github.com/kenny26/maynooth-api/utils"
)

func ListCategory(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json")

	categories := models.GetActiveCategories()

	ctx.SetStatusCode(fasthttp.StatusOK)

	resp := utils.Message(true, "success")
	resp["data"] = categories
	utils.Respond(ctx, resp)
}
