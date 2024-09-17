package redis

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"log"
	"os"
)

var Rdb *redis.Client

func Init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	redisUrl, exists := os.LookupEnv("REDIS_URL")
	if !exists {
		log.Fatal("REDIS_URL is not set")
	}

	Rdb = redis.NewClient(&redis.Options{
		Addr:     redisUrl,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// Check connection
	ctx := context.Background()
	_, err = Rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}

	fmt.Println("Connected to Redis")
}

func GetClient() *redis.Client {
	return Rdb
}
