package city

import (
	"github.com/valyala/fasthttp"
	"github.com/kenny26/maynooth-api/models"
	"github.com/kenny26/maynooth-api/utils"
)

func ListCity(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json")

	cities := models.GetActiveCities(nil)

	ctx.SetStatusCode(fasthttp.StatusOK)

	resp := utils.Message(true, "success")
	resp["data"] = cities
	utils.Respond(ctx, resp)
}
