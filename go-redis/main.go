package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func Example() {
	redisConnObject := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", //No Password set
		DB:       0,  //Use default DB
	})

	errResponse := redisConnObject.Set(ctx, "Object1", "Value", 1).Err()
	if errResponse != nil {
		panic(errResponse)
	}
	fmt.Println("Exiting Redis Program.")
}

func main() {
	Example()
}
