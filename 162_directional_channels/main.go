package main

import (
	"fmt"
)

func main() {
	definitions()
	sendChannel()
	receiveChannel()
	bidirectionalChannel()
	specificToGeneral()
	generalToSpecific()
	conversionGeneralToSpecificInChannels()
	conversionSpecificToGeneralInChannels()
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

func specificToGeneral() {
	fmt.Println(`============= SPECIFIC TO GENERAL ASSIGN ===================
	1. c := make(chan int, 1)
	2. cr := make(<-chan int, 1)
	3. c = cr ==> It WONT WORK because specfic to general does not assign
	`)
}

func generalToSpecific() {
	fmt.Println(`============= GENERAL TO SPECIFIC ASSIGN ===================
	1. c := make(chan int, 1)
	2. cr := make(<-chan int, 1)
	3. cs := make(chan<- int,1)
	4. cr = c ====> Assign General to specific WORKS
	5. cs = c  ====> Assign General to specific WORKS
	`)
}

func conversionGeneralToSpecificInChannels() {
	c := make(chan int)
	//cr := make(<-chan int)
	//cs := make(chan<- int)
	fmt.Println(` ==== GENERAL TO SPECIFIC CONVERTS ====
	1. c := make(chan int)
	2. cr := make(<-chan int)
	3. cs := make(chan<- int)
	THEN ==> (<-chan)(c) is possible.
	THEN ==> (chan<-)(c) is possible
	`)
	fmt.Printf("General to specific converts WORKS ==> (<-chan int)(c): \t%T\n", (<-chan int)(c))
	fmt.Printf("General to specific converts WORKS ==> (chan<- int)(c): \t%T\n", (chan<- int)(c))
}

func conversionSpecificToGeneralInChannels() {
	fmt.Println(` ==== SPECIFIC TO GENERAL CONVERTS ====
	1. cr := make(<-chan int)
	2. cs := make(chan<- int)
	THEN ==> (chan int)(cr) is NOT possible.
	THEN ==> (chan int)(cs) is NOT possible
	`)
}
