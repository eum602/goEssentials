package main

import "fmt"

func main() {
	x := 3
	fmt.Printf("%d\t\t%b\n", x, x)

	y := x << 1 //bitshifting 11 ===> 110 (increments 1 zeroes in the right hand side)
	fmt.Printf("%d\t\t%b\n", y, y)

	w := x << 2 //bitshifting 11 ===> 1100 (increments 2 zeroes in the right hand side)
	fmt.Printf("%d\t\t%b\n", w, w)
}
