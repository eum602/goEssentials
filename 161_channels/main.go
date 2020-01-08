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

	c := make(chan int, 1) //channel onto which i can put ints
	//the added number '1' indicates that it is a buffer channel which will allow certain values to sit in
	//that channel regardless or not something is ready to pull it off. For example '1' indicates that
	//the buffer channel will allow 'one value' to sit in there.

	c <- 42 //now it is not necessay to add it to a goroutine.

	fmt.Println(<-c) //take out of 'c' the content ==> the previous  anonymous goroutine waas launched up,
	//now both the definition goroutine and the anonymous one are RUNNING AT THE SAME TIME. Now we are ready
	//to pull the value stored on 'c' off the goroutine.

}
