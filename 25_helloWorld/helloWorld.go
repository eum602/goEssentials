package main

import "fmt"

func main() {
	fmt.Println("Hello world")
	foo()
	fmt.Println("After foo")

	for i := 0; i < 100; i++ {
		if i%50 == 0 {
			fmt.Println(i)
		}
	}

	bar()
}

func foo() {
	fmt.Println("I am in foo")
}

func bar() {
	fmt.Println("Finishing with bar; all this is the control flow in go, I mean how statements and declarations are analyzed and executed")
}
