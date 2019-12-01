package main

import "fmt"

var structDefinition = `Struct is a data structure which allows us to compose together values of different types. This
is known as THE COMPOSITE DATA STRUCTURE, also known as the AGGREGATE DATA STRUCTURE(for complex data strucure)`

type person struct { /*Definign a new TYPE of type struct*/
	first string
	last  string
}

func main() {
	fmt.Println("Struct definition:\n", structDefinition, "\n")

	/*Now creating VALUES of type person*/
	p1 := person{
		first: "e",
		last:  "pachped",
	}

	p2 := person{
		first: "john",
		last:  "doe",
	}

	fmt.Println(p1.first, p1.last)
	fmt.Println(p2.first, p2.last)

}
