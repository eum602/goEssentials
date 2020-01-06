package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup //created a variable of type waitGroup

func main() {
	getEnv()  //gooroutines=1
	wg.Add(1) //There is one thing I am waiting for, so I add a thing to wait
	go foo()  //the code when added keyword "go" doesn't get printed nothing to the console because it will
	//be run on another routine thanks to keyword 'go' , also as this function is into main, then when main finishes
	//then all into it is forced to finish even when it has not finished, so foo is forced to terminate in case it is
	//still running ==> so to avoid that behavior a some sort of synchronization is needed.
	//WaitGroups is one option for the synchronization:
	getEnv()  //goroutines=2
	wg.Wait() //It will prevent the main function to finish until receiving the done signal

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

	wg.Done() //As the function has finished then done indicates to finish and remove things added
}
