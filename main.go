package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancelF := context.WithTimeout(ctx, time.Second)
	defer cancelF()
	time.Sleep(500 * time.Millisecond)
	select {
	case <- ctx.Done():
		fmt.Println("not finished")
	default:
		fmt.Println("finished")
	}
}