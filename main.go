package main

import (
	"context"
	"fmt"
	"redis-lock/cache"
	"strconv"
	"time"

	"github.com/google/uuid"
)

func main() {
	ctx := context.Background()
	cache.Connect()

	start := time.Now()
	fmt.Println("---------------")
	fmt.Printf("Start: %s", start)

	lock(ctx)
	// unlock(ctx)

	fmt.Println("---------------")
	fmt.Printf("Finish: %s", time.Now())
	fmt.Println("")
	since := time.Since(start)
	fmt.Printf("Time taken: %s", since)
	fmt.Println("")
	fmt.Println("---------------")
}

func lock(ctx context.Context) {
	var keys []string
	var values []interface{}

	for i := 0; i < 3000000; i++ {
		keys = append(keys, strconv.Itoa(i+1))
		values = append(values, uuid.New().String(), "{\"time\": \""+time.Now().Format(time.RFC3339Nano)+"\", \"status\":\"LOCKED\",\"by\":\"RPaaS\"}")
	}
	res, err := cache.GetAll(ctx, keys...)
	if err != nil {
		fmt.Println("")
		fmt.Printf("Error: %v", err)
	}
	var foundValues []interface{}
	for _, v := range res {
		if v != nil {
			foundValues = append(foundValues, v)
		}
	}
	if len(foundValues) > 0 {
		fmt.Println("")
		fmt.Println("A value has been found")
	} else {
		err = cache.SetAll(ctx, values...)
		if err != nil {
			fmt.Println("")
			fmt.Printf("Error: %v", err)
		}
	}
}

func unlock(ctx context.Context) {
	var unlockKeys []string
	for i := 0; i < 300; i++ {
		unlockKeys = append(unlockKeys, strconv.Itoa(i+1))
	}
	err := cache.RemoveAll(ctx, unlockKeys...)
	if err != nil {
		fmt.Println("")
		fmt.Printf("Error removing: %v", err)
	}
}
