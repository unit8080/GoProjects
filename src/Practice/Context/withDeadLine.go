package main

import (
	"context"
	"fmt"
	"time"
)

func someFunc(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("Dealine exceeded")
	case <-time.After(3 * time.Second):
		fmt.Println("Deadline did not exceed")
	}
}
func main() {

	ctx := context.Background()
	ctx, cancel :=
		context.WithDeadline(ctx, time.Now().Add(2*time.Second))

	go someFunc(ctx)

	time.Sleep(6 * time.Second)

	cancel()

	time.Sleep(2 * time.Second)
}
