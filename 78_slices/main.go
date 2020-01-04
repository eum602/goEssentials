package main

import (
	"fmt"
)

func main() {
	sliceDefinition()
	createAnSlice()
}

func sliceDefinition() {
	fmt.Println(`Definitions:	
	* Slices are a composite data types ==> built based on primitive data types
	* https://golang.org/doc/effective_go.html#composite_literals => 
		an expression that creates a new instance each time it is evaluated.
	`)
}

func createAnSlice() {
	x := []int{1, 2, 3, 4}
	fmt.Println(`Composing values ==> eg. Creating slices of int ==> Groups VALUES OF THE SAME TYPE
	x := []int{1, 2, 3, 4} = `, x)
}
