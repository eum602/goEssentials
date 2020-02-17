package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	example1()
	example2()
}

func example1() {
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
	fmt.Println("Log Fatal example: Log fatal invokes os.Exit(1) ==> which means exits with error")
	f2, err := os.Open("non-existent-file.txt")
	if err != nil {
		log.Fatal(err) //exiting with os 1 status code ==> exits with error
	}
	defer f2.Close()
}

func deferred() {
	fmt.Println("Deferred functions will not be executed when os.Exit(1) is invoked")
}
