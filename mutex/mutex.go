package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var mut sync.Mutex

func main() {

	for i := 0; i < 10000; i++ {
		go add()
		wg.Add(1)
	}
	wg.Wait()
	fmt.Printf("race condition => %d\n", counter)

	counter = 0
	for i := 0; i < 10000; i++ {
		go add1()
		wg.Add(1)
	}
	wg.Wait()

	fmt.Printf("using mutex => %d\n", counter)
}

var counter = 0

func add() {
	counter++
	wg.Done()
}
func add1() {
	mut.Lock()
	counter++
	mut.Unlock()
	wg.Done()
}
