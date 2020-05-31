package mystr

import "strings"

//Cat is a simple implementation of concatenate operation
func Cat(xs []string) string {
	s := ""
	for _, v := range xs {
		s += v
		s += " "
	}

	return s
}

//Join is a concatenate implemetation leveraged by string.Join built-in method
func Join(xs []string) string {
	return strings.Join(xs, " ")
}
