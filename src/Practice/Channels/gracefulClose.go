package main

import "fmt"

func main() {

	ch := make(chan int, 1)

	go func() {
		ch <- 200
		close(ch)
	}()

	for {
		v, ok := <-ch
		if !ok {
			break
		}
		fmt.Println("recv:", v)
	}
}
