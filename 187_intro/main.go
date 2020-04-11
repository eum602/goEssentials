package main

import (
	"fmt"
)

func main() {
	definitions()
}

func definitions() {
	fmt.Println(`
	go doc: part of the go command
	godoc: a command just like 'go' is a command. Allows us to read documentation at the terminal.
		* go help doc
		* go doc myMath 
		* go doc myMath.Sum
		* go doc fmt
		* go doc json.Number
		* go doc json.Number.Float64
		* cd go/src/encoding/json; go doc decode <> go doc json.decode

	godoc.org: Has third party package documentation.

	https://golang.org/pkg/ ==> standard libraries
	`)
}
