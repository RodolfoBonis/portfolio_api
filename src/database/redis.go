package database

import (
	"github.com/go-redis/redis/v7"
	"os"
)

var Client *redis.Client

func InitRedis() {
	dsn := os.Getenv("REDIS_DSN")

	Client = redis.NewClient(&redis.Options{
		Addr: dsn, //redis port
	})

	_, err := Client.Ping().Result()
	if err != nil {
		panic(err)
	}
}
