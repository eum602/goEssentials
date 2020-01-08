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

	go func() { //by using a groutine we are launching off the tasks that are inside this function,
		//the code blocks in THIS GOROUTINE; the definition (func definition(){...}) goroutine fires up this routine,
		//and then the flow continues.
		c <- 42 //onto c put number 42;
	}()

	fmt.Println(<-c) //take out of 'c' the content ==> the previous  anonymous goroutine waas launched up,
	//now both the definition goroutine and the anonymous one are RUNNING AT THE SAME TIME. Now we are ready
	//to pull the value stored on 'c' off the goroutine.

}
