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
		* go doc mymath 
		* go doc mymath.Sum
		* go doc github.com/eum602/goEssentials/190_myMath ==> so see specific package
		* go doc fmt
		* go doc json.Number
		* go doc json.Number.Float64
		* cd go/src/encoding/json; go doc decode <> go doc json.decode

		Install godoc with ==> sudo apt install golang-golang-x-tools
		godoc extracts and generates documentation for go programs.
		run it with godoc -http=:8080 or http :8080
		godoc fmt ==> similar to go doc fmt
		godoc fmt Printf ==> similar to go doc fmt.Printf
		godoc -src fmt Printf ==> shows also comments about the queried method (implementation details)

	godoc.org: Has third party package documentation.

	https://golang.org/pkg/ ==> standard libraries
	`)
}
