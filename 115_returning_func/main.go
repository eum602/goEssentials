package main

import (
	"fmt"
)

func main() {
	definitions()

	//step 1: returning a simple type from a function
	example1()

	example2()

	example3()

	example4()
}

func definitions() {
	fmt.Println(`
	Functions are first class citizens in go. It is possible to return a function from a function, it works in the same
	way as we return a string or something else.
	Function is a TYPE.`)
}

func example1() {
	//example 1
	fmt.Println("Returning a simple type from a function")
	s1 := foo()
	fmt.Println(s1)
}

func example2() {
	fmt.Println("\nRetuning a function that returns a type 'int'")
	s1 := func() int {
		return 345
	}()

	fmt.Println(s1)
	fmt.Printf("%T", s1)
}

func example3() {
	fmt.Println("\nRetuning a function that returns a type 'func() int'")
	s1 := func() int {
		return 345
	}

	fmt.Println(s1)
	fmt.Printf("%T", s1)
}

func example4() {
	fmt.Println("\nUsed a declared function that returns a function() which in turn returns a type int")
	s1 := bar()
	fmt.Printf("%T", s1)

	fmt.Println("\nNow EXECUTING my variable which has the function that returns an int")
	a := s1()
	fmt.Println("The returned value is", a)
	a = bar()()
	fmt.Println("because a=s1() and s1=bar() then a=bar()() =", a)
}

func foo() string {
	s := "hello world"
	return s

}

func bar() func() int {
	return func() int {
		return 123
	}
}
