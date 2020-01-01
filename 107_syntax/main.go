package main

import (
	"fmt"
)

func main() {
	foo() //no ARGUMENTS are passed in this example
	//Everything in GO is PASS BY VALUE
	bar("Erick") //Passing the ARGUMENT "Erick"
}

func foo() {
	fmt.Println(`Generalization of function:
	func receiver(r receiver) identifier(parameters) (returns(s)){...}
	Note: Recall that when you define a function we say PARAMETERS(if any), but when you CALL A FUNCTION you pass ARGUMENTS
	`)
}

func bar(s string) {
	fmt.Println(`This is another function with PARAMETER 's'   :`, s)
}
