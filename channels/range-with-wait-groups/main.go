package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan int)

	wg := sync.WaitGroup{}
	wg.Add(10)

	go publish(ch)
	go reader(ch, &wg)

	wg.Wait()
}

func reader(ch chan int, wg *sync.WaitGroup) {
	for x := range ch {
		fmt.Printf("Received %d\n", x)
		time.Sleep(1 * time.Second)
		wg.Done()
	}
}

func publish(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
		println("publish")
		time.Sleep(1 * time.Second)
	}
	close(ch)
}
