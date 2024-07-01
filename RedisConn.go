package main

import (
	"fmt"
	"reflect"

	"github.com/go-redis/redis"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	result, error := client.Ping().Result()
	fmt.Println(result)
	fmt.Println(error)
	fmt.Println(client.Get("name"))
	fmt.Println(client.Keys("*"))
	a := client.LRange("address", 0, 2)
	fmt.Println(reflect.TypeOf(a.Val()))
	fmt.Println(reflect.TypeOf(client.HGetAll("Employee")))
	counter := 0
	for i := 0; i <= 10; i++ {
		client.LPush("", counter)
		counter++
	}

}
