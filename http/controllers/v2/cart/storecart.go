package cart

import (
	"log"
	"encoding/json"
	"github.com/valyala/fasthttp"
	"github.com/kenny26/maynooth-api/utils"
	"github.com/kenny26/maynooth-api/repository/redis"
	"fmt"
	"strconv"
)

type AddCartItemBody struct {
	ProductID uint `json:"productId"`
	VariantID uint `json:"variantId"`
	Quantity uint `json:"quantity"`
}

type AddCartBody struct {
	CityID int `json:"cityId"`
	Items []AddCartItemBody `json:"items"`
}

func StoreCart(rm *redis.RedisManager) func(*fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.SetContentType("application/json")

		var param AddCartBody

		err := json.Unmarshal(ctx.PostBody(), &param)
		if err != nil {
			// If the structure of the body is wrong, return an HTTP error
			ctx.Error("BadRequest", fasthttp.StatusBadRequest)
			return
		}

		userId := ctx.UserValue("ID").(uint)

		var userCartKey = "cart-user-" + fmt.Sprint(userId)
		res, err := rm.Client.HGetAll(userCartKey).Result()

		log.Println("existing cart", res)
		if err != nil {
			log.Println(err)
		}
		log.Println("key cart", userCartKey)

		cartItemValue, _ := json.Marshal(param.Items)
		cartCityValue, _ := strconv.Atoi(res["cityId"])

		if param.CityID != 0 {
			cartCityValue = param.CityID
		}

		_, err = rm.Client.HMSet(userCartKey, map[string]interface{}{
			"items": cartItemValue,
			"cityId": cartCityValue,
		}).Result()

		if err != nil {
			log.Println(err)
		}

		resp := utils.Message(true, "success")
		resp["data"] = param
		utils.Respond(ctx, resp)
		return
	}
}
