package main

import (
	"context"
	"fmt"
)

func main() {

	ctx := context.Background()
	fmt.Println(ctx)

	ctx = context.TODO()
	fmt.Println(ctx)
}
