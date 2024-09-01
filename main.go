package main

import (
	"context"
	"fmt"
	"redis-lock/cache"
	"strconv"
	"time"
)

func main() {
	ctx := context.Background()
	cache.Connect()

	start := time.Now()
	fmt.Println("---------------")
	fmt.Printf("Start: %s", start)

	var keys []string
	var values []interface{}

	for i := 0; i < 1000000; i++ {
		keys = append(keys, strconv.Itoa(i+1))
		values = append(values, strconv.Itoa(i+1), strconv.Itoa(i+1))
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
		fmt.Printf("A value has been found, %s", foundValues)
	} else {
		fmt.Printf("values, %s", values...)
		err = cache.SetAll(ctx, values...)
		if err != nil {
			fmt.Println("")
			fmt.Printf("Error: %v", err)
		}
	}
	fmt.Println("---------------")
	fmt.Printf("Finish: %s", time.Now())
	fmt.Println("")
	since := time.Since(start)
	fmt.Printf("Took time: %s", since)
	fmt.Println("---------------")
}
