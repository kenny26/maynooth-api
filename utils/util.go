package utils

import (
	"github.com/valyala/fasthttp"
	"encoding/json"
)

func Message(status bool, message string) (map[string]interface{}) {
	return map[string]interface{} {"status" : status, "message" : message}
}

func Respond(ctx *fasthttp.RequestCtx, data map[string] interface{})  {
	resJSON, _ := json.Marshal(data)

	ctx.SetContentType("application/json")
	ctx.SetBody(resJSON)
}
