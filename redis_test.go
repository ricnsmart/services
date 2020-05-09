package services

import (
	"log"
	"testing"
)

func TestConnectRedis(t *testing.T) {

	if err := ConnectRedis("dev.ricnsmart.com:10032", "13c7a45a0d9d"); err != nil {
		t.Errorf("ConnectRedis() error = %v", err)
	}

	err := RedisClient.HSet("test2", 5, []byte("1234556")).Err()
	if err != nil {
		log.Fatal(err)
	}
}
