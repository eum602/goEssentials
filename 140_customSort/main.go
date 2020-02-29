package main

import (
	"fmt"
	"sort"
)

func main() {
	definitions()
	customSortByAge()
	customSortByName()
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

func (p Person) String() string {
	return fmt.Sprintf("%s: %d", p.Name, p.Age)
}

var p1 = Person{
	Name: "Erick",
	Age:  32,
}
var p2 = Person{
	Name: "Alice",
	Age:  45,
}
var a = []Person{p1, p2}

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

//ByName ...
type ByName []Person

//Len ...
func (n ByName) Len() int {
	return len(n)
}

//Swap ...
func (n ByName) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

//Less
func (n ByName) Less(i, j int) bool { //ordering criteria: by age
	return n[i].Name < n[j].Name
}

//*********************************

func customSortByAge() {
	fmt.Println("\nIntial array of persons is\n", a)
	b := ByAge(a)
	fmt.Println(b.Len())
	fmt.Println(b.Less(0, 1))
	b.Swap(0, 1)
	fmt.Println(b)

	fmt.Println("Ordering By using sort with the implementation of 'Interface' with the ordering criteria by Age")
	sort.Sort(b)
	fmt.Println(b)
}

func customSortByName() {
	fmt.Println("\nOrdering by Name... initial array: ")
	n := ByName(a)
	fmt.Println(n)
	sort.Sort(n)
	fmt.Println("Ordering by Name, finally ordered: ")
	fmt.Println(n)
}
