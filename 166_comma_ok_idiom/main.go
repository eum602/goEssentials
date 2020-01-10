package main

import (
	"fmt"
)

func main() {
	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)

	go send(eve, odd, quit)
	receive(eve, odd, quit)

	fmt.Println("about to exit")
}

func send(eve, odd chan<- int, quit chan<- bool) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			eve <- i
		} else {
			odd <- i
		}
	}
	close(quit)
}

func receive(eve, odd <-chan int, quit <-chan bool) {
	for {
		select {
		case v := <-eve:
			fmt.Println("The value received from the even channel is: ", v)
		case v := <-odd:
			fmt.Println("The value received from the odd channel is: ", v)
		case i, ok := <-quit: //when the channel quit gets closed then something will be triggered here.
			//you can use 'i, ok' to check when a channel is closed
			if !ok {
				fmt.Println("From comma ok: ", i, ok) //will print the zero value
				return
				//if ok is false then all is ok and we exit the infinite for, we will continue
				//iterating over until getting the right value
			} else {
				fmt.Println("From comma ok: ", i)
			}
		}
	}
}
