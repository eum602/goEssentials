package main

import "fmt"

func main() {
	x := 3
	fmt.Printf("%d\t\t%b\n", x, x)

	y := x << 1 //bitshifting 11 ===> 110 (increments 1 zeroes in the right hand side)
	fmt.Printf("%d\t\t%b\n", y, y)

	w := x << 2 //bitshifting 11 ===> 1100 (increments 2 zeroes in the right hand side)
	fmt.Printf("%d\t\t%b\n", w, w)

	fmt.Println("\nPrinting equilalences between kilobytes, megabytes and gigiabytes")
	const (
		_  = iota             //throing the first iota away ; //0
		kb = 1 << (iota * 10) //iota => 1;  taking number one and shifting it over 10 so ==> 1 0 000 000 000 (binary) => kb => 2^10
		mb = 1 << (iota * 10) //iota => 2
		gb = 1 << (iota * 10) //iota => 3
	)

	fmt.Printf("%d\t\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t\t%b\n", kb, mb)
	fmt.Printf("%d\t\t\t%b\n", mb, gb)
}
