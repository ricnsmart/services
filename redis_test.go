package services

import (
	"log"
	"testing"
)

const (
	redisAddress  = ""
	redisPassword = ""
)

func TestConnectRedis(t *testing.T) {

	if err := ConnectRedis(redisAddress, redisPassword); err != nil {
		t.Errorf("ConnectRedis() error = %v", err)
	}

	err := RedisClient.HSet("test2", 5, []byte("1234556")).Err()
	if err != nil {
		log.Fatal(err)
	}
}
