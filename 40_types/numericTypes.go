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

	//********************STRINGS*****************************//
	s:= "Hello this is nome non useful text"
	fmt.Println(s)
	fmt.Printf("%T\n",s)
	
	u:=`This is a text created with backticks`

	fmt.Println(u)


	bs:=[]byte(s) //using conversion to convert from type string to a slice of bytes(uint8)
	//basically it convert each letter to the correspondent ascii: http://sticksandstones.kstrom.com/appen.html

	
	fmt.Println("The conversion from type string to a slice of bytes is",bs)
	fmt.Println("Type is:") //[]uint8 ;
	fmt.Printf("%T\n",bs)
}