package main

import (
	"context"
	"fmt"
	"time"
)

func someFunc(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Context canceled")
			return
		default:
			fmt.Println("Working ...")
			time.Sleep(1 * time.Second)
		}
	}
}
func main() {

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	go someFunc(ctx)

	time.Sleep(5 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)

}
