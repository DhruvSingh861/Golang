package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/go-redis/redis"
)

const (
	redisAddr     = "localhost:6379"
	redisPassword = ""
	redisDB       = 0
	outputFile    = "output.txt"
)

var client = redis.NewClient(&redis.Options{
	Addr:     redisAddr,
	Password: redisPassword,
	DB:       redisDB,
})
var wg sync.WaitGroup

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	file, err := os.OpenFile(outputFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Failed to open file: %v\n", err)
		return
	}
	defer file.Close()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for {
			err := readFromRedisAndWriteToFile(ctx, file)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			}
			time.Sleep(1 * time.Second) // Adjust the interval as needed
		}
	}()
	wg.Wait()
}

func readFromRedisAndWriteToFile(ctx context.Context, file *os.File) error {
	val, err := client.Get("name").Result()
	if err == redis.Nil {
		return nil
	} else if err != nil {
		return fmt.Errorf("failed to read from Redis: %v", err)
	}
	_, err = file.WriteString(val + "\n")
	if err != nil {
		return fmt.Errorf("failed to write to file: %v", err)
	}
	fmt.Printf("Data written to file: %s\n", val)
	return nil
}
