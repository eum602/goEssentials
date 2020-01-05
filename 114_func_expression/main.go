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

	f := func(p int) { //By using this example it is shown that type functions can be used as any other type.
		//that is why we assign the tyoe function to the variable f
		fmt.Println("This is a function as an expression, as any function it can accept or not paramenters", p)
	}

	f(34)
}
