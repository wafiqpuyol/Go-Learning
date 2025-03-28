package myRedisClient

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func GetRedisClient(ctx context.Context) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // No password set
		DB:       0,  // Use default DB
		Protocol: 2,  // Connection protocol
	})
	defer client.Close()

	// Ping the server to check if it's reachable
	pong, err := client.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Ping:", pong)

	return client
}
