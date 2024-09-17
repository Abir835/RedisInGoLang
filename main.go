package main

import (
	"context"
	"fmt"
	"github.com/RedisInGoLang/redis"
	_ "github.com/RedisInGoLang/redis"
	"log"
)

func main() {
	// Initialize Redis connection
	redis.Init()

	// Get the Redis client
	client := redis.GetClient()
	ctx := context.Background()

	// Set a key-value pair
	err := client.Set(ctx, "foo", "bar", 0).Err()
	if err != nil {
		log.Fatalf("Could not set key: %v", err)
	}

	// Get the value of the key
	val, err := client.Get(ctx, "foo").Result()
	if err != nil {
		log.Fatalf("Could not get key: %v", err)
	}
	fmt.Println("foo", val)
}
