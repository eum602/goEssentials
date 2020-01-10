package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := fanIn(boring("eum602"), boring("alice"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("You are boring")
}

func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

func fanIn(i1, i2 <-chan string) <-chan string { //this func only returns a receiver channel
	c := make(chan string)
	go func() {
		for {
			c <- <-i1 //<-i1 means extracting the value of  channel i1; c<- passing the value to channel 'c'
		}
	}()

	go func() {
		for {
			c <- <-i2
		}
	}()

	return c
}
