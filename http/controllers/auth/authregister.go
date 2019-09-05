package auth

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
	"github.com/kenny26/maynooth-api/repository/redis"
	"github.com/satori/go.uuid"
	"github.com/kenny26/maynooth-api/models"
)

type AuthRegisterBody struct {
	Username string `json:"username"`
	Email string `json:"email"`
	Password  string `json:"password"`
	Gender string `json:"gender"`
	Dob string `json:"dob"`
}

type AuthRegisterResponse struct {
	ID	uint `json:"id"`
	Username	string `json:"name"`
	Email	string `json:"email"`
	Gender	string `json:"gender"`
	Dob	string `json:"dob"`
}

func AuthRegister(rm *redis.RedisManager) func(*fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.SetContentType("application/json")

		var param AuthRegisterBody

		err := json.Unmarshal(ctx.PostBody(), &param)
		if err != nil {
			// If the structure of the body is wrong, return an HTTP error
			ctx.Error("BadRequest", fasthttp.StatusBadRequest)
			return
		}

		user := &models.User{
			Username: param.Username,
			Password: param.Password,
			Email: param.Email,
			Gender: param.Gender,
			Dob: param.Dob,
		}
		newUser := user.Create()

		if newUser == nil {
			ctx.Error("FailCreateUser", fasthttp.StatusBadRequest)
			return
		}

		sessionToken := uuid.NewV4().String()
		_, err = rm.Client.HMSet(sessionToken, map[string]interface{}{
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

		resJSON, _ := json.Marshal(AuthRegisterResponse{
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
}
