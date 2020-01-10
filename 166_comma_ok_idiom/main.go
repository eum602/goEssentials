package main

import (
	"fmt"
)

func main() {
	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	go send(eve, odd, quit)
	receive(eve, odd, quit)
	commaOkExample2()

	fmt.Println("about to exit")
}

func send(eve, odd, quit chan<- int) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			eve <- i
		} else {
			odd <- i
		}
	}
	close(quit)
}

func receive(eve, odd, quit <-chan int) {
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
			}
			fmt.Println("From comma ok: ", i)
		}
	}
}

func commaOkExample2() {
	fmt.Println("Second example of comma OK")
	c := make(chan int)
	go func() {
		c <- 42
		close(c) //it can or not be , in both cases this runs without problems
	}()
	v, ok := <-c
	fmt.Println(v, ok) //prints v=42 and ok=true
	fmt.Println(c, <-c)

	v, ok = <-c
	fmt.Println(v, ok) //prints v=0 and thus ok=false because there is nothing to pull from here and
	//the channel is closed
}
