package main

import "fmt"

var structDefinition = `Struct is a data structure which allows us to compose together values of different types. This
is known as THE COMPOSITE DATA STRUCTURE, also known as the AGGREGATE DATA STRUCTURE(for complex data strucure)`

type person struct { /*Definign a new TYPE of type struct*/
	first string //identifier and type is the way to declare each row in a struct
	last  string //also 'last' is the unqualified type name =>name of the type ; acts also as the field name
	//when declaring a new instance of a struct
}

type secretAgent struct {
	person //embedding type person into secretAgent type, this is the inner type, this type of declaration is
	//anonymous field or implicit declaration
	ltk bool
}

type secretAgentCollision struct {
	person
	ltk   bool
	first bool
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

	sa := secretAgent{
		person: p1,
		ltk:    true,
	}

	sa1 := secretAgent{
		person: person{
			first: "",
			last:  "",
		},
		ltk: false,
	}

	sa2 := secretAgentCollision{
		person: p2,
		ltk:    false,
		first:  true,
	}

	fmt.Println(p1.first, p1.last)
	fmt.Println(p2.first, p2.last)

	fmt.Println("Structs allows to omit the embeded type so we can omit person or not")
	fmt.Println(sa.first, sa.last, sa.ltk) //by using in this way we say that the inner type(person) gets PROMOTED
	//to the outer type(secretAgent)
	fmt.Println(sa1.person.first, sa1.person.last, sa1.ltk)

	fmt.Println("To avoid COLLISIONS we can opt to use the full path to access an attribute")
	fmt.Println(sa2.person.first, sa2.first)
}
