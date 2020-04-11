package main //to run all tests under this folder use go test ./... and to see the documentation use godoc -http=:8080
//to see on godoc use ==> godoc -http=:8080 ==> packages ==> then find for Sum (the package into this folder)
import (
	"fmt"
	"github.com/eum602/goEssentials/195_examples/operations" //this package can be put not neccesarily here but on
	//the parent folder also or other parent
)

func main() {
	fmt.Println(operations.Sum(1, 2, 3))
}
