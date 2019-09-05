package middleware

import (
	"github.com/valyala/fasthttp"
	"log"
	"strconv"
	"github.com/kenny26/maynooth-api/models"
	"github.com/kenny26/maynooth-api/repository/redis"
)

type Middleware func(h fasthttp.RequestHandler) fasthttp.RequestHandler

func AuthMiddleware(rm *redis.RedisManager) Middleware {
	return func(h fasthttp.RequestHandler) fasthttp.RequestHandler {
		return func(ctx *fasthttp.RequestCtx) {
			token := ctx.Request.Header.Cookie("maynooth_token")

			exist := rm.Client.Exists(string(token)).Val()
			if exist != 1 {
				ctx.SetStatusCode(fasthttp.StatusUnauthorized)
				return
			}

			res, err := rm.Client.HGetAll(string(token)).Result()
			if err != nil {
				log.Println(err)
			}

			userId, _ := strconv.Atoi(res["ID"])

			user := models.GetUserById(userId)

			ctx.SetUserValue("ID", user.ID)
			ctx.SetUserValue("Username", user.Username)
			ctx.SetUserValue("Email", user.Email)

			if user != nil {
				h(ctx)
			} else {
				ctx.Response.SetStatusCode(fasthttp.StatusUnauthorized)
			}
		}
	}
}
