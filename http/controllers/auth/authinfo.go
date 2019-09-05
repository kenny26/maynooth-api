package auth

import (
	"github.com/kenny26/maynooth-api/repository/redis"
	"github.com/valyala/fasthttp"
	"encoding/json"
	"log"
	"strconv"
	"github.com/kenny26/maynooth-api/models"
)

type AuthInfoResponse struct {
	ID	int `json:"id"`
	Username	string `json:"username"`
	Email	string `json:"email"`
	Gender	string `json:"gender"`
	Dob	string `json:"dob"`
}

func AuthInfo(rm *redis.RedisManager) func(*fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.SetContentType("application/json")

		token := ctx.Request.Header.Cookie(sessionKey)

		exist := rm.Client.Exists(string(token)).Val()
		log.Println("exist", string(token))
		if exist != 1 {
			ctx.SetStatusCode(fasthttp.StatusUnauthorized)
			return
		}

		res, err := rm.Client.HGetAll(string(token)).Result()

		if err != nil {
			log.Println(err)
		}

		userId, _ := strconv.Atoi(res[userId])

		user := models.GetUserById(userId)

		if user != nil {
			resJSON, _ := json.Marshal(AuthInfoResponse{
				ID: userId,
				Username: user.Username,
				Email:	user.Email,
				Gender: user.Gender,
				Dob: user.Dob,
			})

			ctx.SetBody(resJSON)
			ctx.SetStatusCode(fasthttp.StatusOK)
		}
	}
}
