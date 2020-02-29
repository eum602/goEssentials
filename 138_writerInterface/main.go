package main

import (
	"fmt"
	"os"
)

func main() {
	definitions()
	fmtExample()
}

func definitions() {
	fmt.Println(`
	**** Writer interface from package "IO"****
	A. type Writer interface {
		Write(p []byte) (n int, err error)    =====> any other type that has a Write method is also of type Writer.
		Close() error
	}

	**** *File type of "OS" package ****
	B. func (f *File) Write(b []byte) (n int, err error) ====> *File type implements the method Write then it is also of type Writer
	(https://golang.org/pkg/os/#File.Write)
	
	**** type NewEncoder of package json ****
	https://golang.org/pkg/encoding/json/#NewEncoder
	func NewEncoder(w io.Writer) *Encoder ====> returns a pointer to an encoder that I can use to straight write to a file.

	**** PACKAGE fmt println ****
	C. func Println(a ...interface{}) (n int, err error) {
		return Fprintln(os.Stdout, a...) ==> os.Stdout is of type io.Writer
	}

	D. func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) {
		p := newPrinter()
		p.doPrintf(format, a)
		n, err = w.Write(p.buf)
		p.free()
		return
	}

	E. func Fprintln(w io.Writer, a ...interface{}) (n int, err error) {
		p := newPrinter()
		p.doPrintln(a)
		n, err = w.Write(p.buf)
		p.free()
		return
	}

	F. https://golang.org/src/os/file.go?s=4844:4893#L139
	func (f *File) Write(b []byte) (n int, err error) {...}

	When calling Fprintln we pass os.stdout as a first parameter; according to golang source code
	os.Stdout is a variable of type *File ; and according to B and F *File implements Write interface
	thus os.Stdout is of type(Writer)

	`)
}

func fmtExample() {
	msg := "Hello W."
	fmt.Println(msg)
	fmt.Println("\n ")
	fmt.Fprintf(os.Stdout, msg) //first parameter receives types that implemets Writer function
	fmt.Println("\n ")
	fmt.Fprintln(os.Stdout, msg)
}
