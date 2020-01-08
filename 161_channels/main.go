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
	* It is possible to create buffer channels by specifying the size of a buffer with the use of 'make':
		c := make(chan int, 2)
	* Channels are a type in go.
	`)

	c := make(chan int, 2) //setting to two the number of elements that can fit into this buffer

	c <- 42 //now it is not necessay to add it to a goroutine.
	c <- 46

	fmt.Println(<-c) //take out of 'c' the content => '42' ==> the previous  anonymous goroutine waas launched up,
	//now both the definition goroutine and the anonymous one are RUNNING AT THE SAME TIME. Now we are ready
	//to pull the value stored on 'c' off the goroutine.
	fmt.Println(<-c) //pulls off the second value of the channel: '46'

}
