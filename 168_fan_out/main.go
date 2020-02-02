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
	var wg sync.WaitGroup
	for v := range chan1 {
		wg.Add(1) //for each time I receive a value from chan1 then I add a new wait group
		
		//For each received value LAUNCHING A NEW GOROUTINE
		go func(v2 int) {
			chan2 <- delaySomeTime(v2) //doing something with a value that is been received by another channel(chan1)
			//so Faning in(Receiving).
			wg.Done() //after finishing this work I inform I have finished in this goroutine
		}(v)
	}
	wg.Wait() //blocks the launched goroutines dont close when "spread" channel closes. wg.wait will
	//wait for wg.Done().
	close(chan2)
}

func delaySomeTime(n int) (int){
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(500)))
	return n + rand.Intn(1000)
}
