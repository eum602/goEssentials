package main

import (
	"fmt"
)

func main() {
	sliceDefinition()
	createAnSlice()
	loopOverSlice()
	getLenOfSlice()
	s := []int{3, 5, 6}
	getContentAtIndexOfSlice(s, 2)
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

func loopOverSlice() {
	fmt.Println("looping over a slice:")
	x := []int{1, 2, 3, 4}
	for i, v := range x {
		fmt.Println(i, v)
	}
}

func getLenOfSlice() {
	x := []int{1, 2, 3, 4}
	fmt.Println("The length of a slice is 'len(x)'", len(x))
}

func getContentAtIndexOfSlice(s []int, i int) {
	fmt.Println(`Index obtained with s[i] =>`, s[i])
}
