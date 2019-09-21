package main

import "fmt"

//go is a static programming language(NOT an dynamic one), because when you assign a value of some tye to a variable, then
//you CANNOT assign another type of variable to that variable
var y = 42
var z = "This is some text"
var k = `==========
Using back quotes is also possible in golang. into this we can use "" 
=> double quotes
note how it allows me to put in in several lines
============`

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y) //prints the type of variable

	fmt.Println(z)
	fmt.Printf("%T\n", z)

	//z = 34 //this is not possible because of static behavior of go
	fmt.Println("k variable", k)
	fmt.Printf("%T\n", k)

	////////primitive data types/////////
	//bool, int,

	//composite data types
	//slices, struct,
}
