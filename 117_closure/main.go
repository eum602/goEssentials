package main

import (
	"fmt"
)

func main() {
	definition()
	example1()
	example2()
}

func definition() {
	fmt.Println(`
	Closure: Closure is the concept of enclosing the scope of a variable to a limited area.
	`)
}

func example1() {
	var x int
	x = 0
	fmt.Println(x)

	{
		var y int = 0 //y is only available on the curly braces code bloque scope
		fmt.Println(y)
	}
}

func example2() {
	a := incrementor()
	b := incrementor()
	fmt.Println(a()) //1
	fmt.Println(a()) //2
	fmt.Println(b()) //1
	fmt.Println(b()) //2
}

func incrementor() func() int { //func foo is going to return a func that in turn returns an int.
	var x int //by default it is initialized with the zero value "0"
	//the scope of "x" is all the incrementor function
	return func() int {
		x++
		return x
	}
}
