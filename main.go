package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	ping, err := client.Ping().Result()

	if err != nil {
		panic(err)
	}

	fmt.Println(ping)

	err = client.Set("name", "Raymond", 0).Err()

	if err != nil {
		panic(err)
	}

	fmt.Println("Setting values is success!")

}
