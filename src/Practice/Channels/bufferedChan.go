package main

import "fmt"

func main() {

	ch := make(chan int, 2)

	go func() {
		ch <- 100
		ch <- 200
	}()

	fmt.Println("first read:", <-ch)
	fmt.Println("second read:", <-ch)
}
