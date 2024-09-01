package cache

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var client *redis.Client

func Connect() {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	p := client.Conn().Ping(context.Background())
	fmt.Printf("Redis connection is ready: %s", p.Val())
	fmt.Println("")
}
