package main //to run all tests under this folder use go test ./... and to see the documentation use godoc -http=:8080
import (
	"fmt"
	"github.com/eum602/goEssentials/195_examples/operations" //this package can be put not neccesarily here but on
	//the parent folder also or other parent
)

func main() {
	fmt.Println(operations.Sum(1, 2, 3))
}
