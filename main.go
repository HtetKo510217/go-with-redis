package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func main() {
	// Connect to Redis
	opt, err := redis.ParseURL("redis://localhost:6379")
	if err != nil {
		panic(err)
	}
	client := redis.NewClient(opt)
	defer client.Close()

	// Create
	err = client.Set(context.Background(), "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}
	fmt.Println("Key-value pair created")

	// Read
	val, err := client.Get(context.Background(), "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Retrieved value:", val)

	// Update
	err = client.Set(context.Background(), "key", "new-value", 0).Err()
	if err != nil {
		panic(err)
	}
	fmt.Println("Key-value pair updated")

	// Delete
	deleted, err := client.Del(context.Background(), "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Key deleted:", deleted)
}
