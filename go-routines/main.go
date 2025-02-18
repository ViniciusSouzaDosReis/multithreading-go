package main

import (
	"fmt"
	"time"
)

func task(name string) {
	for i := 0; i <= 10; i++ {
		fmt.Printf("%d: The task %s is running\n", i, name)
		time.Sleep(time.Second)
	}
}

func main() {
	go task("A")
	go task("B")
	time.Sleep(15 * time.Second)
}
