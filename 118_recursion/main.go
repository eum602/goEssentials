package main

import (
	"fmt"
)

func main() {
	definition()
	example1()
	example2()
}

func definition() {
	fmt.Println(`
	Recursion: Recursion is when a function calls itself
	`)
}

func example1() {
	n := factorial(3)
	fmt.Println(n)
}

func example2() {
	n := loopFact(3)
	fmt.Println(n)

}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func loopFact(n int) int {
	total := 1
	for ; n > 0; n-- {
		total *= n
	}
	return total
}
