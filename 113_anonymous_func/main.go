package main

import (
	"fmt"
)

func main() {
	func() {
		fmt.Println("This is an anonymous function with no arguments")
	}()

	func(x int) {
		fmt.Println("This is an anomymous function that takes an argument int of value", x)
	}(42)
}
