package main

import (
	"fmt"
	"os"
)

func main() {
	panicExample()
}

func panicExample() {
	fmt.Println("*************Showing a Panic example*************")
	fmt.Println("With panic all deferred functions are run!!")
	f2, err := os.Open("non-existent-file.txt")
	if err != nil {
		panic(err)
	}
	defer f2.Close()
	defer executedDeferred()
}

func executedDeferred() {
	fmt.Println("This function will be excuted because it is called with panic example")

}
