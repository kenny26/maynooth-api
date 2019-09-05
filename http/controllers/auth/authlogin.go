package auth

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
	"github.com/kenny26/maynooth-api/repository/redis"
	"github.com/satori/go.uuid"
	"github.com/kenny26/maynooth-api/models"
)

const (
	sessionKey = "maynooth_token"
	userId     = "ID"
	username   = "NAME"
	userEmail  = "EMAIL"
	userGender = "GENDER"
	userDob = "DOB"
)

type AuthLoginBody struct {
	Username string `json:"username"`
	Password  string `json:"password"`
}

type AuthLoginResponse struct {
	ID	uint `json:"id"`
	Username	string `json:"username"`
	Email	string `json:"email"`
	Gender	string `json:"gender"`
	Dob	string `json:"dob"`
}

func AuthLogin(rm *redis.RedisManager) func(*fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {
		var param AuthLoginBody

		err := json.Unmarshal(ctx.PostBody(), &param)
		if err != nil {
			// If the structure of the body is wrong, return an HTTP error
			ctx.Error("BadRequest", fasthttp.StatusBadRequest)
			return
		}

		json.Unmarshal(ctx.PostBody(), &param)

		ctx.SetContentType("application/json")

		user := models.AuthenticateUser(param.Username, param.Password)

		if user != nil {
			sessionToken := uuid.NewV4().String()
			_, err := rm.Client.HMSet(sessionToken, map[string]interface{}{
				userId:    user.ID,
				username:  user.Username,
				userEmail: user.Email,
				userGender: user.Gender,
				userDob: user.Dob,
			}).Result()

			if err != nil {
				ctx.Error("InternalServerError", fasthttp.StatusInternalServerError)
				return
			}

			sessionCookie := fasthttp.Cookie{}
			sessionCookie.SetKey(sessionKey)
			sessionCookie.SetValue(sessionToken)
			sessionCookie.SetSecure(false)
			sessionCookie.SetPath("/")
			ctx.Response.Header.SetCookie(&sessionCookie)

			resJSON, _ := json.Marshal(AuthLoginResponse{
				ID: user.ID,
				Username:	user.Username,
				Email:	user.Email,
				Gender:	user.Gender,
				Dob:	user.Dob,
			})

			ctx.SetBody(resJSON)
			ctx.SetStatusCode(fasthttp.StatusOK)

			return
		}

		ctx.Error("InvalidCredential", fasthttp.StatusUnauthorized)
		return
	}
}
