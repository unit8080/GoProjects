package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup

	ch := make(chan int) // synchronizing who even or odd
	input := 100

	for i := 0; i < 2; i++ {
		wg.Add(1)

		go func(thread int) {
			select {
			case v := <-ch:
				var even bool
				if v == 0 {
					even = true
					ch <- 1
				} else {
					even = false
				}
				if even {
					for i := 0; i <= input; i++ {
						if i%2 == 0 {
							fmt.Printf("thread:%d, %d\n", thread, i)
							time.Sleep(time.Microsecond)
						}
					}
				} else {
					for i := 0; i <= input; i++ {
						if i%2 != 0 {
							// fmt.Println("odd thread", i)
							fmt.Printf("thread:%d, %d\n", thread, i)
							time.Sleep(time.Microsecond)
						}
					}
				}
				wg.Done()
			}
		}(i)
	}

	ch <- 0 //
	wg.Wait()
}
