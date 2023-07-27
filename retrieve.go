package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

var client = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "",
	DB:       0,
})

func Retrieve() {
	val, err := client.Get("name").Result()

	if err != nil {
		panic(err)
	}

	fmt.Println(val)
}
