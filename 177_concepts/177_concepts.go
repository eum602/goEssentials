package main

import (
	"fmt"
)

func main() {
	definitions()
}

func definitions() {
	fmt.Println(`
	Golang does not have exceptions 
	Error handling in go is pleasant but quite different that in other languages.
	https://blog.golang.org/error-handling-and-go

	There exist the type error which is an interface that states only one function called Erro() ==>
	type error interface {
		Error() string
	} ==> so any type that implements an Error() function is also of type error.
	`)
}
