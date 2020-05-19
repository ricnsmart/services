package cache

import (
	"github.com/go-redis/redis/v7"
)

var Client *redis.Client

func Connect(address, password string) error {
	Client = redis.NewClient(&redis.Options{Addr: address, Password: password})
	_, err := Client.Ping().Result()
	return err
}
