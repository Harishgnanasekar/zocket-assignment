package cache

import (
	"github.com/redis/go-redis/v9"
	"log"
	"context"
	"github.com/spf13/viper"
)

var RedisClient *redis.Client

func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr: viper.GetString("REDIS_HOST"),
	})
	if err := RedisClient.Ping(context.Background()).Err(); err != nil {
		log.Fatalf("Error connecting to Redis: %v", err)
	}
	log.Println("Redis connected")
}
