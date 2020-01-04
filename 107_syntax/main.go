package main

import (
	"fmt"
)

func main() {
	foo() //no ARGUMENTS are passed in this example
	//Everything in GO is PASS BY VALUE
	bar("Erick") //Passing the ARGUMENT "Erick"
	s := woo("Erick")
	fmt.Println(s)
	x, y := greeting("Erick", "Pacheco")
	fmt.Println(x)
	fmt.Println(y)
	sum := sumInts(2, 3, 4, 5, 6, 7) //passing variadic values
	fmt.Println(sum)
	fmt.Println("==========================Spread operator==========================")
	xi := []int{1, 2, 3, 4, 5, 6, 7}
	sum = sumInts(xi...)
	fmt.Println("Print sum after spread slice was passed", sum)
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
func woo(s string) (returns string) {
	return fmt.Sprint("Hello from woo ", s) //Sprint => String print => this is gonna print a string
}
func greeting(fn string, ln string) (string, bool) {
	a := fmt.Sprint(fn, " ", ln, ` says "hello!"`)
	b := false
	return a, b
}

func sumInts(x ...int) int { //declaring VARIADIC PARAMETERS, in this way it is possible to declare many paramenters
	//then when received, those parameters will be stored as a SLICE.
	//variadic functions like this can receive ZERO or more arguments.
	println("===================SUMS===================")
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("For index ", i, " the value ", v, " is going to be added to sum ==> ", sum)
	}
	println("============>RESULTS<==============")
	fmt.Println("The total sum is", sum)
	return sum
}
