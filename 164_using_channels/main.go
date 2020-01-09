package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	//send
	go foo(c) //thanks to 'go' keyword the foo function is launched up from the main function

	//receive
	bar(c) //this function will run in the main loop and will pull a value from the goroutine
	//launched in foo function!! ==> this is the way how to get values from goroutines

	fmt.Println("about to exit")
}

//send
func foo(c chan<- int) { //going from general to specific by converting bidirectional channel to a send channel.
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c) //closing the channel when all this done, otherwise an error will be thrown because
	//bar function will be waiting for more values to come.
}

//receive
func bar(c <-chan int) { //going from general to specific by converting bidirectional channel to a receive channel.
	for v := range c { //the range just keeps looping over the channel until the channel gets closed
		fmt.Println(v) //pulling out the value
	}
}
