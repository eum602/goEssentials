package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	example1()
}

func example1() {
	chan1 := make(chan int)
	chan2 := make(chan int)

	go spread(chan1)
	go fanOutIn(chan1, chan2)
	for v := range chan2 {
		fmt.Println(v)
	}
}

func spread(c chan<- int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}

func fanOutIn(chan1 <-chan int, chan2 chan<- int) {
	const ngoroutines = 10
	var wg sync.WaitGroup
	wg.Add(ngoroutines) //adding generally only 10 goroutines

	for i := 0; i < ngoroutines; i++ { //launching 10 goroutines
		go func() {
			for v := range chan1 { //structure to wait for new values from channel chan1
				chan2 <- delaySomeTime(v)
			}
			wg.Done() //after finishing this work I inform I have finished in this goroutine
		}()
	}

	wg.Wait()
	close(chan2)
}

func delaySomeTime(n int) int {
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(500)))
	return n + rand.Intn(1000)
}
