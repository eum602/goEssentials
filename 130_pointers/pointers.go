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
	fmt.Println(`Using https://golang.org/ref/spec#Operators_and_punctuation ==> * (ASTERISC OPERATOR)
	# When using * with a pointer(where the the variable is located) ==> then it DEREFERENCE, so it indicates what is the value
		of what is stored on that address:
	`)

	fmt.Println("I stored an address in 'b' => b = &a.")
	fmt.Println("The value stored on b is", b)
	fmt.Println("The type of b is:")
	fmt.Printf("%T\n", b)
	fmt.Println("As 'b' is of type pointer(*int) then by using *b golang returns us whatever is stored on that address=> *b: ", *b)
	fmt.Println("Also b is stored on the address &b", &b)
	fmt.Println("Also as explained above we can DEREFERENCE the value stored on an address(&b) with '*' operator, then *&b: ", *&b)
	fmt.Println(*&b)

	*b = 87
	fmt.Println(`Lets make an assing by knowing the address!! ==> b is an address of type *int. *b references to the value
	stored on that address; then we can ASSIGN  A NEW VALUE to that ADDRESS!!! => *b = 
	`, *b, "!!!")

	/*Everything in go is passed by value, and now we are able to share addresses*/
	/*Of course we could also have made a simple assignment ==> var c := &a   */

	fmt.Println(`Pointers can be used when:
		* You want to change the value of of something that is at certain location.
		* Pass the address where lot of data is stored.
	`)

	fmt.Println("Everything in GO is pass by value, what you see is what you get.")

	z := 0
	fmt.Println("z before", &z)
	fmt.Println("z before", z)
	foo(&z) //key point is that now we are passing the address
	fmt.Println("z after", &z)
	fmt.Println("z after", z) //z has changed!! because now function foo modifies it by using pointers

}

func foo(y *int) { //the received value is ASSIGNED to "y"
	fmt.Println("y before", y)
	fmt.Println("y before", *y)
	*y = 43 //dereferencing to get the value AT the address of y
	fmt.Println("y after", y)
	fmt.Println("y after", *y)
}
