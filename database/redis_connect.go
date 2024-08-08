package database

import (
	"context"
	"lmizania/config"
	"log"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func RedisDBInstance() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     config.REDIS_ADDR,
		Password: config.REDIS_PASSWORD,
		Username: config.REDIS_USERNAME,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to Redis successfully!")
	return client
}
