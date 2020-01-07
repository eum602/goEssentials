package main

import (
	"fmt"
)

func main() {
	definition()

}

func definition() {
	fmt.Println(`
	* Channels are a higher order level way for synchronizing code and writing concurrent code.
	* Channels are a little place where we can send data
	`)

	c := make(chan int) //channel onto which i can put ints
	c <- 42             //onto c put number 42; in this line of code the channel BLOCKs, because the send and receive
	//must happen at the same time. In this case we are entering '42' but there is nothing to pull it off; then blocking occurs.
	// This is one of the most important concepts.
	fmt.Println(<-c) //take out of 'c' the content

}
