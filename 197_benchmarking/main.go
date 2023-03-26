package main

import (
	"fmt"

	"github.com/eum602/goEssentials/197_benchmarking/greet"
)

func main() {
	definitions()
	fmt.Print(greet.Greet("Erick"))
}

func definitions() {
	fmt.Println(`
	* Benchmarking: tests the performance of our code.
	* To run test use : 'go test' or 'go test -v'
	* To get help run: 'go test --help'
	* To run all test inside a folder: go test ./... -v
	* to run tests made on a 'Greet package'
	* go test -bench . -v
	* go test -bench . -v -count 5 ===> '5' consecutive benchamark test just to see if results on each are consistent
	* To show how much of our code is addressed in percentage by our test code: go test -cover
	* go test -cover greet
	* go test -cover .
	* To print result in a file: go test -coverprofile greet c.out
	* To print results in an html file: go tool cover -html c.out
	`)
}
