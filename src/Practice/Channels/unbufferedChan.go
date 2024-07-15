package main

import "fmt"

func main() {

	ch := make(chan int)

	go func() {
		ch <- 300
	}()

	fmt.Println("read from channel:", <-ch)
}
