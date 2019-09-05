package redis

import (
	"log"
	"github.com/go-redis/redis"
)

type RedisManager struct {
	Client *redis.Client
}

func (rm *RedisManager) Connect(host, port, pass string) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: pass,
		DB:       0,
	})

	pong, err := redisClient.Ping().Result()

	if err != nil {
		log.Fatalln("Redis Error:", err)
	}

	log.Println("Redis connected:", pong)

	rm.Client = redisClient
}
