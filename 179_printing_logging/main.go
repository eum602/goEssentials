package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	example1()
	example3()
	example2()
}

func definitions() {
	fmt.Println(`
	Defer statements allow us to think about closing each file right after opening it, guaranteeing that, 
	regardless of the number of return statements in the function, the files will be closed.
	`)

}

func example1() {
	fmt.Println("*************Print logs example*************")
	f, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()  //closing at the end
	log.SetOutput(f) //setting that all log information input will go to the log.txt file

	//Now simulating an error to write it to a file
	f2, err := os.Open("non-existent-file.txt")
	if err != nil {
		log.Println("An error has happened: ", err) //this log will be written to log.txt
	}
	defer f2.Close()
}

func example2() {
	defer deferred() //this will not be executed because deferred functions will not be executed.
	fmt.Println("*************Log Fatal example: Log fatal invokes os.Exit(1) ==> which means exits with error*************")
	f2, err := os.Open("non-existent-file.txt")
	if err != nil {
		log.Fatal(err) //exiting with os 1 status code ==> exits with error
	}
	defer f2.Close()
}

func example3() {
	f, err := os.Create("log1.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()  //closing at the end
	log.SetOutput(f) //setting that all log information input will go to the log.txt file

	written, err := CopyFile("copy", "log.txt")
	if err != nil {
		log.Println("Error in example 3", err)
	}
	log.Println("Successfull copy", written)
}

func deferred() {
	fmt.Println("Deferred functions will not be executed when os.Exit(1) is invoked")
}

//CopyFile ...
func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.Create(dstName)
	if err != nil {
		return
	}
	defer dst.Close() //runs after return is executed. Guarantee even when errors happens the file will be closed.

	return io.Copy(dst, src)
}
