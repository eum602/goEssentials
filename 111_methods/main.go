package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

type human interface { //Any TYPE that has METHOD speak is ALSO of TYPE human
	speak()
}

func main() {
	fmt.Println(`Definition of function: 
	func (r receiver) identifier(parameters) (returns s){...}
	* INTERFACES: 
		- It is a type
		- General syntax :  type human interface { speak() } //keyword identifier type
			+interpretation: Any TYPE that has METHOD speak is ALSO of TYPE human
		- A value can be of more than one type
		
		}`)

	sa1 := secretAgent{
		person: person{
			"Erick",
			"Pacheco",
		},
		ltk: true,
	}

	sa2 := secretAgent{
		person: person{
			"John",
			"Doe",
		},
		ltk: true,
	}

	p1 := person{
		first: "Alice",
		last:  "Spl",
	}

	fmt.Println(`Because 'human' interface states that "Any type that has the method speack() is of type 'human'"
	AND speak function has been ATTACHED to be used by 'secretAgent' types
	THEN all VALUES of types 'secretAgent' are also type 'human'
	THEN => A value can be of more than one type `)

	bar(sa1)
	bar(sa2)
	bar(p1)
}

func bar(h human) {
	fmt.Println("\n\nI'm called human, let me speak:")
	h.speak()
}

func (s secretAgent) speak() {
	fmt.Println(`=======I AM INTO SPEAK FUNCTION ===========
	This function can only be accessed by secretAgent type, in other words this function has been 
	ATTACHED to secretAgent type, so any value of type secretAgent can access here`)
	fmt.Println(`I'm `, s.first, s.last, `- I'm of type secretAgent`)
}

func (p person) speak() {
	fmt.Println(`=======I AM INTO SPEAK FUNCTION ===========
	This function speak has been attached to values of type person`)
	fmt.Println(`I'm `, p.first, p.last, ` - I'm of type person`)
}
