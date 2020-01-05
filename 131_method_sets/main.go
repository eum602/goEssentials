package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

func main() {
	definitions()
}

func definitions() {
	fmt.Println(`
	* Methods that are attached to a certain type are known as METHOD-SETS, that is, a set of methods attached
	to a type.
	* Method sets are a type => https://golang.org/ref/spec#Method_sets
	* Method set are all the methods associated with some type.	
	`)

	c := newCircle(3)
	printArea(c)
}

func (c circle) area() (returns float64) {
	fmt.Println("area() is a method attached to type 'circle'")
	return math.Pi * c.radius * c.radius
}

func newCircle(r float64) (returns circle) {
	fmt.Println("New circle created, radius: ", r)
	return circle{
		radius: r,
	}
}

func printArea(c circle) {
	fmt.Println("The area of the circle is", c.area())
}
