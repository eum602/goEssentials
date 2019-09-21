package main

import "fmt"

var a int

type myStruct int //we have created our own type myStruct and the type is an int
var b myStruct    //now I am declaring the variable b of type myStruct

func main() {
	a = 4
	fmt.Printf("%T\n", a)
	b = 42
	fmt.Printf("%T\n", b) //b is of type myStruct from package main

	//a=b => it is not possible because both are of different types

	//But what is possible is to make conversions(in other programming languages it is called casting)
	//see golang.org/ref/spec and search  in  the conversion section
	a = int(b)
}
