package main

import (
	"github.com/go-redis/redis"
)

var client = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "",
	DB:       0,
})

func Retrieve() (*string, error) {
	val, err := client.Get("name").Result()

	if err != nil {
		return nil, err
	}

	return &val, nil
}
