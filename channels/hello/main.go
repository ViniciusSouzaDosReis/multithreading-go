package main

// Thread 1
func main() {
	channel := make(chan string) // Channel is empty

	// Thread 2
	go func() {
		channel <- "Hello World!" // Chanell full
	}()

	// Thread 1
	msg := <-channel
	print(msg)
}
