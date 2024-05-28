package database

import (
	"context"
	"github.com/go-redis/redis/v8"
	"os"
)

var Context = context.Background()

func CreateClient(dbNo int) *redis.Client {

	redisDb := redis.NewClient(&redis.Options{

		Addr:     os.Getenv("DB_ADDR"),
		Password: os.Getenv("DB_PASS"),
		DB:       dbNo,
	})

	return redisDb
}
