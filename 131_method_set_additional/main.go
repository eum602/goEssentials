package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type circle1 struct {
	radius float64
}

func main() {
	definitions()
}

func definitions() {
	c := newCircle(3)
	fmt.Println("function with a receiver of type T can be used by T or *T",
		"For example if area is: func (c circle) area() (returns float64) { return math.Pi * c.radius * c.radius}",
		"Then we can use c.area() or &c.area()")
	fmt.Println("Printing the area using c.area()", c.area())
	fmt.Println("Printing the area using (&c).area()", (&c).area())

	fmt.Println("\nfunction with a receiver of type *T can be used by T or *T",
		"For example if area1 is: func (c *circle) area1() (returns float64) { return math.Pi * c.radius * c.radius}",
		"Then we can use c.area1() or &c.area1()")
	fmt.Println("Printing the area using c.area1()", c.area1())
	fmt.Println("Printing the area using (&c).area1()", (&c).area1())
}

func (c circle) area() (returns float64) {
	return math.Pi * c.radius * c.radius
}

func (c *circle) area1() (returns float64) {
	return math.Pi * c.radius * c.radius
}

func newCircle(r float64) (returns circle) {
	fmt.Println("New circle created, radius: ", r)
	return circle{
		radius: r,
	}
}
