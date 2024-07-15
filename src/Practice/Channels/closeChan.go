package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int, 2)

	go func() {
		fmt.Println("send 100")
		ch <- 100
		time.Sleep(1 * time.Millisecond)
		fmt.Println("send 200")
		ch <- 200
		time.Sleep(1 * time.Millisecond)
		fmt.Println("send 300")
		ch <- 300
		close(ch)
	}()

	id := 1
	for v := range ch {
		fmt.Println("Recv", v)
		id++
	}
}
