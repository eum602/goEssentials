package main

import (
	"fmt"
	"sync"
)

func main() {
	definitions()
	eve := make(chan int)
	odd := make(chan int)
	fanin := make(chan int)

	go send(eve, odd)
	go receive(eve, odd, fanin)

	for v := range fanin {
		fmt.Println(v) //receiving and printing all the values available in fanin UNTIL it is closed.
	}

}

func definitions() {
	fmt.Println(`
	* Fan In/Fan out are  concurrent design patterns.
	* It is like chunk many work out as many goroutines as possible, then when those gives results
		 we'll fan those results back into another channel as a result.
	`)
}

func send(eve, odd chan<- int) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			eve <- i
		} else {
			odd <- i
		}
	}
	close(eve)
	close(odd)
}

func receive(eve, odd <-chan int, fanin chan<- int) {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for v := range eve {
			fanin <- v
		}
		wg.Done()
	}()

	go func() {
		for v := range odd {
			fanin <- v
		}
		wg.Done()
	}()

	wg.Wait() //waiting for all channels 'odd' and 'eve' to pass all values to fanin channel until
	//both finishes
	close(fanin) //closing the fanin channel
}
