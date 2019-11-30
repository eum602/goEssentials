package main
import "fmt"
import "runtime"

var a int
var b float32

func main(){
	//find all about numerics here: https://golang.org/ref/spec#Numeric_types
	//a=42.456 //this is not going to run because a is previously declared as int
	a=42
	b=42.56

	fmt.Println("a variable is",a, "its type is")
	fmt.Printf("%T\n",a)
	fmt.Println("b variable is",b,"its type is")	
	fmt.Printf("%T\n",b)

	var c int8 //-128 ==> 127
	//c=-129 //this is not going to work
	fmt.Println("c variable is",c,"its type is")	
	fmt.Printf("%T\n",c)

	//ALIASES:
	/*
	* byte: uint8
	* rune: int32
	*/

	fmt.Println("Using runtime package to get the OS",runtime.GOOS)
	fmt.Println("Using runtime package to print the running program architecture", runtime.GOARCH)
}