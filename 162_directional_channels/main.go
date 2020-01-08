package main

import (
	"fmt"
)

func main() {
	definitions()
	sendChannel()
	receiveChannel()
	bidirectionalChannel()
}

func definitions() {
	fmt.Println(`
	* Up to this point we have seen BIDIRECTIONAL CHANNELS, where we can push and pull values from a channel
	* But is possible to specify channel in only one direction: where it is possible to only send or only
		receive from a channel.
	`)
}

func sendChannel() {
	fmt.Println("\n=========SEND CHANNEL========")
	c := make(chan<- int, 2) // '<-' a channel that only send values on to the channel. We are
	//only able to send TO the channel
	c <- 34
	c <- 789
	fmt.Println(`Sending value to a channel buffer:
		 'c <- 34'
		 'c <- 789'
	`)
	fmt.Printf("The type of the channel is: %T\n", c)
}

func receiveChannel() {
	fmt.Println("\n=========RECEIVE CHANNEL========")
	c := make(<-chan int, 2) //We can only receive FROM the channel
	//c <- 5 //it is not allowed because the channel is RECEIVE ONLY thus we can only RECEIVE FROM the channel.
	fmt.Println("Receiving value from a channel, means we can only receive FROM a channel")
	fmt.Printf("The type of the channel is: %T\n", c)
}

func bidirectionalChannel() {
	fmt.Println("\n=========BIDIRECTIONAL CHANNEL========")
	c := make(chan int, 2) //We can only receive FROM the channel
	c <- 5
	fmt.Println("Sending value to a channel 'c <- 5' ")
	fmt.Println("Receiving value from a channel:", <-c)
	fmt.Printf("The type of the channel is: %T\n", c)
}
