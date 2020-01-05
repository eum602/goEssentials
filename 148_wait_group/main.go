package main

import (
	"fmt"
	"runtime"
)

func main() {
	getEnv()

}

func getEnv() {
	fmt.Println("OS\t ", runtime.GOOS)
	fmt.Println("ARCH\t ", runtime.GOARCH)
	fmt.Println("CPUs\t ", runtime.NumCPU())
	fmt.Println("Goroutines\t ", runtime.NumGoroutine())
}
