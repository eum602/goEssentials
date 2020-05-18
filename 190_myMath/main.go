// Package mymath provides an example how to write documentation
package mymath

//Sum function returns the sum of 'n' integer arguments
func Sum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}
