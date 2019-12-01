package main

import "fmt"

func main() {
	x := "Pointer is just a pointing to some location in memory for a value stored."
	fmt.Println(x)

	a := 43
	fmt.Println("The VALUE of variable 'a' is", a)
	fmt.Println("The ADDRESS of variable 'a' is", &a) //0xc000014140 => returns a direction https://golang.org/ref/spec#Operators_and_punctuation

	fmt.Println("The type of variable 'a' is:")
	fmt.Printf("%T\n", a) //int
	fmt.Println("The type of the address where variable 'a' is stored is:")
	fmt.Printf("%T\n", &a) //*int ===> means a POINTER TO AN int ; WHICH IS A VERY DIFFERENT TYPE
	//All (*int) included the '*' is part of the type

	/*To summarize you have an '&' whcihc gives you the address  where the value is stored. Then the address where
	the value is stored we are talking about the pointer to an int(in this example) => a pointer to where the int
	is stored(that pointer has its own type.
	*/

	//var b int = &a //=========> this wont work because &a is OF TYPE '*int' (in the example)
	//But we can make something like this
	//var b *int = &a or something like this:
	var b = &a
	fmt.Println("Assigning type pointer to another variable of the same type=> var *int = &a ======> b = ", b)

	/*Everything in go is passed by value, and now we are able to share addresses*/
	/*Of course we could also have made a simple assignment ==> var c := &a   */
}
