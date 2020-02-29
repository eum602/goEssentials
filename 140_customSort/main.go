package main

import (
	"fmt"
	"sort"
)

func main() {
	customSort()
}

func definitions() {
	fmt.Println(`
	All types that implements the "Interface" interface are also of type "Interface" and thus can call "Sort" method of package "sort" ==>
	https://golang.org/pkg/sort/#Sort
	
	To be of type "Interface", types must implement Len(), Less() and Swap() methods ==>
	https://golang.org/pkg/sort/#Interface
	 `)
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
func (b ByAge) Less(i, j int) bool { //ordering criteria: by age
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

	fmt.Println("Ordering By using sort with the implementation of 'Interface' with the ordering criteria by Age")
	sort.Sort(b)
	fmt.Println(b)
}

func (p Person) String() string {
	return fmt.Sprintf("%s: %d", p.Name, p.Age)
}
