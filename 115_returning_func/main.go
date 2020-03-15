package main

import (
	"fmt"
)

func main() {
	definitions()

	//step 1: returning a simple type from a function
	example1()
}

func definitions() {
	fmt.Println(`
	Functions are first class citizens in go. It is possible to return a function from a function, it works in the same
	way as we return a string or something else.`)
}

func example1() {
	//example 1
	s1 := foo()
	fmt.Println(s1)
}

func foo() string {
	s := "hello world"
	return s

}
