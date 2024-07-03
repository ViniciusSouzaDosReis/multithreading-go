package main

import "fmt"

func main() {
	ch := make(chan string)
	go input("hello", ch)
	output(ch)
}

func input(name string, ch chan<- string) {
	ch <- name
}

func output(ch <-chan string) {
	fmt.Print(<-ch)
}
