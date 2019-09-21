package main

import "fmt"

//reference :  https://godoc.org/fmt
func main() {
	y := 255
	fmt.Printf("%T\n", y)  //print the type of y
	fmt.Printf("%b\n", y)  //printing the value in base 2
	fmt.Printf("%X\n", y)  //prints the value in HEXADECIMAL in UPPERCASE
	fmt.Printf("%x\n", y)  //hexadecimal in lowercase
	fmt.Printf("%#x\n", y) //adding 0x

	fmt.Printf("%b\t%#x\t%X\n", y, y, y) //printing the same value in different format separated by \t => tab

	fmt.Printf("%v\n", y) //it prints the normal value

	//it is possible assign a output of fmt to a variable
	s := fmt.Sprintf("%b\t%#x\t%X\n", y, y, y)
	fmt.Println(s) //now printing that with a simple Println

	//there is also possible to print to a file by using fmt.Fprint

}
