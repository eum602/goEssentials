package greet

import (
	"fmt"
)

//Greet returns a simple greet
func Greet(s string) string {
	return fmt.Sprint("Welcome ", s)
}
