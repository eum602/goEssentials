package main

import (
	"fmt"
)

func main() {
	definitions()
	selectExample()
}

func definitions() {
	fmt.Println(`
	* Channels BLOCK
	* We can use RANGE LOOP to continue to listening over a channel until it gets closed.
	* There are some tools we can use to work wit channels
	`)
}

func selectExample() {
	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	//send
	go send(eve, odd, quit)

	//receive
	receive(eve, odd, quit)

	fmt.Println("about to exit")
}

func send(e, o, q chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}

	//close(e)
	//close(o)
	q <- 0
}

func receive(e, o, q <-chan int) {
	for { //infinite for
		select {
		case v := <-e:
			fmt.Println("From the eve channel", v)

		case v := <-o:
			fmt.Println("From the odd channel", v)

		case v := <-q:
			fmt.Println("From the quit channel", v)
			return //when it is quit then exits the infinite for loop
		}
	}
}
