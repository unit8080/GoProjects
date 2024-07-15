package main

import (
	"context"
	"fmt"
	"time"
)

func someFunc(ctx context.Context) {

	select {
	case <-ctx.Done():
		fmt.Println("context canceled")
	case <-time.After(5 * time.Second):
		fmt.Println("ctx not cancelled, timer expired")
	}
}
func main() {

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	go someFunc(ctx)

	time.Sleep(4 * time.Second)
	cancel()

	time.Sleep(2 * time.Second)
}
