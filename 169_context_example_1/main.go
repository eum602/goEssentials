package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {
	example1()
}

func example1() {
	ctx, cancel := context.WithCancel(context.Background())
	fmt.Println("Checking error first time", ctx.Err())
	fmt.Println("Num of goroutines", runtime.NumGoroutine())

	go func() {
		n := 0
		for {
			select {
			case <-ctx.Done(): //Done is provided for use in "select" statements.
				//Done returns when a context is "canceled"
				return
			default:
				n++
				time.Sleep(time.Millisecond * 200)
				fmt.Println("working on smthng", n)
			}
		}
	}()

	fmt.Println("Simulating a process runs two seconds")
	time.Sleep(time.Second * 2)
	fmt.Println("error checking number second time", ctx.Err())
	fmt.Println("Num of goroutines", runtime.NumGoroutine())

	fmt.Println("Now stopping by cancelling the context ==> cancel()")
	cancel()
	fmt.Println("Context cancelled!")

	fmt.Println("Sleeping 2 more seconds in the main thread so giving time to the goroutine to finish")
	time.Sleep(time.Second * 2)
	fmt.Println("error checking number second time", ctx.Err())
	fmt.Println("Num of goroutines", runtime.NumGoroutine())
}
