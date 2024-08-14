package main

import "fmt"

func main() {

	ch := make(chan int, 4)
	go func() {
		ch <- 1
		ch <- 2
		ch <- 3
		ch <- 4
		ch <- 5
		close(ch)
	}()
	for data := range ch {
		fmt.Println("data", data)
	}
}

// So the correct answer is:
// The program will output all values that were sent
// into the buffered channel.
// Why?
// Because we have 2 goroutines one to produce data
// (anonymous) and another to consume (main). This way
// we ensure a flawless execution that empties the
// channel on time.
