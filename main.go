package main

import (
	"os"
	"log"
	"github.com/valyala/fasthttp"
	"github.com/kenny26/maynooth-api/repository/redis"
	"github.com/kenny26/maynooth-api/http"
	"github.com/joho/godotenv"
	"github.com/kenny26/maynooth-api/http/middleware"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")

	if port == "" {
		panic("Please define port number!")
	}

	// setup redis connection
	redisManager := redis.RedisManager{}
	redisManager.Connect(os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"), os.Getenv("REDIS_PASS"))

	router := http.ApiRoutes(&redisManager)

	log.Fatal(fasthttp.ListenAndServe(":"+port, middleware.CORS(router.Handler)))
}
