package main

import (
	"fmt"
)

func main() {
	arrayDefinition()
	createAnArray()
	modifyElementInArray()
}

func arrayDefinition() {
	fmt.Println(`
	* Array is a built in block in go
	* it is a zero based index
	* Elements in an array are of the same type
	* length is part of its type ==> [1]int != [2]int
	* passing array to a function will receive a COPY of ot NOT a pointer`)
}

func createAnArray() {
	var x [5]int //size five
	fmt.Println("Creating an array ==> var x [5]int\n", x)
}

func modifyElementInArray() {
	var x [5]int
	x[3] = 42
	fmt.Println("Modifying an element in an array ==> x[3] = \n", x)
}
