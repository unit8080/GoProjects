package main

import (
	"fmt"
	"sync"
)

func main() {

	ch := make(chan bool, 1)
	defer close(ch)

	var wg sync.WaitGroup
	counter := 0
	for i := 0; i < 100; i++ {
		wg.Add(1)

		go func() {
			ch <- true
			counter++
			<-ch
			fmt.Println(counter)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Counter:", counter)

}
