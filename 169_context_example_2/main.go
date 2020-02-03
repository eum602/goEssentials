package main

import (
	"context"
	"fmt"
)

func main() {
	example1()
}

func example1() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for v := range generate(ctx) {
		fmt.Println(v)
		if v == 5 {
			break //this line finishes the program triggering the cancel() at the end in the main goroutine
			//because of defer keyword
		}
	}
}

func generate(ctx context.Context) <-chan int {
	dst := make(chan int)
	n := 1
	go func() {
		for {
			select {
			case <-ctx.Done(): //when cancel is triggered then ctx.Done() matches here, helping to finish the
				//goroutine
				return
			case dst <- n:
				n++
			}
		}
	}()
	return dst
}
