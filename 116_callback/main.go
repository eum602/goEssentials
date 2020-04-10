package main

import (
	"fmt"
)

func main() {
	definitions()
	example1()
}

func definitions() {
	fmt.Println(`
	A callback is passing a func as an argument.
	* we can return functions, 
	* we can assging functions to variables and 
	* we can also pass a function to another function  as an argument.

	This is what is called functional programming 
	`)
}

func example1() {
	ii := []int{1, 2, 3, 4, 5, 6}
	s1 := sum(ii...) //unfurling
	fmt.Println(s1)

	s2 := even(sum, ii...)
	fmt.Println("sum of odd numbers ", s2)

}

func sum(xi ...int) int {
	fmt.Printf("%T\n", xi) //input is a slice of int
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

//this is a callback ==> taking in a function as an argument
func even(f func(xi ...int) int, vi ...int) int {
	/*
		Passing as parameters a function with a specific SIGNATURE
		also another parameter which is a variadic one
	*/
	var yi []int
	for _, v := range vi {
		if v%2 == 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}
