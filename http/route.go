package http

import (
	"github.com/buaazp/fasthttprouter"
	"github.com/kenny26/maynooth-api/repository/redis"
	"github.com/kenny26/maynooth-api/http/middleware"
	"github.com/kenny26/maynooth-api/http/controllers"
	"github.com/kenny26/maynooth-api/http/controllers/auth"
	"github.com/kenny26/maynooth-api/http/controllers/product"
	"github.com/kenny26/maynooth-api/http/controllers/category"
	"github.com/kenny26/maynooth-api/http/controllers/cart"
	"github.com/kenny26/maynooth-api/http/controllers/city"
	"github.com/kenny26/maynooth-api/http/controllers/order"
	cartV2 "github.com/kenny26/maynooth-api/http/controllers/v2/cart"
)

func ApiRoutes(rm *redis.RedisManager) *fasthttprouter.Router {
	router := fasthttprouter.New()

	router.GET("/", controllers.Index)

	router.GET("/auth/info", auth.AuthInfo(rm))
	router.POST("/auth/login", auth.AuthLogin(rm))
	router.POST("/auth/register", auth.AuthRegister(rm))
	router.POST("/auth/logout", auth.AuthLogout(rm))

	router.GET("/cities", city.ListCity)

	router.GET("/categories", category.ListCategory)
	router.GET("/categories/:categoryId", category.DetailCategory)

	router.GET("/products", product.ListProduct)
	router.GET("/products/:productId", product.DetailProduct)

	authMiddleware := middleware.AuthMiddleware(rm)

	router.GET("/carts/show", authMiddleware(cart.ShowCart))
	router.POST("/carts/add", authMiddleware(cart.AddCart))
	router.POST("/carts/clear", authMiddleware(cart.ClearCart))
	router.POST("/carts/remove", authMiddleware(cart.RemoveCartItem))
	router.POST("/carts/update", authMiddleware(cart.UpdateCartItem))

	router.POST("/orders/create", authMiddleware(order.SubmitOrder))

	router.POST("/v2/carts/store", authMiddleware(cartV2.StoreCart(rm)))
	router.GET("/v2/carts/show", authMiddleware(cartV2.ShowCart(rm)))

	return router
}
