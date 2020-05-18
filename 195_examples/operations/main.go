//Package operations gives many basic arithmetic operations.
//Find more testable examples here: https://blog.golang.org/examples
package operations

import (
	"fmt"
)

func main() {
	definitions()
	fmt.Println("1 + 4 = ", Sum(1, 4))
}

func definitions() {
	fmt.Println(`
	Test files must have *_test.go format
	It is reccomended that test files be located in the same location where package files are
	All tests names must have Test* format name!!
	To run test use : 'go test' or 'go test -v'
	To get information about this package on terminal use: go doc github.com/eum602/goEssentials/195_examples/operations
	To check for details and observations use golint ./...
	Other similar to golint is go vet ./...
	`)

}

//Sum : A custom sum function. Only for learning purpose in order to show how to wirte documentation
func Sum(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}
