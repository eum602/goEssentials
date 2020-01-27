package main

import (
	"fmt"
)

func main() {
	definitions()
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
