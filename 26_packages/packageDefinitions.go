package main //means this code belongs to package main

import "fmt"

//you can find packages from golang.org but it is more clear to find them on :
//https://godoc.org/fmt or whatever package
//packages helps organize code into different folder and are ready to be IMPORTED and used
//to see how a package is composed go to its index section in godoc.org/fmt (foor example and check index section)
/*
there you will have something like this:
Index
...........
func Print(a ...interface{}) (n int, err error)
func Printf(format string, a ...interface{}) (n int, err error)
func Println(a ...interface{}) (n int, err error) ============> "...interface{}" means it can accept any amount of parameters
........
*/
//here for example method Println can accept mamy parameters as input and return s an int and an error which we can
//catch if we want
func main() {
	fmt.Println("Passing only one argument") //in this case from PACKAGE FMT we are getting some method.
	n, err := fmt.Println("one", "two", "any amount of arguments")
	//now catching the n and error
	fmt.Println("n is", n)       //returned 32 => indicates the number of bytes
	fmt.Println("error is", err) //this is only an example how you can interpret the index on each package and use it
	//at your convenience => returned nil => there was no error

	//ignoring error (similar to javascript)
	n, _ = fmt.Println("something")
	fmt.Println("n returned", n)
}
