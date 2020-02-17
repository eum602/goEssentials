package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	definitions()
	example1()
	inputExample()
	createFile()
	readFile()
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

func example1() {
	n, err := fmt.Println("Some text")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)
}

func inputExample() {
	var ans1, ans2 string
	fmt.Println("Name: ")
	_, err := fmt.Scan(&ans1)
	if err != nil {
		panic(err)
	}

	fmt.Println("Last name: ")
	_, err = fmt.Scan(&ans2)
	if err != nil {
		panic(err)
	}
}

func createFile() {
	f, err := os.Create("myFile.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer f.Close() //closing after all

	reader := strings.NewReader("Some text to add")
	io.Copy(f, reader)
}

func readFile() {
	//opening the file
	r, err := os.Open("myFile.txt")
	if err != nil {
		fmt.Println(err)
	}

	text, err := ioutil.ReadAll(r)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Read text is: ", string(text))
}
