package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go publish(ch)
	reader(ch)
}

func reader(ch chan int) {
	for x := range ch {
		fmt.Printf("Received %d\n", x)
		time.Sleep(1 * time.Second)
	}
}

func publish(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
		print("publish\n")
		time.Sleep(1 * time.Second)
	}
	close(ch)
}
