package services

import (
	"github.com/go-redis/redis/v7"
)

var RedisClient *redis.Client

func ConnectRedis(address, password string) error {
	RedisClient = redis.NewClient(&redis.Options{Addr: address, Password: password})
	_, err := RedisClient.Ping().Result()
	return err
}
