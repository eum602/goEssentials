package main

import (
	"fmt"
)

func main() {
	definitions()
	fmt.Println("1 + 4 = ", sum(1, 4))
}

func definitions() {
	fmt.Println(`
	Test files must have *_test.go format
	It is reccomended that test files be located in the same location where package files are
	All tests names must have Test* format name!!
	To run test use : 'go test' or 'go test -v'
	`)

}

func sum(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}
