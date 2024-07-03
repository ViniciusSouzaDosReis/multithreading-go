package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type Message struct {
	id  int64
	msg string
}

func main() {
	ch1 := make(chan Message)
	ch2 := make(chan Message)

	var i int64 = 0

	go func() {
		for {
			atomic.AddInt64(&i, 1)
			msg := Message{
				id:  i,
				msg: "Hello from first channel",
			}
			ch1 <- msg
			time.Sleep(2 * time.Second)
		}
	}()

	go func() {
		for {
			atomic.AddInt64(&i, 1)
			msg := Message{
				id:  i,
				msg: "Hello from second channel",
			}
			ch2 <- msg
			time.Sleep(1 * time.Second)
		}
	}()

	for {
		select {
		case msg := <-ch1:
			fmt.Printf("Received from first channel: %s - ID: %d\n", msg.msg, msg.id)

		case msg := <-ch2:
			fmt.Printf("Received from second channel: %s - ID: %d\n", msg.msg, msg.id)

		case <-time.After(3 * time.Second):
			fmt.Println("Time out")
		}
	}
}
