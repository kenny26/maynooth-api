package controllers

import (
	"encoding/json"

	"github.com/valyala/fasthttp"
)

// Index Api Handler
func Index(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(fasthttp.StatusOK)

	res := struct {
		Status string `json:"status"`
		Live   bool   `json:"isAlive"`
	}{
		Status: "OK",
		Live:   true,
	}
	resJSON, _ := json.Marshal(res)

	ctx.SetBody(resJSON)
}
