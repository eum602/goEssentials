package main

import (
	"fmt"
)

func main() {
	sliceDefinition()
	createAnSlice()
	loopOverSlice()
	getLenOfSlice()
	s := []int{3, 5, 6, 45, 256, 76, 4}
	getContentAtIndexOfSlice(s, 2)
	sliceASlice(s, 3, 5)
	appendToSliceExamples(s)
	deleteFromSliceExample()

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

func sliceASlice(s []int, from, upToButNotIncluding int) {
	fmt.Println("Slicing a Slice", s)
	fmt.Println("General systanxis: s[from:upToButNotIncluding]")
	fmt.Println("s[", from, ":", upToButNotIncluding, "]", s[from:upToButNotIncluding])
	fmt.Println("s[:]  => ", s[:])
	fmt.Println("s[2:] => ", s[2:])
	fmt.Println("s[:4] => ", s[:4])
}

func appendToSliceExamples(s []int) {
	values := []int{0, 4, 79}
	fmt.Println(`using '...' to decompose the input (in case it is another slice ==> s1 := append(s, values...) ==>`)
	s1 := append(s, values...)
	fmt.Println("input => s=", s, "elements to append", values, `
	s + appended elements => s1 := append(s, values...) =  `, s1)
	s2 := append(s, 1, 4, 5)
	fmt.Println(`A simple way to append is by using something like this: 
	s = append(s,1,4,5)= `, s2)
}

func deleteFromSliceExample() {
	fmt.Println("\nDeleting some elements from an slice")
	s := []int{41, 25, 33, 455, 555, 256, 987, 45}
	fmt.Println(`Given the following slice:
	* `, s, `
	* Lets delete the elements 555 and 256, then we can do something like that: s = append(s[:4],s[6:]...) =>`)
	s = append(s[:4], s[6:]...)
	fmt.Println(s)

}
