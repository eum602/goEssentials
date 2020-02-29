package main

import (
	"fmt"
)

func main() {
	customSort()
}

// Person  ...
type Person struct {
	Name string
	Age  int
}

//ByAge ...
type ByAge []Person

//Len ...
func (b ByAge) Len() int {
	return len(b)
}

//Swap ...
func (b ByAge) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

//Less
func (b ByAge) Less(i, j int) bool {
	return b[i].Age < b[j].Age
}

//*********************************

func customSort() {
	p1 := Person{
		Name: "Erick",
		Age:  32,
	}
	p2 := Person{
		Name: "Alice",
		Age:  45,
	}
	a := []Person{p1, p2}
	fmt.Println("Intial array of persons is\n", a)
	b := ByAge(a)
	fmt.Println(b.Len())
	fmt.Println(b.Less(0, 1))
	b.Swap(0, 1)
	fmt.Println(b)

}

func (p Person) String() string {
	return fmt.Sprintf("%s: %d", p.Name, p.Age)
}
