package main

import "fmt"

var x = "hello" //var allows to declare a variable outside the function body
//the scope of this variable is at a package level
var z int //here we are declaring a INT value with the name z with a package scope level (var); z is initialized with
//the ZERO value
//false for booleans, 0 for numeric types, "" for strings, and nil for pointers, functions, interfaces, slices, channels, and maps. This initialization is done recursively, so for instance each element of an array of structs will have its fields zeroed if no value is specified.
//https://golang.org/ref/spec#The_zero_value
func main() {
	//https://golang.org/ref/spec
	some()
	fmt.Println("z varaible declared and initialized with 'var z int' which initializes with the zero value", z)
}

func some() {
	fmt.Println("calling var variable from an auxiliary function", x)
}
