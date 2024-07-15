package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("Worker %d: got job: %d\n", id, j)
		time.Sleep(time.Second)
		results <- j * j
	}
}

func manager(wg *sync.WaitGroup, numJobs int, results <-chan int) {

	fmt.Println("starting manager...")
	for {
		select {
		case v := <-results:
			fmt.Println(v)
			numJobs--
			if numJobs == 0 {
				wg.Done()
				return
			}
		}
	}

	// for a := 1; a <= numJobs; a++ {
	// 	fmt.Println(<-results)
	// }

	// for v := range results {
	// 	fmt.Println(v)
	// }
}

func fanOutFanIn(wg *sync.WaitGroup) {

	fmt.Println("starting finoutfanIn ...")
	numJobs := 10
	jobs := make(chan int, 5)
	results := make(chan int, 5)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	wg.Add(1)
	go func() {
		fmt.Println("calling  manager ...")
		manager(wg, numJobs, results)
	}()

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)
	wg.Done()
}

func main() {

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		fanOutFanIn(&wg)
	}()
	// go fanOutFanIn(&wg)
	wg.Wait()
}
