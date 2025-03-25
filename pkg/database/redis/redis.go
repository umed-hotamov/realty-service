package redis

import (
	"log"
  
  "github.com/go-redis/redis/v7"
)

func NewClient() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
	  Addr: "localhost:6379",
  })
  
	_, err := client.Ping().Result()
	if err != nil {
    log.Fatalf("failed to ping storage: %s", err.Error())
		return nil, err
	}

	return client, nil
}
