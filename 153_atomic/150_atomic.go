package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

//NOTE: run with go run -race *.go ==> for more info see go help build
func main() {
	//definitions()
	example1()
}

func example1() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine()) //GOMAXPROCS ==> number of maximum goroutines that can be executing simultaneously, by default
	//it is set to the max number of cores available

	var counter int64 //int64 ==> package atomic

	const gs = 100

	var wg sync.WaitGroup

	wg.Add(gs) //waiting fot a hundred goroutines to finish
	for i := 0; i < gs; i++ {
		go func() { //ATOMIC definitely avoids race conditions, it delivers the same functionality of
			//mutex ==> AVOID RACE CONDITIONS
			atomic.AddInt64(&counter, 1) //write to my counter, addint a number
			runtime.Gosched()
			fmt.Println("Counter\t", atomic.LoadInt64(&counter)) //reading my counter
			wg.Done()
		}()

		fmt.Println("Goroutines:", runtime.NumGoroutine())

	}

	wg.Wait()

	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("count", counter)

}

//run with go run -race *.go ==> for more info see go help build
