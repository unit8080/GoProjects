package main

import (
	"fmt"
	"strconv"
	"sync"
)

/*
func selectDrop(done chan int) {
	const cap = 5

	ch := make(chan string, cap)

	go func() {
		for p := range ch {
			fmt.Println("employee: received paper:", p)
		}
		done <- 1
	}()

	const work = 20
	for w := 0; w < work; w++ {
		select {
		case ch <- "paper" + strconv.Itoa(w):
			fmt.Println("manager : send ack")
			time.Sleep(50 * time.Nanosecond)
		default:
			fmt.Println("manager : drop")
		}
	}

	close(ch)
}
*/

func waitForTask(done chan int, ch chan string, wg *sync.WaitGroup) {

	go func() {
		defer wg.Done()
		for p := range ch {
			fmt.Println("employee : working", p)
		}

		done <- 1
	}()

	fmt.Println("Manager : Assigning all work")
	const work = 10
	for w := 0; w < work; w++ {
		ch <- "paper: " + strconv.Itoa(w)
	}
	fmt.Println("Manager : Assigned all work")

}

var done = make(chan int)

func main() {
	ch := make(chan string, 1)
	var wg sync.WaitGroup

	wg.Add(1)
	// selectDrop(done)
	waitForTask(done, ch, &wg)

	// signal range to stop looping
	close(ch)

	<-done
	wg.Wait()
	close(done)
}
