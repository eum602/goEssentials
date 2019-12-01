package main

import (
	"fmt"
	"runtime"
)

var a int
var b float32

func main() {
	//find all about numerics here: https://golang.org/ref/spec#Numeric_types
	//a=42.456 //this is not going to run because a is previously declared as int
	a = 42
	b = 42.56

	fmt.Println("a variable is", a, "its type is")
	fmt.Printf("%T\n", a)
	fmt.Println("b variable is", b, "its type is")
	fmt.Printf("%T\n", b)

	var c int8 //-128 ==> 127
	//c=-129 //this is not going to work
	fmt.Println("c variable is", c, "its type is")
	fmt.Printf("%T\n", c)

	//ALIASES:
	/*
	* byte: uint8
	* rune: int32
	 */

	fmt.Println("Using runtime package to get the OS", runtime.GOOS)
	fmt.Println("Using runtime package to print the running program architecture", runtime.GOARCH)

	//********************STRINGS*****************************//
	fmt.Println("**********************************************")
	fmt.Println("****************STRINGS***********************")
	s := "Hello this is nome non useful text"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	u := `This is a text created with backticks`

	fmt.Println(u)

	bs := []byte(s) //using conversion to convert from type string to a slice of bytes(uint8)
	//basically it convert each letter to the correspondent ascii: http://sticksandstones.kstrom.com/appen.html

	fmt.Println("The conversion from type string to a slice of bytes is", bs)
	fmt.Println("Type is:") //[]uint8 ;
	fmt.Printf("%T\n", bs)

	/* **********************************************************

	 */
	fmt.Println("****************Ways to print Strings with fmt***********************")
	message := `Ways to print Strings with fmt package; 
	%s	the uninterpreted bytes of the string or slice
	%q	a double-quoted string safely escaped with Go syntax
	%x	base 16, lower-case, two characters per byte
	%X	base 16, upper-case, two characters per byte
	%#U ==> to see the utf character
	check : https://godoc.org/fmt`
	fmt.Println(message)

	fmt.Println("\nPrinting utf characters for a text:")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U ", s[i]) //each utf code is expresed in int32 (rune)
	}

	fmt.Println("\n\nShowing the text in their correspondent bytes(uint8) in HEXADECIMAL format:")

	for i, v := range s {
		fmt.Printf("At index position %d we have the %#x\n", i, v) //%d ==> base 10; %x prints without leading format (0b
		//for  binary; 0x for hexadecimal and so on); if we add # (eg %#x) then it will print with the leading format.
		//read fmt package from godoc.org/fmt
	}

	//A string is a slice of bytes and is IMMUTABLE,  meaning that you cannot change
	//what is the value stored there. BUT I CAN ASSIGN A NEW VALUE
	s = "This another text, demonstrates that a new value can be REASSIGNED to a string BUT a string cannot be modified"
	fmt.Println(s)

	//adittional resource to understand strings in : https://blog.golang.org/strings

	//***********************constants***********************
	//https://golang.org/ref/spec
	const (
		w1 int     = 42
		w2 float64 = 42.78
		w3 string  = "something"
		//by adding types now those are TYPED constants
	)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	//***********************constants***********************
	//https://golang.org/ref/spec#Constants
	/*IOTA is a predeclared identifier*/
	fmt.Println("************************************IOTA CONSTANT************************************************")
	const (
		m = iota //3
		n        //4
		p        //5
	)

	const ( //resets!!! when we start a new set of constants
		q  = iota //0
		r         //1
		ss        //2
	)

	fmt.Println(m)
	fmt.Println(n)
	fmt.Println(p)
	/*We have only declared variables but numbers are automatically incremented*/
	fmt.Printf("%T\n", m)
	fmt.Printf("%T\n", n)
	fmt.Printf("%T\n", p)

	fmt.Println(q)
	fmt.Println(r)
	fmt.Println(ss)
	/*We have only declared variables but numbers are automatically incremented*/
	fmt.Printf("%T\n", q)
	fmt.Printf("%T\n", r)
	fmt.Printf("%T\n", ss)
}
