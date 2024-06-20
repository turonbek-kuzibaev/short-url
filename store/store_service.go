package store

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type StorageSerice struct {
	RedisClient *redis.Client
}

var (
	storageSerice = &StorageSerice{}
	ctx           = context.Background()
)

const CacheDuration = 6 * time.Hour

func InitializeStore() *StorageSerice {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := redisClient.Ping(ctx).Result()

	if err != nil {
		panic(fmt.Sprintf("Error init Redis: %v", err))
	}

	fmt.Printf("\nRedis started successfully: pong message = {%s}", pong)
	storageSerice.RedisClient = redisClient
	return storageSerice
}

func SaveUrlMapping(shortUrl string, originalUrl string, userId string) {
	err := storageSerice.RedisClient.Set(ctx, shortUrl, originalUrl, CacheDuration)

	if err != nil {
		panic(fmt.Sprintf("Failed saving key url | Error: %v - shortUrl: %s - originalUrl: %s\n", err, shortUrl, originalUrl))
	}
}

func RetrieveInitialUrl(shortUrl string) string {
	result, err := storageSerice.RedisClient.Get(ctx, shortUrl).Result()
	if err != nil {
		panic(fmt.Sprintf("Failed RetrieveInitialUrl url | Error: %v - shortUrl: %s\n", err, shortUrl))
	}
	return result
}
