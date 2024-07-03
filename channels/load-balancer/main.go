package main

import (
	"fmt"
	"time"
)

func worker(workerId int, ch <-chan int) {
	for x := range ch {
		fmt.Printf("The worker %d received %d\n", workerId, x)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	channel := make(chan int)
	qntWorker := 50

	for i := 0; i < qntWorker; i++ {
		go worker(i, channel)
	}

	for i := 0; i < 100; i++ {
		channel <- i
	}
}
