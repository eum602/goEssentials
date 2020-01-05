package main

import (
	"fmt"
	"runtime"
)

func main() {
	getEnv() //gooroutines=1
	go foo() //the code when added keyword "go" doesn't get printed nothing to the console
	getEnv() //goroutines=2

}

func getEnv() {
	fmt.Println("======= Go runtime =======")
	fmt.Println("OS\t ", runtime.GOOS)
	fmt.Println("ARCH\t ", runtime.GOARCH)
	fmt.Println("CPUs\t ", runtime.NumCPU())
	fmt.Println("Goroutines\t ", runtime.NumGoroutine())
	if runtime.NumGoroutine() > 1 {
		fmt.Println("At this point there is concurrent code up and running")
	}
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("num is", i)
	}
}
