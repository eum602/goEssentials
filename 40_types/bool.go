package main
import "fmt"

var x bool //x is of type bool, with the zero value "false"

func main(){
	x = 5==2
	fmt.Println("hello World")
	fmt.Printf("%T\n",x)
	fmt.Println("value of x is",x)
}