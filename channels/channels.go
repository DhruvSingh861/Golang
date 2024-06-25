package main

import (
	"fmt"
	"sync"
)

func main() {

	ch := make(chan int)
	wg := &sync.WaitGroup{}

	wg.Add(3)

	go func(c chan<- int, w *sync.WaitGroup) {
		ch <- 1
		ch <- 2
		close(ch)
		wg.Done()
	}(ch, wg)

	go func(c <-chan int, w *sync.WaitGroup) {
		val, isClosed := <-ch
		fmt.Println(val, isClosed)
		wg.Done()
	}(ch, wg)

	go func(c <-chan int, w *sync.WaitGroup) {
		fmt.Println(<-ch)
		val, isClosed := <-ch
		fmt.Println(val, isClosed)
		wg.Done()
	}(ch, wg)

	wg.Wait()
}
