package main

import (
	"context"
	"fmt"
	"time"
)

func someFunc(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("context timedout")
	case <-time.After(5 * time.Second):
		fmt.Println("context didn't timeout")
	}
}
func main() {

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)

	go someFunc(ctx)

	time.Sleep(4 * time.Second)
	cancel()

	time.Sleep(2 * time.Second)

}
