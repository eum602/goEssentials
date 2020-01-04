package main

import (
	"fmt"
)

func main() {
	defer sumInts("This text shoud be shown at the end of this final go program because a 'defer' keyword is deferring its execution")
	foo() //no ARGUMENTS are passed in this example
	//Everything in GO is PASS BY VALUE
	bar("Erick") //Passing the ARGUMENT "Erick"
	s := woo("Erick")
	fmt.Println(s)
	x, y := greeting("Erick", "Pacheco")
	fmt.Println(x)
	fmt.Println(y)
	sum := sumInts("some", 2, 3, 4, 5, 6, 7) //passing variadic values
	fmt.Println(sum)
	fmt.Println("==========================Spread operator==========================")
	xi := []int{1, 2, 3, 4, 5, 6, 7}
	sum = sumInts("some text", xi...)
	fmt.Println("Print sum after spread slice was passed", sum)
}

func foo() {
	fmt.Println(`Generalization of function:
	* func receiver(r receiver) identifier(parameters) (returns(s)){...}
	* Note: Recall that when you define a function we say PARAMETERS(if any), but when you CALL A FUNCTION you pass ARGUMENTS
	* When passing VARIADIC parameters to a function, you MUST ensure it has to be the last parameter to be defined =>
		func some(a int, s ...string).
	* If no variadic argument is passed the slice will still be created BUT as "nil" and the UNDERLYING array would have not been
	created, so the slice exits but the pointer is pointing to nothing
	* Defer keyword: https://golang.org/ref/spec#Keywords => invokes a function whose execution is deferred to the moment the
	 surrounding function returns:
			 either because the surrounding function executed a return statement, 
			 reached the end of its function body, 
			 or because the corresponding goroutine is panicking.
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

func sumInts(s string, x ...int) int { //declaring VARIADIC PARAMETERS, in this way it is possible to declare many paramenters
	//then when received, those parameters will be stored as a SLICE.
	//variadic functions like this can receive ZERO or more arguments.
	println("\n\n", s)
	println("===================SUMS===================")
	fmt.Println("content of variadic slice is", x, "and the type of that variadic is:")
	fmt.Printf("%T\n", x)
	fmt.Println(`The length of variadic slice is len(x) `, len(x))
	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("For index ", i, " the value ", v, " is going to be added to sum ==> ", sum)
	}
	println("============>RESULTS<==============")
	fmt.Println("The total sum is", sum)
	return sum
}
