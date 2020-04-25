package services

import (
	"github.com/go-redis/redis/v7"
	"log"
)

var RedisClient *redis.Client

func ConnectRedis(address, password string) {
	RedisClient = redis.NewClient(&redis.Options{Addr: address, Password: password})
	_, err := RedisClient.Ping().Result()
	if err != nil {
		log.Fatal(err)
	}
}
