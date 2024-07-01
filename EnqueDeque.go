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
	fmt.Println(client.Ping())
	enque(client)
	dequeue(client)
}

func enque(client *redis.Client) {
	for i := 0; i <= 10; i++ {
		client.LPush("queue", fmt.Sprint("DATA:", i))
		client.LPush("reliableQueue", fmt.Sprint("DATA:", i))
	}
	fmt.Println("10 entries is Enqued in 'queue' and 'reliableQueue'.")
}

func dequeue(client *redis.Client) {
	//removing from simple queue.
	for i := 0; i <= 10; i++ {
		fmt.Println(client.RPop("queue"))
	}

	//Reliable queue
	// a task is not removed from the queue immediately when it is dequeued.
	// Instead, it is moved to a temporary queue where it is stored until
	// the consumer confirms that the task has been processed
	for i := 0; i <= 10; i++ {
		//RPOPLPUSH source destination
		client.RPopLPush("reliableQueue", "tempQueue")
	}

	fmt.Println("queue and reliableQueue is Empty now")
	fmt.Println("Data is still present in the tempQueue. ->")
	l := client.LRange("tempQueue", 0, 10)
	fmt.Println(l)
}
