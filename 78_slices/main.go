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
	createSliceWithMake()
}

func sliceDefinition() {
	fmt.Println(`Definitions:	
	* Slices are a composite data types ==> built based on primitive data types
	* https://golang.org/doc/effective_go.html#composite_literals => 
		an expression that creates a new instance each time it is evaluated.
	* slices are built on top of an array, they can dinamycally change in size. When a slice is created
		 it is built on top of an array
	* MAKE: if the array, on which a slice will be built, is known then we can use MAKE to create a new hiddeb array:
		https://golang.org/ref/spec#Slice_types
		 + It requires to indicate, the type of the undeliying array, length of the array, the capacity.
		 + The length indicate the length of the SLICE.
		 + The capacity of a slice is the number of elements for which there is space allocated IN THE UNDERLYING ARRAY
		 + From both previous definitions  0 <= len <= cap
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

func createSliceWithMake() {
	s := make([]int, 10, 11)
	fmt.Println("\n\nCreating new slice with make => s := make([]int, 10, 20)", s)
	fmt.Println("length of slice => len(s):", len(s), ". Capacity of slice => cap(s) ", cap(s))

	fmt.Println("Adding some elements to a slice:")
	s[1] = 200
	s[8] = 7
	fmt.Println(s)

	fmt.Println(`Trying to add an element to a defined slice will throw out of range error.
		For example if slice 's' has length 10 => s[12]=3 will throw an error.
	`)

	fmt.Println(`But we are able to append a new element to a slice: 
	s = append(s,2) ==>`, s)
	s = append(s, 2)
	fmt.Println("length of slice => len(s):", len(s), ". Capacity of slice => cap(s) ", cap(s))

	s = append(s, 45)
	fmt.Println(`If we increase the length of the defined slice up to the point that
	the 'length' reaches 'capacity' then the golang runtime will throw the undelying array and create a new one with the
	existing elements but much larger ,==> the capacity will be increased to the double `)
	fmt.Println("length of slice => len(s):", len(s), ". Capacity of slice => cap(s) ", cap(s), "!!!")

}
