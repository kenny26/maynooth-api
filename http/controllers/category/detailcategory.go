package category

import (
	"github.com/valyala/fasthttp"
	"github.com/kenny26/maynooth-api/models"
	"github.com/kenny26/maynooth-api/utils"
	"fmt"
	"strconv"
)

func DetailCategory(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json")

	categoryId, _ := strconv.Atoi(fmt.Sprintf("%v", ctx.UserValue("categoryId")))
	category := models.GetDetailCategory(categoryId)

	ctx.SetStatusCode(fasthttp.StatusOK)

	resp := utils.Message(true, "success")
	resp["data"] = category
	utils.Respond(ctx, resp)
}
