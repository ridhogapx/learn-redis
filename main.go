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

func main() {
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
