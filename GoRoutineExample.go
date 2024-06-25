package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("started....")
	start := time.Now()

	// Normal Execution
	for i := 0; i <= 10; i++ {
		task1()
	}
	fmt.Println("done")
	fmt.Println(time.Since(start))
	start = time.Now()

	//using go routine
	fmt.Println("sync started....")
	wg.Add(10)
	for i := 1; i <= 10; i++ {
		go task2()
	}
	wg.Wait()
	fmt.Println("done")
	fmt.Println(time.Since(start))
}
func task1() {
	time.Sleep(1000 * time.Millisecond)
	// fmt.Printf("Task %d completed", i)
}

func task2() {
	time.Sleep(1000 * time.Millisecond)
	// fmt.Printf("Task %d completed", i)
	wg.Done()
}
