package auth

import (
	"github.com/kenny26/maynooth-api/repository/redis"
	"github.com/valyala/fasthttp"
)

func AuthLogout(rm *redis.RedisManager) func(*fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.SetContentType("application/json")

		token := ctx.Request.Header.Cookie(sessionKey)

		exist := rm.Client.Exists(string(token)).Val()
		if exist != 1 {
			ctx.SetStatusCode(fasthttp.StatusOK)
			return
		}

		rm.Client.Del(string(token))
		ctx.Response.Header.DelCookie(sessionKey)

		ctx.SetStatusCode(fasthttp.StatusOK)
	}
}
