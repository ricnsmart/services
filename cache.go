package services

import (
	"github.com/go-redis/redis/v7"
	"log"
	"time"
)

type cache struct{}

var Cache cache

var redisClient *redis.Client

func InitCache(address, password string) {
	redisClient = redis.NewClient(&redis.Options{Addr: address, Password: password})
	_, err := redisClient.Ping().Result()
	if err != nil {
		log.Fatal(err)
	}
}

func (cache) Set(key string, value interface{}, expiration time.Duration) error {
	return redisClient.Set(key, value, expiration).Err()
}

func (cache) Get(key string) (string, error) {
	return redisClient.Get(key).Result()
}

func (cache) Del(key string) error {
	return redisClient.Del(key).Err()
}

func (cache) GetBytes(key string) ([]byte, error) {
	return redisClient.Get(key).Bytes()
}

func (cache) HSet(key string, values ...interface{}) error {
	return redisClient.HSet(key, values).Err()
}
