package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Printf("GOROUTINES RUNNING UP %d\n", runtime.NumGoroutine())
	ctx := context.Background()
	ctx, cancelF := context.WithCancel(ctx)
	defer cancelF()

	for i := 0; i < 100; i++ {
		go func(n int) {
			fmt.Println("launching goroutine", n)
			for {
				select {
				case <-ctx.Done():
					runtime.Goexit()
					//return
				default:
					fmt.Printf("goroutine %d is doing work\n", n)
					time.Sleep(50 * time.Millisecond)
				}
			}
		}(i)
	}
	time.Sleep(time.Millisecond)
	fmt.Printf("GOROUTINES RUNNING AFTER ONE MILLISECOND!!! %d\n", runtime.NumGoroutine())
	cancelF()
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("GOROUTINES RUNNING AFTER CANCELFUNC CALLED!!! %d\n", runtime.NumGoroutine())

}
