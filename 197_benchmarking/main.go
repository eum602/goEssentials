package main

import (
	"fmt"

	"github.com/eum602/goEssentials/197_benchmarking/greet"
)

func main() {
	definitions()
	fmt.Print(greet.Greet("Erick"))
}

func definitions() {
	fmt.Println(`
	* Benchmarking: tests the performance of our code.
	`)
}
