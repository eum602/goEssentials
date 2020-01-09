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
	c <- 42
}

//receive
func bar(c <-chan int) { //going from general to specific by converting bidirectional channel to a receive channel.
	fmt.Println(<-c) //pulling out the value
}
