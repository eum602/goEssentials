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

func main() {
	fmt.Println(`Definition of function: 
	func (r receiver) identifier(parameters) (returns s){...}`)
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

	fmt.Println(sa1)
	sa1.speak()
	sa2.speak()
}

func (s secretAgent) speak() {
	fmt.Println(`This function can only be accessed by secretAgent type, in other words this function has been 
	ATTACHED to secretAgent type, so any value of type secretAgent can access here`)
	fmt.Println(`I'm `, s.first, s.last)
}
