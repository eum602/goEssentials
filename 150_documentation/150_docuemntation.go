package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	//definitions()
	example1()
}

func definitions() {
	fmt.Println(`Concurrency requires some subtleties to implement.
	Shared values are passed to channels; only one goroutine has access to the value at
	any given time, so data RACE cannot occur by design
	==> DO NOT COMMUNICATE BY SHARING MEMORY, SHARE MEMORY BY COMMUNICATING
	
	GOROUTINE ==> A function executing concurrently with other goroutines in the same address space.
	A function or a method can be launched into a goroutine by using the "go" keyworkd, this launches
	the function as an INDEPENDENT CONCURRENT THREAD OF CONTROL, or goroutine in the same address space.
	When the goroutine terminates its goroutine also terminates.
	
	Function literals are closures`)
}

func example1() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine()) //GOMAXPROCS ==> number of maximum goroutines that can be executing simultaneously, by default
	//it is set to the max number of cores available
	counter := 0
	var wg sync.WaitGroup
	const gs = 100
	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() { //launching a hundred of goroutines to simulate race condition
			v := counter //reading a value
			//time.Sleep(time.Second) //sleep the goroutine
			runtime.Gosched() //contributes to race by allowing something else to run, this is the breaking point to
			//simulate the race condition
			v++
			counter = v //writing a value to the shared variable
			wg.Done()
		}()

		fmt.Println("Goroutines:", runtime.NumGoroutine())

	}

	wg.Wait()

	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("count", counter)

}

//run with go run -race *.go ==> for more info see go help build
//Found 1 data race(s)
