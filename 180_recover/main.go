package main

import (
	"fmt"
)

func main() {
	definitions()
	f()
	fmt.Println("Returned normally from f even when a panic occurred in g() that was called in f()")
}

func definitions() {
	fmt.Println(`
	Recover is a built-in function that regains control of a panicking goroutine. Recover is only useful
	inside deferred functions. During normal execution, a call to recover will return nil and have no other
	effect. If the current goroutine is panicking, a call to recover will capture the value given to panic
	and resume normal execution.
	`)
}

func f() {
	defer func() {
		if r := recover(); r != nil { //calling recover() and then analyzing if recover happenned.
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from g.")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}

/*
Calling g.
Printing in g 0
Printing in g 1
Printing in g 2
Printing in g 3
Panicking!
Defer in g 3
Defer in g 2
Defer in g 1
Defer in g 0
Recovered in f 4
Returned normally from f even when a panic occurred in g() that was called in f()
*/
