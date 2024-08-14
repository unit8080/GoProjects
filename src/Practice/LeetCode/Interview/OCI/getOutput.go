// generate API to give output for last one minute
// Put API to implement to increment key-value, no return value
// return Get api output in array
// maintain time stamp

package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	duration = time.Second
)

type tickerService struct {
}

func NewTickerService() *tickerService {
	return &tickerService{}
}
func (ds *tickerService) Run(ctx context.Context) {
	ticker := time.NewTicker(duration * 5)
	for {
		select {
		case <-ticker.C:
			go func() {
				ds.task()
			}()
		case <-ctx.Done():
			fmt.Println("goroutine terminated!")
			ticker.Stop()
			return
		}
	}
}

func (ds *tickerService) task() {
	fmt.Println("do some task here")
}
func main() {
	fmt.Println("Implement ticker service")

	ctx, cancelFunc := context.WithCancel(context.Background())
	ds := NewTickerService()
	go ds.Run(ctx)

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)
	<-shutdown
	fmt.Println("Program shutdown!")
	cancelFunc()
}
